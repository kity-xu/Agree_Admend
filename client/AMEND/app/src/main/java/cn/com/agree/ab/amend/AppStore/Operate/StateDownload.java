package cn.com.agree.ab.amend.AppStore.Operate;

import java.util.Map;

/**
 * Created by Administrator on 2015/11/10.
 */
public class StateDownload {
    private int position;
    private boolean Ing = false;
    private boolean Over = false;
    private boolean Flase = false;

    public StateDownload(int position){
        this.position = position;
    }

    public void setDownloadIng(){
        Ing = true;
    }
    public void setDownloadOver(){
        Over = true;
    }
    public void setDownloadFlase(){
        Flase = true;
    }

    public int getPosition(){
        return position;
    }
    public boolean isDownloadIng(){
        return  Ing;
    }
    public boolean isDownloadOver(){
        return Over;
    }
    public boolean isDownloadFlase(){
        return Flase;
    }
}
