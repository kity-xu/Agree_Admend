package cn.com.agree.ab.amend.AppStore.Operate;

import android.app.Activity;
import android.content.Context;
import android.os.AsyncTask;
import android.os.Environment;
import android.os.Handler;
import android.os.Message;
import android.util.Log;
import android.view.View;
import android.widget.ProgressBar;
import android.widget.TextView;
import android.widget.Toast;

import java.io.File;
import java.io.FileInputStream;
import java.io.InputStream;
import java.io.RandomAccessFile;
import java.net.HttpURLConnection;
import java.net.URL;

import cn.com.agree.ab.amend.AppStore.ApplicationState;
import cn.com.agree.ab.amend.AppStore.ConfigerClass;
import cn.com.agree.ab.amend.AppStore.ItemInfo;

/**
 * Created by Administrator on 2015/11/16.
 */
public class DownAsynctask extends AsyncTask<Integer, Integer, String>{
    private Activity activity;
    private Context context;
    private ProgressBar bar;
    private TextView txt;
    private Handler handler;
    private ConfigerClass config;
    private ItemInfo info;
    private String apkName;


    int threadNum = 3;
    int currentProgress = 0;
    int threadRunning = 3;
    int threadCount = 3;
    int fileLength;
    int downsize;
    String sdFile;

    public DownAsynctask(Activity activity, ItemInfo info, ProgressBar progressBar, TextView txt, Handler handler){
        this.activity = activity;
        this.context = activity;
        this.bar = progressBar;
        this.txt = txt;
        this.handler = handler;
        this.info = info;
        this.apkName = info.getFilename();
        this.config = new ConfigerClass(activity);
    }


    /*
     * 第一个执行的方法
     * 执行时机：在执行实际的后台操作前，被UI 线程调用
     * 作用：可以在该方法中做一些准备工作，如在界面上显示一个进度条，或者一些控件的实例化，这个方法可以不用实现。
     * @see android.os.AsyncTask#onPreExecute()
     */
    @Override
    protected void onPreExecute() {
        Log.d("Asynctask", "Download Start...");

        super.onPreExecute();
    }


    /*
     * 执行时机：在onPreExecute 方法执行后马上执行，该方法运行在后台线程中
     * 作用：主要负责执行那些很耗时的后台处理工作。可以调用 publishProgress方法来更新实时的任务进度。该方法是抽象方法，子类必须实现。
     * @see android.os.AsyncTask#doInBackground(Params[])
     */
    @Override
    protected String doInBackground(Integer... params) {
        Downloader();
        return "Download Over";
    }


    /*
    * 执行时机：这个函数在doInBackground调用publishProgress时被调用后，UI 线程将调用这个方法.虽然此方法只有一个参数,但此参数是一个数组，可以用values[i]来调用
    * 作用：在界面上展示任务的进展情况，例如通过一个进度条进行展示。此实例中，该方法会被执行100次
    * @see android.os.AsyncTask#onProgressUpdate(Progress[])
    */
    @Override
    protected void onProgressUpdate(Integer... values) {
        // bar.setProgress(downsize);
        //txt.setText(values[0]);
        super.onProgressUpdate(values);
    }


    /*
     * 执行时机：在doInBackground 执行完成后，将被UI 线程调用
     * 作用：后台的计算结果将通过该方法传递到UI 线程，并且在界面上展示给用户
     * result:上面doInBackground执行后的返回值，所以这里是"执行完毕"
     * @see android.os.AsyncTask#onPostExecute(java.lang.Object)
     */
    @Override
    protected void onPostExecute(String result) {
        bar.setVisibility(View.INVISIBLE);
        txt.setVisibility(View.INVISIBLE);
        super.onPostExecute(result);
    }


