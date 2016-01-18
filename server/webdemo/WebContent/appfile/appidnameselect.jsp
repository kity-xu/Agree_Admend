<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@ page import="bean.AppPermission"%>
<%
	String path = "/webdemo";
	String basePath = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort()
			+ path + "/";
%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
<base href="<%=basePath%>">

<title>按App应用名称和App下载地址查询</title>

<meta http-equiv="pragma" content="no-cache">
<meta http-equiv="cache-control" content="no-cache">
<meta http-equiv="expires" content="0">
<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
<meta http-equiv="description" content="This is my page">
<link rel="stylesheet" href="/webdemo/style/page.css" type="text/css">

</head>

<body>
	<div class="container">
		<div id="menu">
			<table style="margin-top: 3px;">
				<thead>
					<tr>
						<th>App名称</th>
						<th>App原文件名</th>
						<th>应用权限</th>
						<th>操作</th>
					</tr>
				</thead>
				<tbody>
					<%
						response.setCharacterEncoding("UTF-8");
						request.setCharacterEncoding("UTF-8");

						try {
							@SuppressWarnings("unchecked")
							ArrayList<AppPermission> result = (ArrayList<AppPermission>) request.getAttribute("result");
							if (!result.isEmpty()) {
								for (int i = 0; i < result.size(); i++) {
					%>
					<tr>
						<%
							AppPermission st = result.get(i);
										out.print("<td width=" + '"' + "200" + '"' + ">" + st.getAppName() + "</td>");
										out.print("<td width=" + '"' + "200" + '"' + "><div style=\"word-wrap:break-word;\">"
												+ st.getAppPath() + "</div></td>");
										out.print("<td width=" + '"' + "200" + '"' + ">" + st.getPermission() + "</td>");
						%>
						<td width="40"><a
							href="AppServlet?appName=<%=st.getAppName()%>&&methodName=<%=4%>"
							title="修改"> <img src="style/img/pencil.png" alt="修改" /></a> <a
							href="AppServlet?appName=<%=st.getAppName()%>&&methodName=<%=2%>"
							title="删除"> <img src="style/img/cross.png" alt="删除" /></a></td>
					</tr>
					<%
						}
							}
						} catch (Exception e) {

						}
					%>
				</tbody>
				<tfoot>
					<tr>
						<td colspan="6">
							<div class="pagination2">
								<br> <input type="button" value="返回信息查询页面"
									style='font-size: 18px'
									onclick="javascript:window.location.href='/webdemo/AppServlet?methodName=<%=1%>&&appName=<%=""%>&&methodName=<%=""%>';" />
							</div>
						</td>
					</tr>
				</tfoot>
			</table>
		</div>
	</div>
</body>
</html>
