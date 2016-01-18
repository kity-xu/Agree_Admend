package bean;

public class AppPermission {
	private String appName;
	private String appPath;
	private String permission;

	public AppPermission() {

	}

	public AppPermission(String name, String path, String permission) {
		this.appName = name;
		this.appPath = path;
		this.permission = permission;
	}

	public String getAppName() {
		return appName;
	}

	public void setAppName(String name) {
		appName = name;
	}

	public String getAppPath() {
		return appPath;
	}

	public void setAppPath(String path) {
		appPath = path;
	}

	public String getPermission() {
		return permission;
	}

	public void setPermission(String permission) {
		this.permission = permission;
	}
}
