<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@ page import="bean.PermissionPassword"%>
<%
	String path = "/webdemo";
	String basePath = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort()
			+ path + "/";
%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
<base href="<%=basePath%>">

<title>按权限标号和权限描述查询</title>

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
						<th>权限编号</th>
						<th>权限描述</th>
						<th>密码</th>
						<th>操作</th>
					</tr>
				</thead>
				<tbody>
					<%
						response.setCharacterEncoding("UTF-8");
						request.setCharacterEncoding("UTF-8");

						try {
							@SuppressWarnings("unchecked")
							ArrayList<PermissionPassword> result = (ArrayList<PermissionPassword>) request.getAttribute("result");
							if (!result.isEmpty()) {
								for (int i = 0; i < result.size(); i++) {
					%>
					<tr>
						<%
							PermissionPassword st = result.get(i);
										out.print("<td width=" + '"' + "200" + '"' + ">" + st.getPermissionId() + "</td>");
										out.print("<td width=" + '"' + "200" + '"' + "><div style=\"word-wrap:break-word;\">"
												+ st.getPermissionInfo() + "</div></td>");
										out.print("<td width=" + '"' + "200" + '"' + ">" + st.getPassword() + "</td>");
						%>
						<td width="40"><a
							href="PerpasswordServlet?permissionID=<%=st.getPermissionId()%>&&methodName=<%=4%>"
							title="修改"> <img src="style/img/pencil.png" alt="修改" /></a></td>
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
									onclick="javascript:window.location.href='/webdemo/PerpasswordServlet?methodName=<%=1%>&&permissionID=<%=""%>&&methodName=<%=""%>';" />
							</div>
						</td>
					</tr>
				</tfoot>
			</table>
		</div>
	</div>
</body>
</html>
