package dbservlet;

import java.io.IOException;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.ArrayList;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;

import bean.AppPermission;
import bean.CurrentMac;
import bean.Page;
import bean.MacPermission;
import bean.PermissionPassword;

public class MyServlet extends HttpServlet {

	private static final long serialVersionUID = 1L;
	private static Logger logger = Logger.getLogger(MyServlet.class);

	public static enum Servlet {
		MacServlet, AppServlet, PerpasswordServlet, CurrentMacServlet
	}

	private String key1 = "", key2 = "", key3 = "";
	private String table = "", field1 = "", field2 = "", field3 = "";
	private String attribute1 = "", attribute2 = "", attribute3 = "";
	private String delete = "", idname = "", success = "", layout = "", input = "", nothing = "", updated = "",
			update = "";

	// TODO 配置数据库
	public Connection connect() throws ClassNotFoundException, SQLException {
		Connection conn = null;
		Class.forName("com.mysql.jdbc.Driver");
		String url = "jdbc:mysql://192.168.1.185:3306/test20150902?characterEncoding=UTF-8";
		String user = "test_server";
		String password = "qwerty";
		conn = DriverManager.getConnection(url, user, password);
		return conn;
	}

	public void close(Statement stat, Connection conn) throws SQLException {
		if (stat != null) {
			stat.close();
		}
		if (conn != null) {
			conn.close();
		}
	}

	public void changePage(HttpServletRequest request, Servlet servlet) {
		switch (servlet) {
		case MacServlet:
			delete = "/dbservlet/delete.jsp";
			idname = "/dbservlet/idnameselect.jsp";
			success = "/dbservlet/InsertSuccess.jsp";
			layout = "/dbservlet/layout.jsp";
			input = "/dbservlet/input.jsp";
			nothing = "/dbservlet/selectnothing.jsp";
			updated = "/dbservlet/updated.jsp";
			update = "/dbservlet/update.jsp";
			table = "table_mac";
			field1 = "MAC_address";
			field2 = "MAC_name";
			field3 = "MAC_authorization";
			key1 = "macAddress";
			key2 = "macName";
			key3 = "macPermission";
			break;

		case AppServlet:
			delete = "/appfile/appdelete.jsp";
			idname = "/appfile/appidnameselect.jsp";
			success = "/appfile/appInsertSuccess.jsp";
			layout = "/appfile/applayout.jsp";
			input = "/appfile/appupload.jsp";
			nothing = "/appfile/appselectnothing.jsp";
			updated = "/appfile/appupdated.jsp";
			update = "/appfile/appupdate.jsp";
			table = "table_app";
			field1 = "App_Name";
			field2 = "App_Path";//TODO 可能应调整为应用信息
			field3 = "App_Authorization";
			key1 = "appName";
			key2 = "appPath";
			key3 = "appPermission";
			break;

		case PerpasswordServlet:
			delete = "/perpasswordfile/perdelete.jsp";
			idname = "/perpasswordfile/peridnameselect.jsp";
			success = "/perpasswordfile/perInsertSuccess.jsp";
			layout = "/perpasswordfile/perlayout.jsp";
			input = "/perpasswordfile/perinput.jsp";
			nothing = "/perpasswordfile/perselectnothing.jsp";
			updated = "/perpasswordfile/perupdated.jsp";
			update = "/perpasswordfile/perupdate.jsp";
			table = "table_authorization";
			field1 = "authorization";
			field2 = "authorization_info";
			field3 = "authorization_password";
			key1 = "permissionID";
			key2 = "permissionInfo";
			key3 = "permissionPassword";

			break;

		default:
			break;
		}

		if (request != null) {
			attribute1 = request.getParameter(key1);
			attribute2 = request.getParameter(key2);
			attribute3 = request.getParameter(key3);
		}
	}

	public void insert(HttpServletRequest request, HttpServletResponse response, Servlet servlet)
			throws ClassNotFoundException, SQLException, IOException, ServletException {
		Connection conn = null;
		Statement stat = null;
		ResultSet rs = null;
		conn = connect();
		stat = conn.createStatement();

		changePage(request, servlet);
		String page = input;
		rs = stat.executeQuery("select * from " + table + " where " + field1 + " = '" + attribute1 + "'");
		if (rs.next() == false) {
			stat.execute("insert into " + table + "(" + field1 + "," + field2 + "," + field3 + ")values('" + attribute1
					+ "','" + attribute2 + "','" + attribute3 + "')");
			page = success;
		} else {
			request.setAttribute("userError", attribute1 + "记录已存在！");
		}
		close(stat, conn);
		request.getRequestDispatcher(page).forward(request, response);
	}