    void Downloader(){
        sdFile = config.getLocalFilePath();
        String spec = config.UrlFileInWeb+apkName+"/"+apkName+".apk";
        try {
            // 根据下载的地址构建URL对象
            URL url = new URL(spec);
            // 通过URL对象的openConnection()方法打开连接，返回一个连接对象
            HttpURLConnection httpURLConnection = (HttpURLConnection) url
                    .openConnection();
            // 设置请求的头
            //httpURLConnection.setRequestMethod("GET");
            httpURLConnection.setReadTimeout(5000);
            httpURLConnection.setConnectTimeout(5000);
            // 判断是否响应成功
            if (httpURLConnection.getResponseCode() == 200) {
                // 获取下载文件的长度
                fileLength = httpURLConnection
                        .getContentLength();
                //设置进度条的最大值
                Message msg = handler.obtainMessage();
                msg.what = 0;
                msg.getData().putInt("Max", fileLength);
                handler.sendMessage(msg);
                //判断sd卡是否管用
                if (Environment.getExternalStorageState().equals(Environment.MEDIA_MOUNTED)) {
                    // 保存文件
                    // 外部存储设备的路径
                    //创建保存的文件
                    File file = new File(sdFile, apkName);
                    //创建可以随机访问对象
                    RandomAccessFile accessFile = new RandomAccessFile(
                            file, "rwd");
                    // 保存文件的大小
                    // accessFile.setLength(fileLength);
                    // 关闭
                    accessFile.close();
                    // 计算出每个线程的下载大小
                    int threadSize = fileLength / threadNum;
                    // 计算出每个线程的开始位置，结束位置
                    for (int threadId = 1; threadId <= 3; threadId++) {
                        int startIndex = (threadId - 1) * threadSize;
                        int endIndex = threadId * threadSize - 1;
                        if (threadId == threadNum) {// 最后一个线程
                            endIndex = fileLength - 1;
                        }

                        System.out.println("当前线程：" + threadId
                                + " 开始位置：" + startIndex + " 结束位置："
                                + endIndex + " 线程大小：" + threadSize);
                        // 开启线程下载
                        new DownLoadThread(threadId, startIndex,
                                endIndex, spec).start();
                    }
                }else {
                    activity.runOnUiThread(new Runnable() {
                        public void run() {
                            Toast.makeText(context, "SD卡不管用", Toast.LENGTH_LONG).show();
                        }
                    });
                }
            }else {
                //在主线程中运行
                activity.runOnUiThread(new Runnable() {
                    public void run() {
                        Toast.makeText(context, "无法连接服务器", Toast.LENGTH_LONG).show();
                    }
                });
            }

        } catch (Exception e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
    }


    class DownLoadThread extends Thread {

        private int threadId;
        private int startIndex;
        private int endIndex;
        private String path;
        /**
         * 构造函数
         *
         * @param threadId
         *            线程的序号
         * @param startIndex
         *            线程开始位置
         * @param endIndex
         * @param path
         */
        public DownLoadThread(int threadId, int startIndex, int endIndex,
                              String path) {
            super();
            this.threadId = threadId;
            this.startIndex = startIndex;
            this.endIndex = endIndex;
            this.path = path;
        }

        @Override
        public void run() {
            RandomAccessFile raf = null;
            InputStream is = null;
            HttpURLConnection httpURLConnection = null;
            try {
                //获取每个线程下载的记录文件
                File recordFile = new File(sdFile, threadId+"_"+apkName + ".txt");
                if (recordFile.exists()) {
                    // 读取文件的内容
                    InputStream isInput = new FileInputStream(recordFile);
                    // 利用工具类转换
                    String value = StreamTools.streamToStr(isInput);
                    // 获取记录的位置
                    int recordIndex = Integer.parseInt(value);
                    // 将记录的位置赋给开始位置
                    startIndex = recordIndex;
                }

                // 通过path路径构建URL对象
                URL url = new URL(path);
                // 通过URL对象的openConnection()方法打开连接，返回一个连接对象
                httpURLConnection = (HttpURLConnection) url
                        .openConnection();
                // 设置请求的头
                httpURLConnection.setRequestMethod("GET");
                httpURLConnection.setReadTimeout(10000);
                // 设置下载文件的开始位置结束位置
                httpURLConnection.setRequestProperty("Range", "bytes="
                        + startIndex + "-" + endIndex);
                // 获取的状态码
                int code = httpURLConnection.getResponseCode();
                // 判断是否成功
                if (code == 206) {
                    // 获取每个线程返回的流对象
                    is = httpURLConnection.getInputStream();
                    //获取文件的名称
                    String fileName = path.substring(path.lastIndexOf("/")+1);
                    // 根据路径创建文件
                    File file = new File(sdFile, fileName);
                    // 根据文件创建RandomAccessFile对象
                    raf = new RandomAccessFile(file, "rwd");
                    raf.seek(startIndex);
                    // 定义读取的长度
                    int len = 0;
                    // 定义缓冲区
                    byte b[] = new byte[1024 * 100];
                    int total = 0, size =0;
                    int time = 0;
                    // 循环读取
                    File app = new File(new ConfigerClass(context).getLocalFilePath(),apkName+".apk");
                    while ((len = is.read(b)) != -1) {
                        RandomAccessFile threadFile = new RandomAccessFile(
                                new File(sdFile, threadId + "_" + apkName + ".txt"), "rwd");
                        threadFile.writeBytes((startIndex + total) + "");
                        threadFile.close();
                        raf.write(b, 0, len);
                        // 已经下载的大小
                        downsize = total += len;
                        size += len;
                        time++;
                        //解决同步问题
                        synchronized (context) {
                            if (time > 50) {        //  防止过于频繁的更新UI线程
                                Message msg = handler.obtainMessage();
                                msg.what = 1;
                                msg.getData().putInt("Size", size);
                                handler.sendMessage(msg);
                                size = 0;
                                time = 0;

                                publishProgress((int)(downsize*100l/fileLength));
                                if (fileLength - (new ApplicationState(activity).getFileSize(app)) < 200 * 1024) {
                                    time = 51;
                                    sleep(10);
                                    System.out.println("<<<<<<<<<<<<<<<<<<<<<");
                                }
                            }
                            //计算百分比的操作 l表示long型
                            currentProgress += len;
                            final String percent = currentProgress * 100l / fileLength + "%";
                            //创建保存当前进度和百分比的操作
                            RandomAccessFile pbFile = new RandomAccessFile(
                                    new File(sdFile, apkName + "_pb.txt"), "rwd");
                            pbFile.writeBytes(total + ";" + currentProgress + ";" + percent); //fileLength
                            pbFile.close();
                        }
                    }
                    raf.close();
                    is.close();
                    httpURLConnection.disconnect();
                    activity.runOnUiThread(new Runnable() {
                        public void run() {
                            Toast.makeText(context, "当前线程--" + threadId + "--下载完毕", Toast.LENGTH_LONG).show();
                        }
                    });
                    deleteRecordFiles();
                } else {
                    Message msg = handler.obtainMessage();
                    msg.what = 5;
                    handler.sendMessage(msg);
                }
            } catch (Exception e) {
                // TODO Auto-generated catch block
                e.printStackTrace();
                if (is != null) {
                    try{
                        is.close();
                    }catch (Exception e2){
                        e2.printStackTrace();
                    }
                }else if(raf != null){
                    try {
                        raf.close();
                    }catch (Exception e3){
                        e3.printStackTrace();
                    }
                }else if(httpURLConnection != null){
                    try {
                        httpURLConnection.disconnect();
                    }catch (Exception e4){
                        e4.printStackTrace();
                    }
                }try {
                    sleep(3* 1000);
                }catch (Exception e5){
                    e5.printStackTrace();
                }

                threadCount--;
                if(threadCount == 0){
                    try{
                        sleep(5000);
                        Message msg = handler.obtainMessage();
                        msg.what = 6;
                        handler.sendMessage(msg);
                    }catch (Exception e9){
                        e9.printStackTrace();
                    }
                }
            }

        }

    }

    // synchronized避免线程同步
    public synchronized void deleteRecordFiles() {
        threadRunning--;
        if (threadRunning == 0) {
            for (int i = 1; i <= 3; i++) {
                File recordFile = new File(sdFile, i +"_"+apkName+".txt");
                if (recordFile.exists()) {
                    // 删除文件
                    recordFile.delete();
                }
                File pbFile = new File(sdFile,apkName+"_pb.txt");
                if (pbFile.exists()) {
                    pbFile.delete();
                }
            }

            Message msg = handler.obtainMessage();
            msg.what = 2;
            handler.sendMessage(msg);
        }
    }
}
