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

<title>提交成功</title>

<meta http-equiv="pragma" content="no-cache">
<meta http-equiv="cache-control" content="no-cache">
<meta http-equiv="expires" content="0">
<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
<meta http-equiv="description" content="This is my page">

</head>

<body>
	<center>
		<br> <br>
		<h3>提交成功！</h3>
		<input type="button" value="返回"
			onclick="javascript:window.location.href='perpasswordfile/perinput.jsp';" />
		<input type="button" value="查询"
			onclick="javascript:window.location.href='PerpasswordServlet?methodName=<%=1%>&permissionID=<%=""%>&permissionInfo=<%=""%>';"
			style="margin-left: 50px" />
	</center>
</body>
</html>
