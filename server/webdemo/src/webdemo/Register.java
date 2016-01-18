package webdemo;

import java.io.IOException;

import javax.servlet.RequestDispatcher;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;

public class Register extends DBServlet {

	/**
	 * 
	 */
	private static final long serialVersionUID = 1L;
	private static Logger logger = Logger.getLogger(Register.class);

	@Override
	protected void service(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		String userName = null;
		if (request.getParameter("login") != null) {
			response.sendRedirect("login.jsp");
			return;
		}
		try {
			super.service(request, response);
			userName = request.getParameter("username");
			String password = request.getParameter("password");
			String email = request.getParameter("email");
			String validationCode = request.getParameter("validation_code");
			if (userName == null || password == null || validationCode == null)
				return;
			if (userName.equals("") || password.equals("") || validationCode.equals(""))
				return;
			userName = new String(userName.getBytes("ISO-8859-1"), "UTF-8");
			request.setAttribute("page", "register.jsp");
			if (!checkValidationCode(request, validationCode))
				return;
			email = (email == null) ? "" : email;
			String passwordMD5 = webdemo.Encrypter.md5Encrypt(password);
			String sql = "insert into table_users(User_Name,User_Password,User_Email) values(?,?,?)";
			execSQL(sql, userName, passwordMD5, email);
			request.setAttribute("info", "用户注册成功");
			logger.info(userName + "用户注册成功");
		} catch (Exception e) {
			request.setAttribute("info", userName + "用户已存在");
			logger.info("用户:\"" + userName + "\"已存在");
		}

		finally {
			RequestDispatcher rd = request.getRequestDispatcher("result.jsp");
			rd.forward(request, response);
		}
	}
}