	public ArrayList<MacPermission> macSelect(String attr1, String attr2) throws ClassNotFoundException, SQLException {
		Connection conn = null;
		Statement stat = null;
		ResultSet rs = null;
		conn = connect();
		stat = conn.createStatement();

		if (attr1 == "" && attr2 == "")
			changePage(null, Servlet.MacServlet);

		ArrayList<MacPermission> result = new ArrayList<MacPermission>();
		if (attr1 == "" && attr2 == "") {
			rs = stat.executeQuery("select * from " + table);
		}
		if (attr1 != "" && attr2 == "") {
			rs = stat.executeQuery("select * from " + table + " where " + field1 + " = '" + attr1 + "'");
		}
		if (attr1 == "" && attr2 != "") {
			rs = stat.executeQuery("select * from " + table + " where " + field2 + " = '" + attr2 + "'");
		}
		if (attr1 != "" && attr2 != "") {
			rs = stat.executeQuery("select * from " + table + " where " + field1 + " = '" + attr1 + "' and " + field2
					+ " = '" + attr2 + "'");
		}
		while (rs.next()) {
			MacPermission st = new MacPermission();
			st.setMacAddress(rs.getString(field1));
			st.setMacName(rs.getString(field2));
			st.setPermission(rs.getString(field3));
			result.add(st);
		}
		if (rs != null) {
			rs.close();
		}
		close(stat, conn);
		return result;
	}

	public ArrayList<AppPermission> appSelect(String attr1, String attr2) throws ClassNotFoundException, SQLException {
		Connection conn = null;
		Statement stat = null;
		ResultSet rs = null;
		conn = connect();
		stat = conn.createStatement();
		if (attr1 == "" && attr2 == "")
			changePage(null, Servlet.AppServlet);

		ArrayList<AppPermission> result = new ArrayList<AppPermission>();
		if (attr1 == "" && attr2 == "") {
			rs = stat.executeQuery("select * from " + table);
		}
		if (attr1 != "" && attr2 == "") {
			rs = stat.executeQuery("select * from " + table + " where " + field1 + " = '" + attr1 + "'");
		}
		if (attr1 == "" && attr2 != "") {
			rs = stat.executeQuery("select * from " + table + " where " + field2 + " = '" + attr2 + "'");
		}
		if (attr1 != "" && attr2 != "") {
			rs = stat.executeQuery("select * from " + table + " where " + field1 + " = '" + attr1 + "' and " + field2
					+ " = '" + attr2 + "'");
		}
		while (rs.next()) {
			AppPermission st = new AppPermission();
			st.setAppName(rs.getString(field1));
			st.setAppPath(rs.getString(field2));
			st.setPermission(rs.getString(field3));
			result.add(st);
		}
		if (rs != null) {
			rs.close();
		}
		close(stat, conn);
		return result;
	}

	public ArrayList<PermissionPassword> permSelect(String attr1, String attr2)
			throws ClassNotFoundException, SQLException {
		Connection conn = null;
		Statement stat = null;
		ResultSet rs = null;
		conn = connect();
		stat = conn.createStatement();
		if (attr1 == "" && attr2 == "")
			changePage(null, Servlet.PerpasswordServlet);

		ArrayList<PermissionPassword> result = new ArrayList<PermissionPassword>();
		if (attr1 == "" && attr2 == "") {
			rs = stat.executeQuery("select * from " + table);
		}
		if (attr1 != "" && attr2 == "") {
			rs = stat.executeQuery("select * from " + table + " where " + field1 + " = '" + attribute1 + "'");
		}
		if (attr1 == "" && attr2 != "") {
			rs = stat.executeQuery("select * from " + table + " where " + field2 + " = '" + attribute2 + "'");
		}
		if (attr1 != "" && attr2 != "") {
			rs = stat.executeQuery("select * from " + table + " where " + field1 + " = '" + attribute1 + "' and "
					+ field2 + " = '" + attribute2 + "'");
		}
		while (rs.next()) {
			PermissionPassword st = new PermissionPassword();
			st.setPermissionId(rs.getString(field1));
			st.setPermissionInfo(rs.getString(field2));
			st.setPassword(rs.getString(field3));
			result.add(st);
		}
		if (rs != null) {
			rs.close();
		}
		close(stat, conn);
		return result;
	}

