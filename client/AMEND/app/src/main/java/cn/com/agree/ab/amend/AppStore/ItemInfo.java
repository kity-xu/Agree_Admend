package cn.com.agree.ab.amend.AppStore;

import java.util.ArrayList;
import java.util.HashMap;

/**
 * Created by Administrator on 2015/10/29.
 */
public class ItemInfo {
    private String fileName;
    private String appLable;
    private String packageName;
    private String imgPath;
    private int Size;
    private int versionCode;
    static public int flag;

    ArrayList<HashMap<String, Object>> Itemlist;
    int position;
    public ItemInfo(int position,ArrayList<HashMap<String, Object>> Itemlist){
        this.Itemlist = Itemlist;
        this.position = position;
        setInfo();
    }

    void setInfo(){
        fileName = (String)Itemlist.get(position).get("fileName");
        appLable = (String)Itemlist.get(position).get("appLable");
        imgPath = (String)Itemlist.get(position).get("icon");
        packageName = (String)Itemlist.get(position).get("packageName");
        Size = (int)Itemlist.get(position).get("Size");
        versionCode = (int)Itemlist.get(position).get("versionCode");
    }

    public String getFilename(){
        return fileName;
    }
    public String getApplable(){
        return appLable;
    }
    public String getPackageName(){
        return packageName;
    }
    public String getImgpath(){
        return imgPath;
    }
    public int getFilesize(){
        return Size;
    }
    public int getVersionCode(){
        return versionCode;
    }
}
