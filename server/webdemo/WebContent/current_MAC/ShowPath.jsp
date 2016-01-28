<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@ page import="bean.CurrentMac"%>
<%@ page import="dbservlet.CurrentMacServlet"%>
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta name="viewport" content="initial-scale=1.0, user-scalable=no" />
<style type="text/css">
body, html, #allmap {
	width: 100%;
	height: 100%;
	overflow: hidden;
	margin: 0;
	font-family: "微软雅黑";
}
</style>
<script type="text/javascript"
	src="http://api.map.baidu.com/api?v=2.0&ak=GeUGW3bsBzYwb6DGhGPVKOQt"></script>
<script type="text/javascript"
	src="http://api.map.baidu.com/library/CurveLine/1.5/src/CurveLine.min.js"></script>
<title>显示路线</title>
</head>
<body>
	<%
		String macAddress = request.getParameter("MAC_address");
		//String MAC_address = "111111111111";
		response.setCharacterEncoding("UTF-8");
		request.setCharacterEncoding("UTF-8");
		CurrentMacServlet servlet = new CurrentMacServlet();
		ArrayList<CurrentMac> result = servlet.viewPath(macAddress);
		CurrentMac st = new CurrentMac();
		String[] LO = new String[20];
		String[] LA = new String[20];
		String[] DA = new String[20];
	%>
	<div id="allmap"></div>
	<script type="text/javascript">
			var points = [new BMap.Point(113.288, 23.139)
			<%if (!result.isEmpty()) {
				for (int i = 0; i < result.size(); i++) {
					st = result.get(i);
					LO[i] = st.getLongitude();
					LA[i] = st.getLatitude();
					DA[i] = st.getDateAdd();%>
					,new BMap.Point(<%=LO[i]%>, <%=LA[i]%>)
					<%}
			}%>
		//var points = [ new BMap.Point(113.2886889372559, 23.13962965106183),
			//new BMap.Point(113.28932786853032, 23.13995884517671),
				//new BMap.Point(113.29034009082035, 23.139432133833574),
				//new BMap.Point(113.29124058825688, 23.13989300648029),
				//new BMap.Point(113.29213701159672, 23.13995884517671) 
				];


				var pointsname = ["公司"
				<%if (!result.isEmpty()) {
				for (int j = 0; j < result.size(); j++) {%>
						,"<%=DA[j]%>"	
							<%}
			}%>];
		
				
		//地图初始化
		var bm = new BMap.Map("allmap");
		bm.centerAndZoom(new BMap.Point(113.288, 23.139), 16);
		//坐标转换完之后的回调函数
		translateCallback = function(data) {
			if (data.status === 0) {
				for (var i = 0; i < data.points.length; i++) {
					bm.addOverlay(new BMap.Marker(data.points[i]));
					bm.setCenter(data.points[i]);
					var marker = new BMap.Marker(data.points[i]);
					bm.addOverlay(marker);
					var label = new BMap.Label(pointsname[i], {
						offset : new BMap.Size(20, -10)
					});
					marker.setLabel(label);
					if (i > 1) {
						var polyline = new BMap.Polyline([ data.points[i - 1],
								data.points[i] ], {
							strokeColor : "blue",
							strokeWeight : 2,
							strokeOpacity : 0.5
						}); //创建折线   
						bm.addOverlay(polyline); //增加折线
					}
					}
				}
			}
		
		setTimeout(function() {
			var convertor = new BMap.Convertor();
			convertor.translate(points, 1, 5, translateCallback)
		}, 1000);
	</script>
</body>
</html>