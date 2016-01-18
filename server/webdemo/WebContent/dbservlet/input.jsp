
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
<title>添加可信MAC地址配置</title>
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
		//设备权限禁止为999，不可用设备权限为000
		var macAddress = document.forms[0].macAddress.value;
		var macName = document.forms[0].macName.value;
		var permission = document.forms[0].macPermission.value;
		if (macAddress.length <= 0) {
			alert("MAC地址不能为空，请输入MAC地址！");
			return false;
		} else if (macName.length <= 0) {
			alert("设备名称不能为空，请输入设备名称！");
			return false;
		} else if (permission.length <= 0) {
			alert("设备权限不能为空，请输入设备权限！");
			return false;
		} else if (permission == 999) {
			alert("设备权限不能为999");
			return false;
		} else if (permission == 000) {
			alert("提示：000权限表示要禁用此设备");
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
		<h2>添加可信MAC地址配置</h2>
		<hr>
		<form action="MacServlet" method="post" id="form"
			onSubmit="return validate()">

			<input type="hidden" name="methodName" value="0" />
			<h4 id="MAC">
				<b> MAC地址：</b><input type="text" name="macAddress" title="MAC地址不能为空"
					onkeyup="this.value=this.value.replace(/[^0-9A-F]/g,'')"
					onafterpaste="this.value=this.value.replace(/[^0-9A-F]/g,'')"
					maxLength="12"></input> （12位16进制数字,字母大写）<font color="#FF0000">${requestScope.userError}</font><br>
			</h4>
			<h4>
				设备名称：<input type="text" name="macName" title="设备名称不能为空"
					maxLength="50" />（长度限制50）<br>
			</h4>
			<h4>
				权限设置：<input type="text" name="macPermission" title="设备权限不能为空"
					onkeyup="this.value=this.value.replace(/[^0-9]/g,'')"
					onafterpaste="this.value=this.value.replace(/[^0-9]/g,'')" size="3"
					maxLength="3" />（三位数字）<br>
			</h4>
			<input type="submit" value="提交" style="margin-left: 30px" /> <input
				type="button" value="返回" style="margin-left: 30px"
				onclick="javascript:window.location.href='/webdemo/MacServlet?methodName=<%=1%>&macAddress=<%=""%>&macName=<%=""%>';" />
		</form>
		<br>
	</center>
</body>
</html>
