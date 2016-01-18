/**
 * 
 */
function checkRegister() {
	var username = document.getElementById("username");
	if (username.value == "") {
		alert("必须输入用户名！");
		username.focus();
		return;
	}

	var password = document.getElementById("password");
	if (password.value == "") {
		alert("必须输入密码！");
		password.focus();
		return;
	}

	var repassword = document.getElementById("repassword");
	if (password.value != repassword.value) {
		alert("两次输入的密码不一致！");
		repassword.focus();
		return;
	}

	var email = document.getElementById("email");
	if (email.value != "") {
		if (!checkEmail(email))
			return;
	}

	var validation_code = document.getElementById("validation_code");
	if (validation_code.value == "") {
		alert("必须输入验证码！");
		validation_code.focus();
		return;
	}
	register_form.submit();
}

function checkEmail(email) {
	var email2 = email;
	var email = email2.value;
	var pattern = /^([a-zA-Z0-9._-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/;
	flag = pattern.test(email);
	if (!flag) {
		alert("eamil格式不正确！");
		email.focus();
		return flase;
	}
	return true;
}

function refresh() {
	var img = document.getElementById("img_validation_code");
	img.src = "validation_code?" + Math.random();
}

function checkLogin() {
	var username = document.getElementById("username");
	if (username.value == "") {
		alert("必须输入用户名！");
		username.focus();
		return;
	}

	var password = document.getElementById("password");
	if (password.value == "") {
		alert("必须输入密码！");
		password.focus();
		return;
	}

	var validation_code = document.getElementById("validation_code");
	if (validation_code.value == "") {
		alert("必须输入验证码！");
		validation_code.focus();
		return;
	}
	login_form.submit();
}
