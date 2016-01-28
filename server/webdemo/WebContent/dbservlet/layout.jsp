<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@ page import="bean.MacPermission"%>
<%@ page import="bean.Page"%>
<%@ page import="dbservlet.MacServlet"%>
<%
	//mac权限管理
	String path = "/webdemo";
	String basePath = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort()
			+ path + "/";
%>

<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 4.01 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
<head>
<base href="<%=basePath%>">

<title>可信MAC地址配置区</title>

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
			return false;
		}
	}
	function validate() {
		var mac_Address = document.forms[0].macAddress.value;
		var mac_Name = document.forms[0].macName.value;
		if (mac_Address.length <= 0 && mac_Name.length <= 0) {
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
		<form action="MacServlet" method="post" style="margin-left: 25px"
			onSubmit="return validate()">
			<input type="hidden" name="methodName" value="5" /> MAC地址： <input
				type="text" name="macAddress" value="" title=""></input> 设备名称： <input
				type="text" name="macName" value="" title=""></input> <input
				type="submit" value="查询" /> <input type="button" value="添加"
				style="margin-left: 1px"
				onClick="javascript:window.location.href='/webdemo/dbservlet/input.jsp';" />
			<input type="button" value="刷新" style="margin-left: 1px"
				onclick="javascript:window.location.href='/webdemo/MacServlet?methodName=<%=1%>&macAddress=<%=""%>&macName=<%=""%>';" />
		</form>
		<div id="menu">
			<table>
				<thead>
					<tr>
						<th>MAC地址</th>
						<th>设备名称</th>
						<th>权限设置</th>
						<th>操作</th>
					</tr>
				</thead>
				<tbody>
					<%
						response.setCharacterEncoding("UTF-8");
						request.setCharacterEncoding("UTF-8");
						MacServlet servlet = new MacServlet();
						ArrayList<MacPermission> result = servlet.select("", "");
						Page pager = new Page();
						pager = servlet.setPage(request, response);
						List<MacPermission> subResult = null;
						int currentRecord = pager.getCurrentRecord();
						if (currentRecord == 0) {
							if (pager.getTotalRecord() < 7) {
								subResult = (List<MacPermission>) result.subList(0, pager.getTotalRecord());
							} else {
								subResult = (List<MacPermission>) result.subList(0, pager.getPageSize());
							}
						} else if (pager.getCurrentRecord() + pager.getPageSize() < result.size()) {
							subResult = (List<MacPermission>) result.subList(pager.getCurrentRecord(),
									pager.getCurrentRecord() + pager.getPageSize());
						} else {
							subResult = (List<MacPermission>) result.subList(pager.getCurrentRecord(), result.size());
						}

						if (!subResult.isEmpty()) {
							for (int i = 0; i < subResult.size(); i++) {
								MacPermission st = subResult.get(i);
					%>
					<tr>
						<%
							out.print("<td width=" + '"' + "200" + '"' + ">" + st.getMacAddress() + "</td>");
									out.print("<td width=" + '"' + "200" + '"' + ">" + st.getMacName() + "</td>");
									out.print("<td width=" + '"' + "200" + '"' + ">" + st.getPermission() + "</td>");
						%>

						<td width="20"><a
							href="MacServlet?macAddress=<%=st.getMacAddress()%>&macName=<%=""%>&methodName=<%=4%>"
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

								<%
									int totalRecord = pager.getTotalRecord();
										int totalPage = pager.getTotalPage();
										int currentPage = pager.getCurrentPage() + 1;
										int pageSize = pager.getPageSize();
								%>
								总<%=totalRecord%>条记录·总<%=totalPage%>页·当前<%=currentPage%>页·每页<%=pageSize%>条&nbsp;
								<%
									int last = pager.getCurrentRecord() - pager.getPageSize();
										int next = pager.getCurrentRecord() + pager.getPageSize();
										//	int currentRecord;
										if (last < 0) {
											out.println(" 首页  ");
										} else {
											out.print("<a href='MacServlet?currentRecord=" + last + "&methodName=1'><<上一页</a>");
										}
										if (next >= pager.getTotalRecord()) {
											out.println(" 尾页  ");
										} else {
											out.print("<a href='MacServlet?currentRecord=" + next + "&methodName=1'>下一页>></a>");
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
