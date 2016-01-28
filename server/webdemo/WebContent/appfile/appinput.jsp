<%@ page language="java" pageEncoding="UTF-8"%>
<%
	String path = "/webdemo";
	String basePath = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort()
			+ path + "/";
%>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
<base href="<%=basePath%>">
<title>添加App应用权限配置</title>
<meta http-equiv="pragma" content="no-cache">
<meta http-equiv="cache-control" content="no-cache">
<meta http-equiv="expires" content="0">
<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
<meta http-equiv="description" content="This is my page">
<style type="text/css">
div#MAC {
	width: 500px;
	margin: 0;
	text-align: left;
	padding-left: 250px;
}
</style>
<script type="text/javascript">
	function validate() {
		//应用权限禁止为000，不可用应用权限为999
		var app_name = document.forms[0].appName.value;
		var app_path = document.forms[0].appPath.value;
		var app_authorization = document.forms[0].appPermission.value;
		if (app_name.length <= 0) {
			alert("App名称不能为空，请输入App名称！");
			return false;
		} else if (app_path.length <= 0) {
			alert("App原文件名不能为空");
			return false;
		} else if (app_authorization.length <= 0) {
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
	<br>
	<center>
		<h2>添加App应用权限配置</h2>
		<hr>
		<form action="AppServlet" method="post" id="form"
			onSubmit="return validate()">


			<input type="hidden" name="methodName" value="0" />
			<h4 id="APP">
				<b> App名称：</b><input type="text" name="appName" title="App名称"
					maxLength="30"></input> （长度限制30）<%
					//TODO 汉字会导致部分三星设备无法读取，原因不明
				%><font color="#FF0000">${requestScope.userError}</font><br>
			</h4>
			<h4 id="APP">
				<b> 文件名称：</b><input type="text" readonly="readonly" name="appPath" title="文件名称"
					maxLength="255" value="<%=request.getAttribute("filename")%>" ></input>
				（长度限制255）<br>
			</h4>
			<h4>
				应用权限：<input type="text" name="appPermission" title="应用权限"
					onkeyup="this.value=this.value.replace(/[^0-9]/g,'')"
					onafterpaste="this.value=this.value.replace(/[^0-9]/g,'')" size="3"
					maxLength="3" />（三位数字）<br>
			</h4>
			<input type="submit" value="提交" style="margin-left: 30px" /> <input
				type="button" value="返回" style="margin-left: 30px"
				onclick="javascript:window.location.href='/webdemo/AppServlet?methodName=<%=1%>&appName=<%=""%>&appPath=<%=""%>';" />
		</form>
	</center>
</body>
</html>
