package bean;

public class PermissionPassword {

	private String permissionID;
	private String permissionInfo;
	private String password;

	public PermissionPassword() {

	}

	public PermissionPassword(String id, String info, String password) {
		this.permissionID = id;
		this.permissionInfo = info;
		this.password = password;
	}

	public String getPermissionId() {
		return permissionID;
	}

	public void setPermissionId(String id) {
		this.permissionID = id;
	}

	public String getPermissionInfo() {
		return permissionInfo;
	}

	public void setPermissionInfo(String info) {
		this.permissionInfo = info;
	}

	public String getPassword() {
		return password;
	}

	public void setPassword(String password) {
		this.password = password;
	}

}
