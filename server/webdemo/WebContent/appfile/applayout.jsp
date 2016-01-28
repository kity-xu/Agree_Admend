<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@ page import="bean.AppPermission"%>
<%@ page import="bean.Page"%>
<%@ page import="dbservlet.AppServlet"%>
<%
	String path = "/webdemo";
	String basePath = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort()
			+ path + "/";
%>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 4.01 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
<link rel="stylesheet" type="text/css" href="/webdemo/style/page.css" />
<head>
<base href="<%=basePath%>">
<title>App应用权限配置区</title>
<meta http-equiv="pragma" content="no-cache">
<meta http-equiv="cache-control" content="no-cache">
<meta http-equiv="expires" content="0">
<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
<meta http-equiv="description" content="This is my page">
<link rel="stylesheet" href="/webdemo/style/page.css" type="text/css">
<script type="text/javascript">
	function confirmdialog() {
		if (window.confirm("您确定要删除此条信息？")) {
			return true;
		} else {
			//  alert("取消删除！");
			return false;
		}
	}
	function validate() {
		var app_name = document.forms[0].appName.value;
		var app_path = document.forms[0].appPath.value;
		if (app_name.length <= 0 && app_path.length <= 0) {
			alert("请输入搜索条件");
			return false;
		} else {
			return true;
		}
	}
</script>
</head>
<body>
	<div class="container">
		<form action="AppServlet" method="post" onSubmit="return validate()"
			style="margin-left: 25px">
			<input type="hidden" name="methodName" value="5" /> 名称： <input
				type="text" name="appName" value="" title=""></input> 原文件名： <input
				type="text" name="appPath" value="" title=""></input> <input
				type="submit" value="查询" /> <input type="button" value="添加"
				style="margin-left: 1px"
				onClick="javascript:window.location.href='/webdemo/appfile/appupload.jsp';" />
			<input type="button" value="刷新" style="margin-left: 1px"
				onclick="javascript:window.location.href='/webdemo/AppServlet?methodName=<%=1%>&appName=<%=""%>&appPath=<%=""%>';" />
		</form>
		<div id="menu">
			<table>
				<thead>
					<tr>
						<th>App名称</th>
						<th>原文件名</th>
						<th>应用权限</th>
						<th>操作</th>
					</tr>
				</thead>
				<tbody>
					<%
						response.setCharacterEncoding("UTF-8");
						request.setCharacterEncoding("UTF-8");
						AppServlet servlet = new AppServlet();
						ArrayList<AppPermission> result = servlet.select("", "");
						Page pager = new Page();
						pager = servlet.setPage(request, response);
						List<AppPermission> subResult = null;
						int currentRecord = pager.getCurrentRecord();
						if (currentRecord == 0) {
							if (pager.getTotalRecord() < 7) {
								subResult = (List<AppPermission>) result.subList(0, pager.getTotalRecord());
							} else {
								subResult = (List<AppPermission>) result.subList(0, pager.getPageSize());
							}
						} else if (pager.getCurrentRecord() + pager.getPageSize() < result.size()) {
							subResult = (List<AppPermission>) result.subList(pager.getCurrentRecord(),
									pager.getCurrentRecord() + pager.getPageSize());
						} else {
							subResult = (List<AppPermission>) result.subList(pager.getCurrentRecord(), result.size());
						}
						if (!subResult.isEmpty()) {
							for (int i = 0; i < subResult.size(); i++) {
								AppPermission st = subResult.get(i);
					%>
					<tr>
						<%
							out.print("<td width=" + '"' + "200" + '"' + ">" + st.getAppName() + "</td>");
									out.print("<td width=" + '"' + "200" + '"' + ">" + st.getAppPath() + "</td>");
									out.print("<td width=" + '"' + "200" + '"' + ">" + st.getPermission() + "</td>");
						%>
						<td width="30"><a
							href="AppServlet?appName=<%=st.getAppName()%>&appPath=<%=""%>&methodName=<%=4%> "
							title="修改"> <img src="style/img/pencil.png" alt="修改" /></a> <a
							href="AppServlet?appName=<%=st.getAppName()%>&appPath=<%=""%>&methodName=<%=2%> "
							title="删除"> <img src="style/img/cross.png" alt="删除" /></a></td>
					</tr>
				</tbody>
				<%
					}
					}
					try {
				%>
				<tfoot>
					<tr>
						<td colspan="6">
							<div class="pagination">
								总<%=pager.getTotalRecord()%>条记录·总<%=pager.getTotalPage()%>页·当前<%=pager.getCurrentPage() + 1%>页·每页<%=pager.getPageSize()%>条&nbsp;
								<%
									int last = pager.getCurrentRecord() - pager.getPageSize();
										int next = pager.getCurrentRecord() + pager.getPageSize();
										//	int currentRecord;
										if (last < 0) {
											out.println(" 首页  ");
										} else {
											out.print("<a href='AppServlet?currentRecord=" + last + "&methodName=1'><<上一页</a>");
										}
										if (next >= pager.getTotalRecord()) {
											out.println(" 尾页  ");
										} else {
											out.print("<a href='AppServlet?currentRecord=" + next + "&methodName=1'>下一页>></a>");
										}
									} catch (Exception e) {
										e.getMessage();
									}
								%>
							</div>
						</td>
					</tr>

				</tfoot>

			</table>
		</div>
	</div>

</body>
</html>
