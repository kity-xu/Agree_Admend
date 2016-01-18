<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@ page import="bean.MacPermission"%>
<%
	String path = "/webdemo";
	String basePath = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort()
			+ path + "/";
%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
<base href="<%=basePath%>">

<title>可信MAC地址配置修改</title>

<meta http-equiv="pragma" content="no-cache">
<meta http-equiv="cache-control" content="no-cache">
<meta http-equiv="expires" content="0">
<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
<meta http-equiv="description" content="This is my page">
<link rel="stylesheet" href="/webdemo/style/page.css" type="text/css">

<style type="text/css">
h2 {
	margin: 0;
}
</style>

<script type="text/javascript">
	function validate() {
		//设备权限禁止为999，不可用设备权限为000
		var macName = document.forms[0].macName.value;
		var permission = document.forms[0].macPermission.value;
		if (macName.length <= 0) {
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
	<div class="container">
		<br>
		<h2>可信MAC地址配置</h2>
		<hr>
		<div id="menu" style="height: 250px;">
			<table style="width: 496;">
				<thead>
					<tr>
						<th>MAC地址</th>
						<th>设备名称</th>
						<th>权限设置</th>
					</tr>
				</thead>
				<%
					String macAddress = null;
					// ArrayList<StudentInfo> result=new ArrayList<StudentInfo>();
					try {
						@SuppressWarnings("unchecked")
						ArrayList<MacPermission> result = (ArrayList<MacPermission>) request.getAttribute("result");
						if (!result.isEmpty()) {
							for (int i = 0; i < result.size(); i++) {
								MacPermission st = result.get(i);
								macAddress = st.getMacAddress();
								out.print("<tr>");
								out.print("<td width=" + '"' + "200" + '"' + ">" + st.getMacAddress() + "</td>");
								out.print("<td width=" + '"' + "200" + '"' + ">" + st.getMacName() + "</td>");
								out.print("<td width=" + '"' + "40" + '"' + ">" + st.getPermission() + "</td>");
								out.print("</tr>");
							}
						}
					} catch (Exception e) {
					}
				%>
			</table>
			<div class="update">
				<div style="font-weight: bold; font-size: 18px;">将MAC地址配置更改为：</div>
				<form action="MacServlet" method="post" onSubmit="return validate()">
					<input type="hidden" name="methodName" value="3" /> <br>
					&nbsp;MAC地址：<input type="text" name="macAddress" readonly
						value="<%=macAddress%>" title="MAC地址不能改变"></input><br> <br>
					设备名称：<input type="text" name="macName" title="设备名称不能为空"></input><br>
					<br> 权限设置：<input type="text" name="macPermission"
						title="设备权限不能为空"></input><br> <br>
					<div style="margin-left: 190px;">
						<input type="submit" value="修改" />
					</div>
				</form>
				<br> <input type="button" value="返回信息添加页面"
					style='font-size: 14px'
					onclick="javascript:window.location.href='/webdemo/dbservlet/input.jsp';" />
				&nbsp; <input type="button" value="返回信息查询页面" style='font-size: 14px'
					onclick="javascript:window.location.href='/webdemo/MacServlet?methodName=<%=1%>&macAddress=<%=""%>&macName=<%=""%>';" />
			</div>
		</div>
	</div>
</body>
</html>
