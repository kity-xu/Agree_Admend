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

<title>删除界面</title>

<meta http-equiv="pragma" content="no-cache">
<meta http-equiv="cache-control" content="no-cache">
<meta http-equiv="expires" content="0">
<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
<meta http-equiv="description" content="This is my page">
</head>

<body>
	<%
		out.print("<center><br><br><h3>删除成功！</h3></center>");
	%>
	<br>
	<br>
	<center>
		<input type="button" value="返回"
			onclick="javascript:window.location.href='MacServlet?methodName=<%=1%>&macAddress=<%=""%>&macName=<%=""%>';" />
		<input type="button" value="添加"
			onclick="javascript:window.location.href='dbservlet/input.jsp';"
			style="margin-left: 50px" />
	</center>
</body>
</html>

