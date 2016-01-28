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

<title>App应用权限配置修改</title>

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
		//应用权限禁止为000，不可用应用权限为999
		var appName = document.forms[0].appName.value;
		var permission = document.forms[0].appPermission.value;
		if (MAC_name.length <= 0) {
			alert("App原文件名不能为空");
			return false;
		} else if (permission.length <= 0) {
			alert("应用权限不能为空，请输入应用权限！");
			return false;
		} else if (app_authorization == 000) {
			alert("应用权限不能为000");
			return false;
		} else if (app_authorization == 999) {
			alert("提示：999权限表示要禁用此应用");
			return true;
		} else {
			return true;
		}
	}
</script>

</head>

<body>
	<div class="container">
		<br>
		<h2>App应用权限配置</h2>
		<hr>
		<div id="menu" style="height: 250px;">
			<table style="width: 496;">
				<thead>
					<tr>
						<th>App名称</th>
						<th>App下载地址</th>
						<th>应用权限</th>
					</tr>
				</thead>
				<%
					String appName = null;
					// ArrayList<StudentInfo> result=new ArrayList<StudentInfo>();
					try {
						@SuppressWarnings("unchecked")
						ArrayList<AppPermission> result = (ArrayList<AppPermission>) request.getAttribute("result");
						if (!result.isEmpty()) {
							for (int i = 0; i < result.size(); i++) {
								AppPermission st = result.get(i);
								appName = st.getAppName();
								out.print("<tr>");
								out.print("<td width=" + '"' + "200" + '"' + ">" + st.getAppName() + "</td>");
								out.print("<td width=" + '"' + "200" + '"' + ">" + st.getAppPath() + "</td>");
								out.print("<td width=" + '"' + "40" + '"' + ">" + st.getPermission() + "</td>");
								out.print("</tr>");
							}
						}
					} catch (Exception e) {
					}
				%>
			</table>
			<div class="update">
				<form action="AppServlet" method="post" onSubmit="return validate()">
					<input type="hidden" name="methodName" value="3" /> <br>

					&nbsp;App名称：<input type="text" name="appName" readonly
						value="<%=appName%>" title="App名称不能改变"></input><br> <br>
					应用信息：<input type="text" name="appPath" title="App应用信息不能为空"></input><br>
					<br> 应用权限：<input type="text" name="appPermission"
						title="应用权限不能为空"></input><br> <br>
					<div style="margin-left: 190px;">
						<input type="submit" value="修改" />
					</div>
				</form>
				<br> <input type="button" value="返回信息添加页面"
					style='font-size: 16px'
					onclick="javascript:window.location.href='/webdemo/appfile/appupload.jsp';" />
				&nbsp; <input type="button" value="返回信息查询页面" style='font-size: 16px'
					onclick="javascript:window.location.href='/webdemo/AppServlet?methodName=<%=1%>&appName=<%=""%>&appPath=<%=""%>';" />
			</div>
		</div>
	</div>
</body>
</html>
