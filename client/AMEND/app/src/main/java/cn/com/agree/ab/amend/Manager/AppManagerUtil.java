package cn.com.agree.ab.amend.Manager;

import android.app.Activity;
import android.app.ActivityManager;
import android.content.Context;
import android.content.Intent;
import android.content.pm.PackageInfo;
import android.content.pm.PackageManager;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.Map;

/**
 * Created by Administrator on 2015/9/25.
 */
public class AppManagerUtil {
    public static List<AppInfo> getInstalledApps(Activity activity, List<String> filterPackage) {
        List<AppInfo> appList = new ArrayList<AppInfo>(); //用来存储获取的应用信息数据
        PackageManager pm = activity.getPackageManager();
        List<PackageInfo> packages = pm.getInstalledPackages(0);
        List<String> runningApps = getRunningApps(activity);

        for (int i = 0; i < packages.size(); i++) {
            PackageInfo packageInfo = packages.get(i);
            AppInfo tmpInfo = new AppInfo();
            tmpInfo.appName = packageInfo.applicationInfo.loadLabel(activity.getPackageManager()).toString();
            tmpInfo.packageName = packageInfo.packageName;
            if (filterPackage.contains(tmpInfo.packageName) || filterPackage.contains(tmpInfo.appName)) {
                tmpInfo.versionName = packageInfo.versionName;
                tmpInfo.versionCode = packageInfo.versionCode;
                tmpInfo.launchIntent = pm.getLaunchIntentForPackage(tmpInfo.packageName);
                tmpInfo.appIcon = packageInfo.applicationInfo.loadIcon(activity.getPackageManager());
                tmpInfo.sourceDir = packageInfo.applicationInfo.sourceDir;
                if (runningApps.contains(tmpInfo.packageName)) {
                    tmpInfo.isRunning = true;
                }
                appList.add(tmpInfo);
            }
        }
        return appList;
    }

    public static boolean isRunning(Activity activity,AppInfo appInfo){
        List<String> runningApps = getRunningApps(activity);
        if(runningApps.contains(appInfo.packageName)){
            return true;
        }
        return  false;
    }

    public static void startApp(Activity activity, AppInfo app, Map<String, String> extras,int[] flags) {
        Intent tmpIntent = app.launchIntent;
        tmpIntent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK);
        for(int flag : flags){
            tmpIntent.addFlags(flag);
        }
        for (Iterator<Map.Entry<String, String>> it = extras.entrySet().iterator(); it.hasNext(); ) {
            Map.Entry<String, String> entry = it.next();
            tmpIntent.putExtra(entry.getKey(), entry.getValue());
        }
        tmpIntent.putExtra("action", "start");
        activity.startActivity(tmpIntent);
    }

    public static void stopApp(Activity activity, AppInfo app, Map<String, String> extras) {
        Intent tmpIntent = app.launchIntent;
        tmpIntent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK);
        tmpIntent.addFlags(Intent.FLAG_ACTIVITY_CLEAR_TASK);
        tmpIntent.addFlags(Intent.FLAG_ACTIVITY_CLEAR_TOP);
        for (Iterator<Map.Entry<String, String>> it = extras.entrySet().iterator(); it.hasNext(); ) {
            Map.Entry<String, String> entry = it.next();
            tmpIntent.putExtra(entry.getKey(), entry.getValue());
        }
        tmpIntent.putExtra("action", "stop");
        activity.startActivity(tmpIntent);
    }


    public static List<String> getRunningApps(Activity activity) {
        ActivityManager activityManager = (ActivityManager) activity
                .getSystemService(Context.ACTIVITY_SERVICE);
        List<ActivityManager.RunningAppProcessInfo> runningAppProcesses
                = activityManager.getRunningAppProcesses();
        List<String> list = new ArrayList<String>();
        for (ActivityManager.RunningAppProcessInfo info : runningAppProcesses) {
            list.add(info.processName);
        }
        return list;
    }
}
