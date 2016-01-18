package cn.com.agree.ab.amend.Manager;

import android.content.Intent;
import android.graphics.drawable.Drawable;

/**
 * Created by Administrator on 2015/9/25.
 */
public class AppInfo {
    public String appName = "";
    public String packageName = "";
    public String versionName = "";
    public int versionCode = 0;
    public Drawable appIcon = null;
    public Intent launchIntent = null;
    public boolean isRunning = false;
    public String sourceDir = null;

    public String toString() {
        return "appName[" + appName + "]  packageName[" + packageName + "]  version[" + versionName
                + "]  icon[" + (appIcon != null ? appIcon.toString() : "null") + "]  launchIntent[" +
                (launchIntent != null ? launchIntent.toString() : "null") + "] isRunning[" + isRunning + "]";
    }
}
