package dbservlet;

import java.io.IOException;
import java.sql.SQLException;
import java.util.ArrayList;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;

import bean.CurrentMac;
import bean.Page;

/**
 * Servlet implementation class CurrentMacServlet
 */
public class CurrentMacServlet extends MyServlet {
	private static final long serialVersionUID = 1L;
	private static Logger logger = Logger.getLogger(CurrentMacServlet.class);

	/**
	 * @see HttpServlet#HttpServlet()
	 */
	public CurrentMacServlet() {
	}

	/**
	 * @see HttpServlet#doGet(HttpServletRequest request, HttpServletResponse
	 *      response)
	 */
	protected void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		logger.info("doGet");
		doPost(request, response);
	}

	/**
	 * @see HttpServlet#doPost(HttpServletRequest request, HttpServletResponse
	 *      response)
	 */
	protected void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		logger.info("doPost");
		request.setCharacterEncoding("UTF-8");
		response.setCharacterEncoding("UTF-8");
	}

	public Page setPage(HttpServletRequest request, HttpServletResponse response)
			throws ClassNotFoundException, SQLException {
		return super.setPage(request, response, Servlet.CurrentMacServlet);
	}

	public ArrayList<CurrentMac> selectCurrentMac() throws ClassNotFoundException, SQLException {
		return super.selectCurrentMac();
	}

	public ArrayList<CurrentMac> displayAllPostion() throws ClassNotFoundException, SQLException {
		return super.displayAllPostion();
	}

	public ArrayList<CurrentMac> viewPath(String address) throws ClassNotFoundException, SQLException {
		return super.viewPath(address);
	}

}