	public void dispatch(HttpServletRequest request, HttpServletResponse response, Servlet servlet)
			throws ClassNotFoundException, SQLException, ServletException, IOException {
		changePage(request, servlet);
		switch (servlet) {
		case MacServlet:
			if (macSelect(attribute1, attribute2).isEmpty()) {
				request.getRequestDispatcher(nothing).forward(request, response);
			} else {
				request.setAttribute("result", macSelect(attribute1, attribute2));
				request.getRequestDispatcher(idname).forward(request, response);
			}
			break;
		case AppServlet:
			if (appSelect(attribute1, attribute2).isEmpty()) {
				request.getRequestDispatcher(nothing).forward(request, response);
			} else {
				request.setAttribute("result", appSelect(attribute1, attribute2));
				request.getRequestDispatcher(idname).forward(request, response);
			}
			break;
		case PerpasswordServlet:
			if (permSelect(attribute1, attribute2).isEmpty()) {
				request.getRequestDispatcher(nothing).forward(request, response);
			} else {
				request.setAttribute("result", permSelect(attribute1, attribute2));
				request.getRequestDispatcher(idname).forward(request, response);
			}
			break;
		case CurrentMacServlet:
			break;
		default:
			break;
		}
	}

	public Page setPage(HttpServletRequest request, HttpServletResponse response, Servlet servlet)
			throws ClassNotFoundException, SQLException {
		String crd = request.getParameter("currentRecord");
		Page pager = new Page();

		switch (servlet) {
		case MacServlet:
			ArrayList<MacPermission> ret1 = macSelect("", "");
			pager.setTotalRecord(ret1.size());
			pager.setTotalPage(ret1.size(), pager.getPageSize());
			break;

		case AppServlet:
			ArrayList<AppPermission> ret2 = appSelect("", "");
			pager.setTotalRecord(ret2.size());
			pager.setTotalPage(ret2.size(), pager.getPageSize());
			break;

		case PerpasswordServlet:
			ArrayList<PermissionPassword> ret3 = permSelect("", "");
			pager.setTotalRecord(ret3.size());
			pager.setTotalPage(ret3.size(), pager.getPageSize());
			break;
		case CurrentMacServlet:
			ArrayList<CurrentMac> ret4 = selectCurrentMac();
			pager.setTotalRecord(ret4.size());
			pager.setTotalPage(ret4.size(), pager.getPageSize());
			break;
		default:
			break;
		}

		if (crd != null) {
			int currentRecord = Integer.parseInt(crd);
			pager.setCurrentRecord(currentRecord);
			pager.setCurrentPage(currentRecord, pager.getPageSize());
		}
		return pager;
	}

	public void difPage(HttpServletRequest request, HttpServletResponse response, Servlet servlet)
			throws ServletException, IOException, ClassNotFoundException, SQLException {
		changePage(request, servlet);
		request.getRequestDispatcher(layout).forward(request, response);
	}

	public void delete(HttpServletRequest request, HttpServletResponse response, Servlet servlet)
			throws ClassNotFoundException, SQLException, ServletException, IOException {
		Connection conn = null;
		Statement stat = null;
		conn = connect();
		stat = conn.createStatement();

		changePage(request, servlet);
		stat.execute("delete from " + table + " where " + field1 + " = '" + attribute1 + "'");
		request.getRequestDispatcher(delete).forward(request, response);
	}

	public void update(HttpServletRequest request, HttpServletResponse response, Servlet servlet)
			throws ClassNotFoundException, SQLException, ServletException, IOException {
		changePage(request, servlet);

		switch (servlet) {
		case MacServlet:
			request.setAttribute("result", macSelect(attribute1, ""));
			break;

		case AppServlet:
			request.setAttribute("result", appSelect(attribute1, ""));
			break;

		case PerpasswordServlet:
			request.setAttribute("result", permSelect(attribute1, ""));
			break;
		case CurrentMacServlet:
			break;
		default:
			break;
		}
		request.getRequestDispatcher(update).forward(request, response);
	}

