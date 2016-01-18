package dbservlet;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.PrintWriter;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;

import net.sf.json.JSONObject;

public class AndroidServlet extends HttpServlet {
	// TODO 独立到单独的项目中，把客户端服务和管理端服务区分开

	/**
	 * 
	 */
	private static final long serialVersionUID = 614031498133677806L;
	private static Logger logger = Logger.getLogger(AndroidServlet.class);

	// TODO 配置数据库
	public Connection connect() throws ClassNotFoundException, SQLException {
		Connection conn = null;
		Class.forName("com.mysql.jdbc.Driver");
		String url = "jdbc:mysql://192.168.1.185:3306/test20150902?characterEncoding=UTF-8";
		String user = "test_android";
		String password = "qwerty";
		conn = DriverManager.getConnection(url, user, password);
		return conn;
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		request.setCharacterEncoding("UTF-8");
		response.setCharacterEncoding("UTF-8");
		try {
			Connection conn = null;
			Statement stat = null;
			ResultSet rs = null;
			conn = connect();
			stat = conn.createStatement();
			BufferedReader reader = request.getReader();
			String info = "";
			info = reader.readLine();
			String logData = info.replace("\"", "").replace("{", "").replace("}", "").replace(",", "");// 流水
			String logMac = "";
			JSONObject json = net.sf.json.JSONObject.fromObject(info);
			response.setCharacterEncoding("UTF-8");
			response.setContentType("application/json; charset=utf-8");
			PrintWriter out = null;
			String macAddress = json.getString("MAC");
			logMac = macAddress;
			logger.info("macAddress:" + macAddress);
			ResultSet row = stat.executeQuery("select count(*) as rowCount from view_mac_app where MAC_Address = '"
					+ macAddress.toUpperCase() + "'");
			row.next();
			int rowCount = row.getInt("rowCount");
			if (rowCount != 0) {// MAC校验
				logger.info("Mac校验正确");
				String method = json.getString("Key");
				logger.info("method:" + method);
				String locJson = json.getString("Location");
				JSONObject locati = net.sf.json.JSONObject.fromObject(locJson);
				String lo = locati.getString("Longitude");
				String la = locati.getString("Latitude");
				logger.info("location:" + locJson.toString());
				if (chkpos(Double.valueOf(lo), Double.valueOf(la))) {// 校验坐标范围
					// TODO 增加method以增加操作
					if ("CheckPwd".equals(method)) {// 判断操作是记录地址还是校验密码
						ResultSet rows = stat
								.executeQuery("select Authorization_Password from view_mac_app where MAC_Address = '"
										+ macAddress + "'");
						JSONObject rdata = json.getJSONObject("Data");
						String pwd = rdata.getString("Data1");
						rows.next();
						String password = rows.getString("Authorization_Password");
						if (pwd.equals(password)) {// 校验密码
							logger.info("right password");
							logger.info("应用数量：" + rowCount);
							String[] appName = new String[rowCount];
							int i = 0;
							rs = stat.executeQuery(
									"select App_Name from view_mac_app where MAC_Address = '" + macAddress + "'");
							while (rs.next()) {
								appName[i] = rs.getString("App_Name");
								i++;
							}
							JSONObject jsonobj = new JSONObject();
							JSONObject appSet = new JSONObject();
							jsonobj.put("Key", "right");
							String data = "";
							String logapps = "";
							appSet.put("long", rowCount);
							for (int j = 0; j < rowCount; j++) {
								data = "data" + j;
								logapps = logapps + appName[j] + ";";
								appSet.put(data, appName[j]);
							}
							logger.info(logapps);
							jsonobj.put("Data", appSet);

							try {
								out = response.getWriter();
								out.print(jsonobj);
							} catch (IOException e) {
								logger.info(e);
							} finally {
								if (out != null) {
									out.close();
								}
							}
						} else {// 密码错误
							logger.info("wrong password");
							JSONObject jsonobj = new JSONObject();
							jsonobj.put("Key", "wrong");
							try {
								out = response.getWriter();
								out.print(jsonobj);
							} catch (IOException e) {
								logger.info(e);
							} finally {
								if (out != null) {
									out.close();
								}
							}
						}
					} else {// 不校验密码，只记录位置
						stat.execute(
								"insert into table_coordinate_record(MAC_address,Coordinate_Longitude,Coordinate_Latitude)values('"
										+ macAddress + "','" + lo + "','" + la + "')");
						logger.info("记录位置:recorded");
						JSONObject jsonobj = new JSONObject();
						jsonobj.put("Key", "Record");
						try {
							out = response.getWriter();
							out.print(jsonobj);
						} catch (IOException e) {
							logger.info(e);
						} finally {
							if (out != null) {
								out.close();
							}
						}
					}
				} else {// 坐标范围错误
					logger.info("wrong position");
				}
			} else {// mac错误，不回复，记录流水mac为0
				logger.info("wrong mac:" + macAddress);
				logMac = "000000000000";
			}
			stat.execute(
					"insert into table_running_record (MAC_address,Data)values('" + logMac + "','" + logData + "')");
		} catch (ClassNotFoundException e) {
			logger.info(e);
		} catch (SQLException e) {
			logger.info(e);
		}
	}

	public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		logger.info("doGet");
		doPost(request, response);
	}

	private boolean chkpos(Double dlo, Double dla) {// TODO 用于限制坐标范围的接口
		return dlo * dla != 0;
	}
}
