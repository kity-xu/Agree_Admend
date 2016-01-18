package webdemo;

import java.io.IOException;
import java.sql.ResultSet;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

import org.apache.log4j.Logger;

public class Login extends DBServlet {

	/**
	 * 
	 */
	private static final long serialVersionUID = -1249482337187956622L;
	private static Logger logger = Logger.getLogger(Login.class);

	@Override
	protected void service(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		if (request.getParameter("register") != null) {
			response.sendRedirect("register.jsp");
			return;
		}
		int page = 0;
		String userName = "";
		try {
			super.service(request, response);
			userName = request.getParameter("user_name");
			String password = request.getParameter("user_password");
			String validationCode = request.getParameter("validation_code");
			if (userName == null || password == null || validationCode == null)
				return;

			if (userName == "" || password == "" || validationCode == "")
				return;

			userName = new String(userName.getBytes("ISO-8859-1"), "UTF-8");
			if (!checkValidationCode(request, validationCode))
				return;

			String sql = "select user_name,user_password from table_users where user_name=?";
			ResultSet rs = execSQL(sql, new Object[] { userName });
			if (rs.next() == false) {
				request.setAttribute("userError", userName + "不存在");
			} else {
				String passwordMD5 = webdemo.Encrypter.md5Encrypt(password);
				if (!rs.getString("user_password").equals(passwordMD5)) {
					request.setAttribute("passwordError", "密码不正确");
				} else {
					HttpSession session = request.getSession();
					session.setAttribute("userName", userName);
					page = 1;
				}
			}
		} catch (Exception e) {
		} finally {
			request.setAttribute("userName", userName);
			if (page == 0) {
				request.getRequestDispatcher("login.jsp").forward(request, response);//显示错误信息
			} else {
				response.sendRedirect("index.jsp");//令主页可以刷新
			}
		}
	}

	public void logout(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		logger.info("logout");
		HttpSession session = request.getSession();
		if (session.getAttribute("userName") != null) {
			logger.info(session.getAttribute("userName") + " is logout");
			session.setAttribute("userName", null);
		}
	}
}