	public void updated(HttpServletRequest request, HttpServletResponse response, Servlet servlet)
			throws ClassNotFoundException, SQLException, ServletException, IOException {
		Connection conn = null;
		Statement stat = null;
		conn = connect();
		stat = conn.createStatement();

		changePage(request, servlet);
		stat.execute("update " + table + " set " + field1 + " = '" + attribute1 + "'," + field2 + " = '" + attribute2
				+ "'," + field3 + " = '" + attribute3 + "' where " + field1 + " = '" + attribute1 + "'");

		switch (servlet) {
		case MacServlet:
			request.setAttribute("result", macSelect(attribute1, ""));
			break;

		case AppServlet:
			request.setAttribute("result", appSelect(attribute1, ""));
			break;
		case PerpasswordServlet:
			request.setAttribute("result", permSelect(attribute1, ""));
			break;
		case CurrentMacServlet:
			break;
		default:
			break;
		}
		request.getRequestDispatcher(updated).forward(request, response);
	}

	public ArrayList<CurrentMac> selectCurrentMac() throws ClassNotFoundException, SQLException {
		Connection conn = null;
		Statement stat = null;
		conn = connect();
		stat = conn.createStatement();
		ArrayList<CurrentMac> result = new ArrayList<CurrentMac>();
		ResultSet rs = null;
		CurrentMac st = new CurrentMac();
		rs = stat.executeQuery("select MAC_Address,Date_Add from view_current_mac;");
		while (rs.next()) {
			st = new CurrentMac();
			st.setMacAddress(rs.getString("MAC_Address"));
			st.setDateAdd(rs.getString("Date_Add"));
			result.add(st);
		}
		if (rs != null) {
			rs.close();
		}
		close(stat, conn);
		return result;
	}

	public ArrayList<CurrentMac> displayAllPostion() throws ClassNotFoundException, SQLException {
		Connection conn = null;
		Statement stat = null;
		conn = connect();
		stat = conn.createStatement();
		ArrayList<CurrentMac> result = new ArrayList<CurrentMac>();
		ResultSet rs = null;
		ResultSet rs2 = null;
		CurrentMac st = new CurrentMac();
		String[] macAddress = new String[255];
		String[] dateAdd = new String[255];
		int i = 0;
		rs = stat.executeQuery("select MAC_Address,Date_Add from view_current_mac;");
		while (rs.next()) {
			macAddress[i] = rs.getString("MAC_address").toUpperCase();
			dateAdd[i] = rs.getString("Date_Add");
			i++;
		}
		while (i > 0) {
			st = new CurrentMac();
			i--;
			rs2 = stat.executeQuery(
					"select Coordinate_Longitude,Coordinate_Latitude from view_mac_coordinate where MAC_address='"
							+ macAddress[i] + "' and Date_Add='" + dateAdd[i] + "';");
			if (rs2.next()) {
				st.setMacAddress(macAddress[i]);
				st.setLongitude(rs2.getString("Coordinate_Longitude"));
				st.setLatitude(rs2.getString("Coordinate_Latitude"));
				result.add(st);
			}
		}
		if (rs != null) {
			rs.close();
		}
		if (rs2 != null) {
			rs2.close();
		}
		close(stat, conn);
		return result;
	}

	public ArrayList<CurrentMac> viewPath(String macAddress) throws ClassNotFoundException, SQLException {
		Connection conn = null;
		Statement stat = null;
		conn = connect();
		stat = conn.createStatement();
		ArrayList<CurrentMac> result = new ArrayList<CurrentMac>();
		ResultSet rs = null;
		ResultSet rsc = null;
		CurrentMac st = new CurrentMac();
		rsc = stat.executeQuery(
				"select count(*) as rowCount from view_mac_coordinate where MAC_address='" + macAddress + "';");
		rsc.next();
		int rowCount = rsc.getInt("rowCount");
		rs = stat.executeQuery("select * from view_mac_coordinate where MAC_address='" + macAddress + "';");
		while (rs.next()) {
			st = new CurrentMac();
			st.setDateAdd(rs.getString("Date_Add"));
			st.setLongitude(rs.getString("Coordinate_Longitude"));
			st.setLatitude(rs.getString("Coordinate_Latitude"));
			logger.info(st.getDateAdd() + st.getMacAddress() + st.getLatitude() + st.getLongitude());
			result.add(st);
			// TODO 解决取过多点无法显示的问题
			for (int r = 0; r < rowCount / 5; r++) {// 取中间五个左右坐标，如果过多超过了JS发送长度就会无法打开
				if (!rs.next()) {
					break;
				}
			}
		}
		if (rs != null) {
			rs.close();
		}
		close(stat, conn);
		return result;
	}
}
