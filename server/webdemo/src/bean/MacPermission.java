package bean;

public class MacPermission {

	private String macAddress;
	private String macName;
	private String permission;

	public MacPermission() {

	}

	public MacPermission(String address, String name, String permission) {
		this.macAddress = address;
		this.macName = name;
		this.permission = permission;
	}

	public String getMacName() {
		return macName;
	}

	public void setMacName(String name) {
		macName = name;
	}

	public String getMacAddress() {
		return macAddress;
	}

	public void setMacAddress(String address) {
		macAddress = address;
	}

	public String getPermission() {
		return permission;
	}

	public void setPermission(String permission) {
		this.permission = permission;
	}

}
