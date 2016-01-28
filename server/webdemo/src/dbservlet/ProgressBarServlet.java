package dbservlet;

import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

public class ProgressBarServlet extends HttpServlet {

	/**
	 * 
	 */
	private static final long serialVersionUID = 1L;

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, java.io.IOException {
		HttpSession session = request.getSession();
		Object is_begin = session.getAttribute("IS_UPLOAD_BEGIN");
		if (is_begin == null) {
			return;
		}
		if ("0".equals(is_begin.toString())) {
			return;
		}
		PrintWriter out = response.getWriter();
		Object upload_percentage = session.getAttribute("UPLOAD_PERCENTAGE");
		out.write("{percentage:'" + upload_percentage.toString() + "'}");
		out.flush();
		out.close();
	}

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, java.io.IOException {
		doPost(request, response);
	}
}
