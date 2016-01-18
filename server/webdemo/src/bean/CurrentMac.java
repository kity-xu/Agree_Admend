package bean;

public class CurrentMac {
	private String macAddress;
	private String dateAdd;
	private String longitude;
	private String latitude;

	public CurrentMac() {

	}

	public CurrentMac(String address, String date, String longitude, String latitude) {
		this.macAddress = address;
		this.dateAdd = date;
		this.longitude = longitude;
		this.latitude = latitude;
	}

	public String getMacAddress() {
		return macAddress;
	}

	public void setMacAddress(String address) {
		macAddress = address;
	}

	public String getDateAdd() {
		return dateAdd;
	}

	public void setDateAdd(String date) {
		this.dateAdd = date;
	}

	public String getLongitude() {
		return longitude;
	}

	public void setLongitude(String longitude) {
		this.longitude = longitude;
	}

	public void setLatitude(String latitude) {
		this.latitude = latitude;
	}

	public String getLatitude() {
		return latitude;
	}
}
