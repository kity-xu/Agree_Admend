<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
<%@ page language="java" contentType="text/html; charset=UTF-8"%>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c"%>
<title>结果页面</title>
</head>
<body>
	<form name="form" action="${requestScope.page}" method="post"></form>
	<c:if test="${requestScope.info != null}">
		<script type="text/javascript">
			alert('${requestScope.info}');
			form.submit();
		</script>
	</c:if>
</body>
</html>