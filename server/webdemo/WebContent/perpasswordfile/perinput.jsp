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
<title>添加权限密码配置</title>
<meta http-equiv="pragma" content="no-cache">
<meta http-equiv="cache-control" content="no-cache">
<meta http-equiv="expires" content="0">
<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
<meta http-equiv="description" content="This is my page">
<link rel="stylesheet" type="text/css" href="styles.css">
<script type="text/javascript">
	function validate() {
		//禁止修改999权限和000权限
		var permissionID = document.forms[0].permissionID.value;
		var permissionInfo = document.forms[0].permissionInfo.value;
		var password = document.forms[0].permissionPassword.value;
		if (permissionID.length <= 0) {
			alert("权限编号不能为空，请输入权限编号！");
			return false;
		} else if (permissionInfo.length <= 0) {
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
	<br>
	<center>
		<h2>添加权限密码配置</h2>
		<hr>
		<form action="PerpasswordServlet" method="post" id="form"
			onSubmit="return validate()">

			<input type="hidden" name="methodName" value="0" />
			<h4 id="MAC">
				<b> 权限编号：</b><input type="text" name="permissionID" title="权限编号不能为空"
					onkeyup="this.value=this.value.replace(/[^0-9]/g,'')"
					onafterpaste="this.value=this.value.replace(/[^0-9]/g,'')" size="3"
					maxLength="3" />（三位数字） <font color="#FF0000">${requestScope.userError}</font><br>
			</h4>
			<h4>
				权限描述： <input name="permissionInfo" title="权限描述不能为空" maxLength="255" />（长度限制255）<br>
			</h4>
			<h4>
				密码：<input type="text" name="permissionPassword" title="密码不能为空"
					maxLength="8" />（八位数字）<br>
			</h4>
			<input type="submit" value="提交" style="margin-left: 30px" /> <input
				type="button" value="返回" style="margin-left: 30px"
				onclick="javascript:window.location.href='/webdemo/PerpasswordServlet?methodName=<%=1%>&permissionID=<%=""%>&permissionInfo=<%=""%>';" />
		</form>
		<br>
	</center>
</body>
</html>
