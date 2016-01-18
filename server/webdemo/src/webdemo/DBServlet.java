package webdemo;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;

public class DBServlet extends HttpServlet {

	/**
	 * 
	 */
	private static final long serialVersionUID = 3464282311164003379L;
	private static Logger logger = Logger.getLogger(DBServlet.class);
	protected java.sql.Connection conn = null;

	protected java.sql.ResultSet execSQL(String sql, Object... args) throws Exception {
		java.sql.PreparedStatement pStmt = conn.prepareStatement(sql);
		for (int i = 0; i < args.length; i++) {
			pStmt.setObject(i + 1, args[i]);
		}
		pStmt.execute();
		return pStmt.getResultSet();
	}

	protected boolean checkValidationCode(HttpServletRequest request, String validationCode) {

		String validationCodeSession = (String) request.getSession().getAttribute("validation_code");
		if (validationCodeSession == null) {
			request.setAttribute("info", "验证码过期");
			request.setAttribute("codeError", "验证码过期");
			return false;
		}

		if (!validationCode.equalsIgnoreCase(validationCodeSession)) {
			request.setAttribute("info", "验证码不正确");
			request.setAttribute("codeError", "验证码不正确");
			return false;
		}
		return true;
	}

	@Override
	protected void service(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {

		try {
			if (conn == null) {
				javax.naming.Context ctx = new javax.naming.InitialContext();
				// TODO 配置数据库
				javax.sql.DataSource ds = (javax.sql.DataSource) ctx.lookup("java:/comp/env/jdbc/test");
				conn = ds.getConnection();
			}
		} catch (Exception e) {
			logger.info("DBconnection is wrong!");
			response.sendError(500);
		}
	}

	@Override
	public void destroy() {
		try {
			if (conn != null) {
				conn.close();
			}
		} catch (Exception e) {

		}
	}

}
