<%@ page language="java" pageEncoding="UTF-8"%>
<%
	String path = "/webdemo";
	String basePath = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort()
			+ path + "/";
%>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 4.01 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
<link rel="stylesheet" type="text/css" href="style/main.css" />
<head>
<base href="<%=basePath%>">
<title>赞同桌面后台配置</title>
<meta http-equiv="pragma" content="no-cache">
<meta http-equiv="cache-control" content="no-cache">
<meta http-equiv="expires" content="0">
<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
<meta http-equiv="description" content="This is my page">
</head>
<body>
	<div id="container_main">
		<div id="header">
			<h1>赞同桌面后台配置</h1>
		</div>
		<div id="link">
			<ul>
				<li><a href="dbservlet/layout.jsp" target="display">可信MAC地址配置</a></li>
				<li><a href="perpasswordfile/perlayout.jsp" target="display">权限密码配置</a></li>
				<li><a href="appfile/applayout.jsp" target="display">App应用权限配置</a></li>
				<li><a href="current_MAC/CurrentMac.jsp" target="display">当前设备信息</a></li>
				<li><a href="login.jsp">退出</a></li>
			</ul>
		</div>
		<iframe name="display" src="dbservlet/layout.jsp" width=1000px
			height="400px" marginheight="0px" frameborder="0"
			style="background-image: url(style/img/manger.png); background-repeat: no-repeat; background-size: 100% 100%;">
		</iframe>
		<div id="footer">
			<b>welcome</b>
		</div>
	</div>
</body>
</html>
