<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<link rel="stylesheet" href="style/background.css" type="text/css">
<head>
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@ page import="bean.CurrentMac"%>
<%@ page import="bean.Page"%>
<%@ page import="webdemo.Login"%>
<%
	Login login = new Login();
	login.logout(request, response);
%>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>登录页面</title>
<script type="text/javascript" src="user.js"></script>
</head>
<body class="backgroundcss">
	<form name="login_form" action="login" method="post">
		<div class="logo">
			<p>
				<img src=style/20142181183.png>
			</p>
		</div>
		<div class="login" align="center">
			<div class="loginborder" align="right">
				<div class="logincont">
					<label>赞同桌面后台配置</label>
					<p style="margin: 0;">用户名：</p>
					<input type="text" id="username" value="${requestScope.username}"
						class="input_list" name="user_name" size="30" maxLength="30"
						style="width: 16em;" />
					<p style="margin: 0;">密 码：</p>
					<input type="password" id="password" class="input_list"
						name="user_password" size="30" maxLength="30" style="width: 16em;" />
					<p style="margin: 0;">验证码：</p>
					<input type="text" id="validation_code" name="validation_code"
						style="margin-left: 0;" size="4" maxlength="4" /> <img
						id="img_validation_code" src="validation_code" /> <input
						type="button" style="margin-left: 0px;" value="刷新"
						onclick="refresh()" /> <br> <input type="submit"
						style="margin-top: 1em;" value="登录" name="login"
						onclick="checkLogin()" /> <input type="button"
						style="margin-top: 1em;" value="注册" name="register"
						onclick="javascript:window.location. href='register.jsp'; " /> <input
						type="reset" style="margin-top: 1em;" value="重置"> <br>
					<font color="#FF0000">${requestScope.userError}</font> <font
						color="#FF0000">${requestScope.passwordError}</font> <font
						color="#FF0000">${requestScope.codeError}</font>
				</div>
			</div>
		</div>
	</form>
</body>
</html>