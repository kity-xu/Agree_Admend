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

<title>查无结果</title>

<meta http-equiv="pragma" content="no-cache">
<meta http-equiv="cache-control" content="no-cache">
<meta http-equiv="expires" content="0">
<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
<meta http-equiv="description" content="This is my page">

</head>

<body>
	<%
		out.print("<center><br><br><h3>没有符合此条件的设备！</h3></center>");
	%>
	<br>
	<br>
	<center>
		<input type="button" value="返回信息添加页面" style='font-size: 16px'
			onclick="javascript:window.location.href='/webdemo/dbservlet/input.jsp';" />
		<input type="button" value="返回信息查询页面" style='font-size: 16px'
			onclick="javascript:window.location.href='/webdemo/MacServlet?methodName=<%=1%>&macAddress=<%=""%>&macName=<%=""%>';" />
	</center>
</body>
</html>
