package cn.com.agree.ab.amend.AppStore.Operate;

import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.net.Uri;
import android.os.Handler;
import android.os.Message;
import android.widget.Toast;

import java.io.File;
import java.lang.String;

import cn.com.agree.ab.amend.AppStore.ApplicationState;
import cn.com.agree.ab.amend.AppStore.ConfigerClass;


/**
 * Created by Administrator on 2015/11/2.
 */
public class Operate {
    private Activity activity;
    private Context context;
    private ConfigerClass config;
    private ApplicationState state;

    public Operate(Activity activity){
        this.activity = activity;
        this.context = activity;
        config = new ConfigerClass(activity);
        state = new ApplicationState(activity);
    }

    //安装
    public void install(String filename){
        if (state.appIsExist(filename)) {
            String apkname = filename + ".apk";
            final Intent i = new Intent(Intent.ACTION_VIEW);
            i.setDataAndType(Uri.parse("file://" + new File(config.getLocalFilePath(), apkname).toString()), "application/vnd.android.package-archive");
            context.startActivity(i);
            //new getBroadcast(ListviewActivity.this,i,holder);
        } else {
            Toast.makeText(activity, "App不存在或路径不正确", Toast.LENGTH_SHORT).show();
        }
    }

    //卸载
    public void uninstall(String packageName){
        try {
            Uri packageURI = Uri.parse("package:"+packageName);
            final Intent uninstallIntent = new Intent(Intent.ACTION_DELETE, packageURI);
            context.startActivity(uninstallIntent);
        }catch (Exception e){
            e.printStackTrace();
        }
    }

    //取消
    public void cancel(Handler handler, int flag){
        Message msg = new Message();
        msg.what = flag;
        handler.sendMessage(msg);
    }

    //更新
    public void update(Handler handler, String fileName){
           new Download(activity, handler, fileName);
    }
}
