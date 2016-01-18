<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@ page import="bean.CurrentMac"%>
<%@ page import="bean.Page"%>
<%@ page import="dbservlet.CurrentMacServlet"%>
<%
	String path = "/webdemo";
	String basePath = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort()
			+ path + "/";
%>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 4.01 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
<link rel="stylesheet" href="/webdemo/style/page.css" type="text/css">
<head>
<base href="<%=basePath%>">
<title>当前MAC设备信息</title>
<meta http-equiv="pragma" content="no-cache">
<meta http-equiv="cache-control" content="no-cache">
<meta http-equiv="expires" content="0">
<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
<meta http-equiv="description" content="This is my page">
<link rel="stylesheet" href="/webdemo/style/page.css" type="text/css">
<style type="text/css">
h2 {
	margin: 0;
}
</style>
</head>
<body>
	<div class="container">
		<div id="menu">
			<br>
			<h2>
				当前MAC设备列表 <input type="button" value="显示全部位置"
					style="margin-left: 610px;"
					onClick="window.open('/webdemo/current_MAC/ShowAllDevices.jsp', '_blank')" />
				<input type="button" value="刷新" style="margin-left: 1px"
					onclick="javascript:window.location.href='/webdemo/current_MAC/CurrentMac.jsp'" />
			</h2>
			<hr>
			<table style="margin-top: 3px;">
				<thead>
					<tr>
						<th width="200">MAC地址</th>
						<th width="200">上次定位时间</th>
						<th width="30">操作</th>
					</tr>
				</thead>
				<tbody>
					<%
						response.setCharacterEncoding("UTF-8");
						request.setCharacterEncoding("UTF-8");
						CurrentMacServlet servlet = new CurrentMacServlet();
						ArrayList<CurrentMac> result = servlet.selectCurrentMac();
						Page pager = new Page();
						pager = servlet.setPage(request, response);
						List<CurrentMac> subResult = null;
						int currentRecord = pager.getCurrentRecord();
						CurrentMac mac1 = new CurrentMac();
						if (currentRecord == 0) {
							if (pager.getTotalRecord() < 7) {
								subResult = (List<CurrentMac>) result.subList(0, pager.getTotalRecord());
							} else {
								subResult = (List<CurrentMac>) result.subList(0, pager.getPageSize());
							}
						} else if (pager.getCurrentRecord() + pager.getPageSize() < result.size()) {
							subResult = (List<CurrentMac>) result.subList(pager.getCurrentRecord(),
									pager.getCurrentRecord() + pager.getPageSize());
						} else {
							subResult = (List<CurrentMac>) result.subList(pager.getCurrentRecord(), result.size());
						}
						if (!subResult.isEmpty()) {
							for (int i = 0; i < subResult.size(); i++) {
								CurrentMac st = subResult.get(i);
					%>
					<tr>
						<%
							out.print("<td width=" + '"' + "300" + '"' + ">" + st.getMacAddress() + "</td>");
									out.print("<td width=" + '"' + "300" + '"' + ">" + st.getDateAdd() + "</td>");
						%>
						<td width="20"><a
							href="current_MAC/ShowPath.jsp?MAC_address=<%=st.getMacAddress()%>"
							target="_blank" title="查看路径"> <img src="style/img/path.png"
								alt="查看路径" /></a></td>
					</tr>
				</tbody>
				<%
					}
					}
					try {
				%>
				<tfoot>
					<tr>
						<td colspan="6">
							<div class="pagination">
								总<%=pager.getTotalRecord()%>条记录·总<%=pager.getTotalPage()%>页·当前<%=pager.getCurrentPage() + 1%>页·每页<%=pager.getPageSize()%>条&nbsp;
								<%
									int last = pager.getCurrentRecord() - pager.getPageSize();
										int next = pager.getCurrentRecord() + pager.getPageSize();
										//	int currentRecord;
										if (last < 0) {
											out.println(" 首页  ");
										} else {
											out.print("<a href='AppServlet?currentRecord=" + last + "&methodName=1'><<上一页</a>");
										}
										if (next >= pager.getTotalRecord()) {
											out.println(" 尾页  ");
										} else {
											out.print("<a href='AppServlet?currentRecord=" + next + "&methodName=1'>下一页>></a>");
										}
									} catch (

									Exception e)

									{
										e.getMessage();
									}
								%>
							</div>
						</td>
					</tr>

				</tfoot>
			</table>
		</div>
	</div>
</body>
</html>