package dbservlet;

import java.io.IOException;
import java.sql.SQLException;
import java.util.ArrayList;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;

import bean.MacPermission;
import bean.Page;

public class MacServlet extends MyServlet {
	/**
	 * 
	 */
	private static final long serialVersionUID = -6800823822498700665L;
	private static Logger logger = Logger.getLogger(MacServlet.class);

	public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		request.setCharacterEncoding("UTF-8");
		response.setCharacterEncoding("UTF-8");
		String methodName = request.getParameter("methodName");
		int method = Integer.parseInt(methodName);
		try {
			switch (method) {
			case 0:
				super.insert(request, response, Servlet.MacServlet);
				break;
			case 1:
				super.difPage(request, response, Servlet.MacServlet);
				break;
			case 2:
				// TODO 删除时应该先依照约束删除其他的内容后再删除。或者不提供删除功能，通过修改为000权限来实现。
				logger.info("未解除约束的删除操作");
				super.delete(request, response, Servlet.MacServlet);
				break;
			case 3:
				super.updated(request, response, Servlet.MacServlet);
				break;
			case 4:
				super.update(request, response, Servlet.MacServlet);
				break;
			case 5:
				super.dispatch(request, response, Servlet.MacServlet);
				break;
			}
		} catch (ClassNotFoundException e) {
			// e.printStackTrace();
			logger.info(e);
		} catch (SQLException e) {
			// e.printStackTrace();
			logger.info(e);
		}
	}

	public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		doPost(request, response);
	}

	public ArrayList<MacPermission> select(String attr1, String attr2) throws ClassNotFoundException, SQLException {
		return super.macSelect(attr1, attr2);
	}

	public Page setPage(HttpServletRequest request, HttpServletResponse response)
			throws ClassNotFoundException, SQLException {
		return super.setPage(request, response, Servlet.MacServlet);
	}

}
