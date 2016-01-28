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

<title>权限编号密码配置修改</title>

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

<script type="text/javascript">
	function validate() {
		//禁止修改999权限和000权限
		var permissionInfo = document.forms[0].permissionInfo.value;
		var password = document.forms[0].permissionPassword.value;
		if (permissionInfo.length <= 0) {
			alert("权限描述不能为空，请输入权限描述！");
			return false;
		} else if (password.length <= 0) {
			alert("权限密码不能为空，请输入权限密码！");
			return false;
		} else if (permission == 999) {
			alert("999权限为管理用保留权限");
			return false;
		} else if (permission == 000) {
			alert("000权限为管理用保留权限");
			return false;
		} else {
			return true;
		}
	}
</script>

</head>

<body>
	<div class="container">
		<br>
		<h2>权限编号密码配置修改</h2>
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
					String permissionID = null;
					// ArrayList<StudentInfo> result=new ArrayList<StudentInfo>();
					try {
						@SuppressWarnings("unchecked")
						ArrayList<PermissionPassword> result = (ArrayList<PermissionPassword>) request.getAttribute("result");
						if (!result.isEmpty()) {
							for (int i = 0; i < result.size(); i++) {
								PermissionPassword st = result.get(i);
								permissionID = st.getPermissionId();
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
			<div class="update">
				<div style="font-weight: bold; font-size: 18px;">将权限密码配置更改为：</div>
				<form action="PerpasswordServlet" method="post"
					onSubmit="return validate()">
					<input type="hidden" name="methodName" value="3" /> <br>
					权限编号：<input type="text" name="permissionID" readonly
						value="<%=permissionID%>" title="权限编号不能改变"></input><br>
					<br> 权限描述：<input type="text" name="permissionInfo"
						title="权限描述不能为空"></input><br>
					<br> &nbsp;&nbsp;&nbsp;&nbsp;密码：<input type="text"
						name="permissionPassword" title="密码不能为空"></input><br>
					<br>
					<div style="margin-left: 190px;">
						<input type="submit" value="修改" />
					</div>
				</form>
				<br> <input type="button" value="返回信息添加页面"
					style='font-size: 16px'
					onclick="javascript:window.location.href='/webdemo/perpasswordfile/perinput.jsp';" />
				&nbsp; <input type="button" value="返回信息查询页面" style='font-size: 16px'
					onclick="javascript:window.location.href='/webdemo/PerpasswordServlet?methodName=<%=1%>&permissionID=<%=""%>&permissionInfo=<%=""%>';" />
			</div>
		</div>
	</div>
</body>
</html>
