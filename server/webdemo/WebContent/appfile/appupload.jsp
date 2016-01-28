<%@ page language="java" pageEncoding="UTF-8"%>
<%
	String path = "/webdemo";
	String basePath = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort()
			+ path + "/";
	request.getSession().setAttribute("IS_UPLOAD_BEGIN", 0);
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
<script type="text/javascript"
	src="http://ajax.microsoft.com/ajax/jquery
/jquery-1.4.min.js"></script>
<script type="text/javascript">
	function validate() {
		//应用权限禁止为000，不可用应用权限为999
		var file_var = document.forms[0].file.value;
		if (file_var.length <= 0) {
			alert("请上传文件");
			return false;
		} else {
			return true;
		}
	}
	function beginUpload() {
		$("#progress_bar").show();
		setInterval("getUploadMeter()", 200);
	}
	function getUploadMeter() {
		$.post("/webdemo/ProgressBarServlet", function(data) {
			var json = eval("(" + data + ")");
			jQuery("#msg").html(json.percentage + "%已上传，文件上传完成后需要处理，请耐心等待跳转");
		});
	}
</script>
</head>
<body style="overflow: hidden;">
	<br>
	<center>
		<h2>上传app</h2>
		<hr>
		<form action="UploadServlet" method="post"
			enctype="multipart/form-data" onSubmit="return validate()">
			<h4>
				<input type="file" name="file" size="50" accept=".apk" /> <br /> <input
					type="submit" value="Upload File" onclick="beginUpload()" /><input
					type="button" value="返回" style="margin-left: 30px"
					onclick="javascript:window.location.href='/webdemo/AppServlet?methodName=<%=1%>&appName=<%=""%>&appPath=<%=""%>';" />
			</h4>
		</form>
		<div id="msg"></div>
	</center>
</body>
</html>
