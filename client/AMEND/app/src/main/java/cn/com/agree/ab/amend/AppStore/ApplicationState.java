package cn.com.agree.ab.amend.AppStore;

import android.app.Activity;
import android.content.pm.ApplicationInfo;
import android.content.pm.PackageInfo;
import android.content.pm.PackageManager;
import android.os.Handler;
import android.util.Log;
import android.view.View;

import java.io.File;
import java.io.FileInputStream;
import java.util.List;

/**
 * Created by Administrator on 2015/10/29.
 */
public class ApplicationState {
    Activity activity;
    ConfigerClass config;

    public ApplicationState(Activity activity){
        config = new ConfigerClass(activity);
        this.activity = activity;
    }

    /**
     * 依据文件名判断该apk是否在本地存在;
     * @param fileName;
     * @return if Exist,return true;
     */
    public boolean appIsExist(String fileName){
        fileName=fileName+".apk";
        File apkfile = new File(config.getLocalFilePath(), fileName);
        if (apkfile.exists()) {
            return true;
        }
        return false;
    }
    /**
     * 依据包名判断该应用是否在本地已安装;
     * @name appIsInstalled;
     * @param targetPackage packageName in HashMap;
     * @return boolean;
     */
    public boolean appIsInstalled(String targetPackage) {
        if("".equals(targetPackage)){
            return false;
        }else {
            List<ApplicationInfo> packages;
            PackageManager pm = activity.getPackageManager();
            packages = pm.getInstalledApplications(0);
            for (ApplicationInfo packageInfo : packages) {
                if (packageInfo.packageName.equals(targetPackage)) {
                    return true;
                }
            }
            return false;
        }
    }

    /**
     * 依据包名得到已安装app的版本号
     * @param packageName
     * @return versionCode
     */
    public int getVersionCode(String packageNmae){
        if("".equals(packageNmae)){
            return 0;
        }else {
            List<PackageInfo> packages = activity.getPackageManager().getInstalledPackages(0);
            for (PackageInfo packageInfo : packages) {
                if (packageInfo.packageName.equals(packageNmae)) {
                    return packageInfo.versionCode;
                }
            }
        }
        return 0;
    }

    //获取指定文件大小
    public  int getFileSize(File file) {
        int size = 0;
        try{
            if (file.exists()) {
                FileInputStream fis = null;
                fis = new FileInputStream(file);
                size = fis.available();
            } else {
                file.createNewFile();
                Log.e("获取文件大小", "文件不存在!");
            }
        }catch (Exception e){
            e.printStackTrace();
        }
        return size;
    }
}
