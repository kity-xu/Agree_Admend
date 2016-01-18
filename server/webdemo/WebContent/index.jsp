<%@ page language="java" pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html lang="en" class="no-js">
<head>
<meta charset="UTF-8" />
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>赞同桌面后台配置</title>
<link rel="stylesheet" type="text/css" href="style/css/normalize.css" />
<link rel="stylesheet" type="text/css" href="style/css/demo.css" />
<link rel="stylesheet" type="text/css" href="style/css/tabs.css" />
<link rel="stylesheet" type="text/css" href="style/css/tabstyles.css" />
<link rel="stylesheet" type="text/css" href="style/background.css" />
<script src="style/js/modernizr.custom.js"></script>
<script type="text/javascript" src="style/js/frame.js"></script>
</head>
<body
	style="background-image: url(style/img/manger.png); background-repeat: no-repeat; background-size: 100% 100%;">
	<svg class="hidden">
<defs>
  <path id="tabshape" d="M80,60C34,53.5,64.417,0,0,0v60H80z"> </path>
</defs>
</svg>
	<div class="container">
		<header class="codrops-header">
			<img src=style/20142181183.png>
			<p class="support">
				Your browser does not support <strong>flexbox</strong>! <br />
				Please view this with <strong><a href="main.jsp">简版</a></strong>.
			</p>
		</header>
		<section>
			<div class="tabs tabs-style-linemove">
				<nav>
					<ul>
						<li><a href="#section-linemove-1"><span>可信MAC地址配置</span></a></li>
						<li><a href="#section-linemove-2"><span>权限密码配置</span></a></li>
						<li><a href="#section-linemove-3"><span>App应用权限配置</span></a></li>
						<li><a href="#section-linemove-4"><span>当前设备信息</span></a></li>
						<li><a href="#section-linemove-5"><span>退出登录</span></a></li>
					</ul>
				</nav>
				<div class="content-wrap">
					<section id="section-linemove-1">
						<iframe src="dbservlet/layout.jsp" class='frame' name=mac></iframe>
					</section>
					<section id="section-linemove-2">
						<iframe src="perpasswordfile/perlayout.jsp" class='frame'></iframe>
					</section>
					<section id="section-linemove-3">
						<iframe src="appfile/applayout.jsp" class='frame' name=main></iframe>
					</section>
					<section id="section-linemove-4">
						<iframe src="current_MAC/CurrentMac.jsp" class='frame' name=main></iframe>
					</section>
					<section id="section-linemove-5">
						<iframe src="exit.html" class='frame'> </iframe>
					</section>
				</div>
				<!-- /content -->
			</div>
			<!-- /tabs -->
		</section>
	</div>
	<div class="footer">
		<b>welcome</b>
	</div>
	<!-- /container -->
	<script src="style/js/cbpFWTabs.js"></script>
	<script>
		(function() {
			[].slice.call(document.querySelectorAll('.tabs')).forEach(
					function(el) {
						new CBPFWTabs(el);
					});
		})();
	</script>
</body>
</html>