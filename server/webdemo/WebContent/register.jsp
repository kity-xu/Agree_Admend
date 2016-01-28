<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<link rel="stylesheet" href="style/background.css" type="text/css">
<head>
<%@ page language="java" contentType="text/html; charset=UTF-8"%>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>注册页面</title>
<script type="text/javascript" src="user.js"></script>
</head>
<body class="backgroundcss">
	<form name="register_form" action="Register" method="post">
		<div class="logo">
			<p>
				<img src=style/20142181183.png>
			</p>
		</div>
		<div class="login" align="center">
			<div class="loginborder" align="right">
				<div class="logincont">
					<label>请输入用户注册信息：</label><br>
					<p style="margin: 0; margin-top: 2px;">
						用户名： <input type="text" id="username" name="username" size="20"
							maxLength="30" style="margin-left: 16px;"
							onkeydown='if(event.keyCode==13){enter.click()}' /> <span
							class="require">*</span>
					</p>
					<p style="margin: 0; margin-top: 2px;">
						密码： <input type="password" id="password" name="password" size="15"
							maxLength="30" style="margin-left: 32px;"
							onkeydown='if(event.keyCode==13){enter.click()}' /><span
							class="require">*</span>
					</p>
					<p style="margin: 0; margin-top: 2px;">
						再次输入： <input type="password" id="repassword" name="repassword"
							size="15" maxLength="30"
							onkeydown='if(event.keyCode==13){enter.click()}' /><span
							class="require">*</span>
					</p>
					<p style="margin: 0; margin-top: 2px;">
						邮箱地址： <input type="text" id="email" name="email" size="20"
							maxLength="30" style="margin-left: 0px;"
							onkeydown='if(event.keyCode==13){enter.click()}' />
					</p>
					<p style="margin: 0; margin-top: 2px;">
						验证码： <input type="text" id="validation_code"
							name="validation_code" size="4" maxLength="30"
							style="margin-left: 16px;"
							onkeydown='if(event.keyCode==13){enter.click()}' /><span
							class="require">*</span> <img id="img_validation_code"
							src="validation_code" />
					</p>
					<input name="提交" type="button" style="margin-top: 1em;"
						onclick="checkRegister()" value="注册" id=enter /> <input
						type="button" value="登录" name="login" style="margin-top: 1em;"
						onclick="javascript:window.location. href='login.jsp'; " /> <input
						type="reset" value="重置" name="reset" style="margin-top: 1em;" /><input
						type="button" value="刷新" onclick="refresh()"
						style="margin-top: 1em; margin-left: 2em;" />
				</div>
			</div>
		</div>
	</form>
</body>
</html>