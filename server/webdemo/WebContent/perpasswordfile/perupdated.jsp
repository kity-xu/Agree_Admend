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

<title>修改成功</title>

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
		<br>
		<h2>修改后的权限密码配置为：</h2>
		<hr>
		<div id="menu" style="height: 250px;">
			<table style="width: 496;">
				<thead>
					<tr>
						<th>权限编号</th>
						<th>权限描述</th>
						<th>密码</th>
					</tr>
				</thead>
				<%
					response.setCharacterEncoding("UTF-8");
					request.setCharacterEncoding("UTF-8");
					//  ArrayList<StudentInfo> result=new ArrayList<StudentInfo>();
					try {
						@SuppressWarnings("unchecked")
						ArrayList<PermissionPassword> result = (ArrayList<PermissionPassword>) request.getAttribute("result");
						if (!result.isEmpty()) {
							for (int i = 0; i < result.size(); i++) {
								PermissionPassword st = result.get(i);
								out.print("<tr>");
								out.print("<td width=" + '"' + "200" + '"' + ">" + st.getPermissionId() + "</td>");
								out.print("<td width=" + '"' + "200" + '"' + ">" + st.getPermissionInfo() + "</td>");
								out.print("<td width=" + '"' + "40" + '"' + ">" + st.getPassword() + "</td>");
								out.print("</tr>");
							}
						}
					} catch (Exception e) {

					}
				%>
			</table>
			<div
				style="position: absolute; margin-top: 180px; margin-left: 630px;">
				<input type="button" value="返回信息添加页面" style='font-size: 16px'
					onclick="javascript:window.location.href='/webdemo/perpasswordfile/perinput.jsp';" />
				&nbsp; <input type="button" value="返回信息查询页面" style='font-size: 16px'
					onclick="javascript:window.location.href='/webdemo/PerpasswordServlet?methodName=<%=1%>&permissionID=<%=""%>&permissionInfo=<%=""%>';" />
			</div>
		</div>
	</div>
</body>
</html>
