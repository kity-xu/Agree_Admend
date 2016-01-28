package cn.com.agree.ab.amend.AppStore;

import android.app.Activity;
import android.content.Context;
import android.content.Intent;
import android.net.Uri;
import android.os.Environment;
import android.os.Handler;
import android.os.Message;
import android.view.View;
import android.widget.Toast;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.InputStream;

import cn.com.agree.ab.amend.AppStore.Operate.DownAsynctask;
import cn.com.agree.ab.amend.AppStore.Operate.Download;
import cn.com.agree.ab.amend.AppStore.Operate.Operate;
import cn.com.agree.ab.amend.AppStore.Operate.StreamTools;

/**
 * Created by Administrator on 2015/10/29.
 */
public class OnClickOperate implements View.OnClickListener {
    static final int DOWN_before = 0;
    static final int DOWN_ing = 1;
    static final int DOWN_over = 2;
    static final int DOWN_cancel = 3;
    static final int DOWN_break = 4;

    int filelen,down=0;
    private Activity activity;
    private Context context;
    private ViewHolder holder;
    private int position;
    ItemInfo info;
    Handler DownHandler;

    public OnClickOperate(Activity activity, int position, ItemInfo info, Handler handler,ViewHolder holder) {
        this.activity = activity;
        this.context = activity;
        this.holder = holder;
        this.info = info;
        this.position = position;
        this.DownHandler = handler;
    }

    // 下载 所用的Handler
    Handler handler = new Handler() {
        public void handleMessage(Message msg){
            if(!Thread.currentThread().isInterrupted()){
                switch(msg.what){
                    case DOWN_before:
                        holder.txt_state.setText("下载中...");
                        holder.progressBar.setVisibility(View.VISIBLE);
                        holder.txt_process.setVisibility(View.VISIBLE);

                        filelen = msg.getData().getInt("Max");
                        //holder.progressBar.setMax(filelen);
                        Message ing = DownHandler.obtainMessage();
                        ing.what = 0;
                        ing.arg1 = position;
                        ing.arg2 = filelen;
                        ing.obj = info.getApplable();
                        DownHandler.sendMessage(ing);
                        break;
                    case DOWN_ing:
                        down += msg.getData().getInt("Size");
                        //holder.progressBar.setProgress(down);
                        int x = (int)(down*100l/filelen);
                        //holder.txt_process.setText(x + "%");

                        Message downing = DownHandler.obtainMessage();
                        downing.what = 1;
                        downing.arg1 = down;
                        downing.arg2 = x;
                        downing.obj = holder;
                        DownHandler.sendMessage(downing);
                        break;
                    case DOWN_over:
                        holder.txt_state.setText("下载完成");

                        holder.txt_process.setText("100%");
                        holder.progressBar.setProgress(filelen);

                        Message over = DownHandler.obtainMessage();
                        over.what = 2;
                        over.arg1 = position;   //操作的是那个position
                        over.arg2 = 2; //1:正在下载; 2:下载完成; 3:下载出错;
                        over.obj = info.getApplable();
                        DownHandler.sendMessage(over);
                        break;
                    case DOWN_cancel: //表示按了取消按钮
                        holder.txt_cancel.setVisibility(View.INVISIBLE);
                        holder.btn_download.setVisibility(View.VISIBLE);
                        holder.txt_process.setVisibility(View.INVISIBLE);
                        holder.progressBar.setVisibility(View.INVISIBLE);
                        //向Download 发送取消下载的消息
                        break;
                    case DOWN_break:
                        readPbFile();
                        break;
                    case 5://连接网络出错
                        Toast.makeText(context,"网络异常，断开连接",Toast.LENGTH_SHORT).show();
                        //holder.btn_download.performClick();
                        break;
                    case 6:
                        //点击继续下载
                        holder.txt_state.setText("下载出错");
                        holder.btn_download.setVisibility(View.INVISIBLE);
                        holder.txt_continue.setVisibility(View.VISIBLE);

                        Message flase = DownHandler.obtainMessage();
                        flase.what = 3;
                        flase.arg1 = position;
                        flase.arg2 = 3;
                        flase.obj = info.getApplable();
                        DownHandler.sendMessage(flase);
                        break;
                    default:
                        break;
                }
            }
        }
    };

    @Override
    public void onClick(View v) {
        if(v == holder.btn_download){
            Message msg = handler.obtainMessage();
            msg.what = 4;
            handler.sendMessage(msg);

            holder.txt_continue.setVisibility(View.INVISIBLE);
            new Download(activity,handler,info.getFilename());
            //new DownAsynctask(activity,info,holder.progressBar,holder.txt_process,handler).execute();
        }else{
            Operate operate = new Operate(activity);
            if(v == holder.txt_install){
                operate.install(info.getFilename());
            }else if(v == holder.txt_uninstall){
                operate.uninstall(info.getPackageName());
            }else if(v == holder.txt_continue){
                Message msg = handler.obtainMessage();
                msg.what = 4;
                handler.sendMessage(msg);

               new Download(activity,handler,info.getFilename());
              // new DownAsynctask(activity,info,holder.progressBar,holder.txt_process,handler).execute();
            }else if(v == holder.txt_cancel){
                //取消
            }else if(v == holder.txt_update){
                File file = new File(new ConfigerClass(context).getLocalFilePath(),info.getFilename()+".apk");
                if ( file.exists()&& (file.length()== info.getFilesize())) {
                    String apkname = info.getFilename() + ".apk";
                    Intent i = new Intent(Intent.ACTION_VIEW);
                    i.setDataAndType(Uri.parse("file://" + new File(new ConfigerClass(context).getLocalFilePath(), apkname).toString()), "application/vnd.android.package-archive");
                    context.startActivity(i);
                } else {
                    operate.update(handler, info.getFilename());
                }
            }
        }
    }


    //读取上次下载进度
    private void readPbFile(){
        String sdDir = Environment.getExternalStorageDirectory().getPath()+"/update/";
        File pbFile = new File(sdDir, info.getFilename()+"_pb.txt");
        InputStream is = null;
        try {
            //判断文件是否存在
            if (pbFile.exists()) {
                is = new FileInputStream(pbFile);
            }
        } catch (FileNotFoundException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
        if (is != null) {
            String value = StreamTools.streamToStr(is);
            String[] arr = value.split(";");
            holder.progressBar.setMax(Integer.valueOf(arr[0]));//最大值
            down = Integer.valueOf(arr[1]);//当前值
            holder.progressBar.setProgress(down);
            holder.txt_process.setText(arr[2]);//显示百分比
        }
    }

}
