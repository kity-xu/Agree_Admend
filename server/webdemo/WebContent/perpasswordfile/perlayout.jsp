<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@ page import="bean.PermissionPassword"%>
<%@ page import="bean.Page"%>
<%@ page import="dbservlet.PerpasswordServlet"%>
<%
	String path = "/webdemo";
	String basePath = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort()
			+ path + "/";
%>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 4.01 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
<head>
<base href="<%=basePath%>">
<title>设备权限密码配置区</title>
<meta http-equiv="pragma" content="no-cache">
<meta http-equiv="cache-control" content="no-cache">
<meta http-equiv="expires" content="0">
<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
<meta http-equiv="description" content="This is my page">
<link rel="stylesheet" type="text/css" href="/webdemo/style/page.css" />
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
		var ID = document.forms[0].permissionID.value;
		var Info = document.forms[0].permissionInfo.value;
		if (ID.length <= 0 && Info.length <= 0) {
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
		<form action="PerpasswordServlet" method="post"
			onSubmit="return validate()" style="margin-left: 25px">
			<input type="hidden" name="methodName" value="5" /> 权限编号： <input
				type="text" name="permissionID" value="" title=""></input> 权限描述： <input
				type="text" name="permissionInfo" value="" title=""></input> <input
				type="submit" value="查询" /> <input type="button" value="添加"
				style="margin-left: 1px"
				onClick="javascript:window.location.href='/webdemo/perpasswordfile/perinput.jsp';" />
			<input type="button" value="刷新" style="margin-left: 1px"
				onclick="javascript:window.location.href='/webdemo/PerpasswordServlet?methodName=<%=1%>&permissionID=<%=""%>&permissionInfo=<%=""%>';" />
		</form>
		<div id="menu">
			<table>
				<thead>
					<tr>
						<th>编号</th>
						<th>权限描述</th>
						<th>密码</th>
						<th>操作</th>
					</tr>
				</thead>
				<tbody>
					<%
						response.setCharacterEncoding("UTF-8");
						request.setCharacterEncoding("UTF-8");
						PerpasswordServlet servlet = new PerpasswordServlet();
						ArrayList<PermissionPassword> result = servlet.select("", ""); //鏉╂柨娲栭弻銉嚄閻ㄥ嫮绮ㄩ弸婊堟肠
						Page pager = new Page();
						pager = servlet.setPage(request, response);
						List<PermissionPassword> subResult = null;
						int currentRecord = pager.getCurrentRecord();
						if (currentRecord == 0) {
							if (pager.getTotalRecord() < 7) {
								subResult = (List<PermissionPassword>) result.subList(0, pager.getTotalRecord());
							} else {
								subResult = (List<PermissionPassword>) result.subList(0, pager.getPageSize());
							}
						} else if (pager.getCurrentRecord() + pager.getPageSize() < result.size()) {
							subResult = (List<PermissionPassword>) result.subList(pager.getCurrentRecord(),
									pager.getCurrentRecord() + pager.getPageSize());
						} else {
							subResult = (List<PermissionPassword>) result.subList(pager.getCurrentRecord(), result.size());
						}
						if (!subResult.isEmpty()) {
							for (int i = 0; i < subResult.size(); i++) {
								PermissionPassword st = subResult.get(i);
					%>
					<tr>
						<%
							out.print("<td width=" + '"' + "200" + '"' + ">" + st.getPermissionId() + "</td>");
									out.print("<td width=" + '"' + "200" + '"' + ">" + st.getPermissionInfo() + "</td>");
									out.print("<td width=" + '"' + "200" + '"' + ">" + st.getPassword() + "</td>");
						%>
						<%
							//有外键约束，不能删除
						%>
						<td width="20"><a
							href="PerpasswordServlet?permissionID=<%=st.getPermissionId()%>&authorization_info=<%=""%>&methodName=<%=4%>"
							title="修改"> <img src="style/img/pencil.png" alt="修改" /></a></td>
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
											out.print("<a href='PerpasswordServlet?currentRecord=" + last + "&methodName=1'><<上一页</a>");
										}
										if (next >= pager.getTotalRecord()) {
											out.println(" 尾页  ");
										} else {
											out.print("<a href='PerpasswordServlet?currentRecord=" + next + "&methodName=1'>下一页>></a>");
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
