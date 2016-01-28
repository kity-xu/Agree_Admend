package apkutil;

public class ApkPath {
	// TODO 根据环境设置默认路径-set default path
	private String AaptPath = "C:\\Users\\sl\\AppData\\Local\\Android\\sdk\\build-tools\\23.0.1\\aapt.exe";// apk反编译工具，在androidSDK中可以找到
	private String InfoPath = "C:\\Users\\sl\\Desktop\\info.txt";// 生成的信息文件路径
	private String IconPath = "C:\\Users\\sl\\Desktop\\icon.png";// 应用图标路径
	private String apkpath = "C:\\Users\\sl\\Desktop\\QQ.apk";// 应用安装文件路径
	private String basePath = "F:\\eclipse\\workspace\\webdemo\\WebContent\\resource";// 服务器资源路径
	// private String basePath = "C:\\Program
	// Files\\apache-tomcat-8.0.24\\webapps\\webdemo\\resource";// 服务器资源文件夹路径
	private String TempPath = "F:\\temp\\";// 上传缓存目录

	public String getAaptPath() {
		return AaptPath;
	}

	public void setAaptPath(String mAaptPath) {
		this.AaptPath = mAaptPath;
	}

	public String getInfoPath() {
		return InfoPath;
	}

	public void setInfoPath(String infoPath) {
		InfoPath = infoPath;
	}

	public String getIconPath() {
		return IconPath;
	}

	public void setIconPath(String iconPath) {
		IconPath = iconPath;
	}

	public String getApkpath() {
		return apkpath;
	}

	public void setApkpath(String apkpath) {
		if (apkpath.contains("/") || apkpath.contains("\\")) {
			this.apkpath = apkpath;
		} else {
			this.apkpath = TempPath + apkpath;
		}
	}

	public String getbasePath() {
		return basePath;
	}

}
