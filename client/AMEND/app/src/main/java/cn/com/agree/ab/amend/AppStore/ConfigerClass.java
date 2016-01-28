package cn.com.agree.ab.amend.AppStore;

import android.content.Context;
import android.content.SharedPreferences;
import android.os.Environment;


/**
 * Created by Administrator on 2015/9/7.
 */
public class ConfigerClass {
    SharedPreferences sp;
    String servlet;
    String ip;
    String port;

    public String Url;
    public String UrlFileInWeb;
    public String filePath = "/sdcard/update/";
    public String tmpFile;

    public ConfigerClass(Context context){
        this.sp = context.getSharedPreferences("filename", 0x8000+0x0008);
        this.servlet = sp.getString("server", "");
        this.ip = sp.getString("ip", "");
        this.port = sp.getString("port", "");
        Url = "http://"+ ip+":"+port+"/"+servlet+"/"+"AndroidServlet";
        UrlFileInWeb = "http://"+ ip+":"+port+"/"+servlet+"/resource/";
        tmpFile = Environment.getExternalStorageDirectory().getPath()+"/Agree/";
    }
    public String getLocalFilePath(){
        return Environment.getExternalStorageDirectory().getPath()+"/update/";
    }

    //public String Url = "http://192.168.1.175:8080/webdemo/AndroidServlet";
    //public String UrlFileInWeb = "http://192.168.1.175:8080/webdemo/resource/";
}
