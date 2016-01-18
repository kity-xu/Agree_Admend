package cn.com.agree.ab.amend.AppStore;

import android.content.Context;
import android.os.AsyncTask;

import java.io.BufferedReader;
import java.io.ByteArrayOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.URL;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/**
 * Created by Administrator on 2015/8/31.
 */
public class getApkInfoFromWeb {
    private ConfigerClass config;
    private HttpURLConnection con;
    Context context;
    private String str = null;
    private String localAppPath;
    public boolean txtfileExist;


    private String appLable;
    private String[] appNames;
    private int Size;
    private String packageName;
    private String versionName;
    private int versionCode;

    public getApkInfoFromWeb(Context context, String[] apkName) {
        appNames = apkName;
        this.context = context;
        config  = new ConfigerClass(context);
        localAppPath = config.getLocalFilePath();
    }

    public ArrayList<HashMap<String, Object>> getAppsInfo() throws NullPointerException{
        ArrayList<HashMap<String, Object>> appsList = new ArrayList<>();
        try {
            for (int i = 0; i < appNames.length; i++) {
                final HashMap<String, Object> appsMap = new HashMap<String, Object>();

                final int idex = i;
                final File txtfile = new File(localAppPath, appNames[i] + ".txt");
                if (txtfile.exists()) {
                    txtfileExist = true;
                    getLocalTxt(txtfile);
                } else {
                    //download
                    Thread thread = new Thread() {
                        public void run() {
                            try {
                                DownloadTxtFile(appNames[idex], ".txt");
                            } catch (Exception e) {
                                e.printStackTrace();
                            }
                        }
                    };
                    thread.start();
                    try {
                        thread.join();
                        getLocalTxt(txtfile);
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                }
                File jpgfile = new File(localAppPath, appNames[i] + ".jpg");
                if (jpgfile.exists()) {
                    appsMap.put("icon", config.filePath + appNames[i] + ".jpg");
                } else {
                    Thread thread = new Thread() {
                        public void run() {
                            try {
                                DownloadTxtFile(appNames[idex], ".jpg");
                            } catch (Exception e) {
                                e.printStackTrace();
                            }
                        }
                    };
                    thread.start();
                    try {
                        thread.join();
                        appsMap.put("icon", config.filePath + appNames[i] + ".jpg");
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                }

                //把服务器上部署的所有应用添加到appsList中;
                appsMap.put("versionCode", versionCode);
                appsMap.put("versionName", versionName);
                appsMap.put("packageName", packageName);
                appsMap.put("appLable", appLable);
                appsMap.put("Size", Size);
                appsMap.put("fileName", appNames[i]);
                appsList.add(appsMap);
            }

            //以包名作为应用的唯一标识，即相同包名的多个应用只显示一个;
            //找出包名相同但版本号较小的应用放入removeIdex中;
            int j = 0;
            ArrayList<Integer> removeIdex = new ArrayList<>();
            for (int i = 0; i < appNames.length; i++) {
                for (j = 1 + i; j < appNames.length; j++) {

                    if (appsList.get(i).get("packageName").equals(appsList.get(j).get("packageName")) && i != j) {
                        if ((int) appsList.get(i).get("versionCode") > (int) appsList.get(j).get("versionCode")) {
                            removeIdex.add(j);

                        } else if ((int) appsList.get(i).get("versionCode") < (int) appsList.get(j).get("versionCode")) {
                            removeIdex.add(i);
                        } else {
                            removeIdex.add(i);
                        }
                    }
                }
            }

            //删除removeIdex中重复的元素后放入list中;
            ArrayList<Integer> list = new ArrayList<>();
            for (int i = 0; i < removeIdex.size(); i++) {
                if (!list.contains(removeIdex.get(i)))
                    list.add(removeIdex.get(i));
            }

            //appsList保留唯一包名（即唯一应用），但版本号较大的应用；
            //用于合理的listview及更新使用
            int k = 0;
            for (int i = 0; i < list.size(); i++) {
                //appsList下标一直在改变，所以要根据删除的数量减k前移
                appsList.remove(appsList.get(list.get(i - k)));
                k++;
            }
        } catch (NullPointerException n) {
            n.printStackTrace();
        }
        return appsList;
    }


    /*************************************************************/

    /**
     * 读取本地.txt文件
     */
    private void getLocalTxt(File txtfile) {
        try {
            String encoding = "GBK";
            InputStreamReader read = new InputStreamReader(new FileInputStream(txtfile), encoding);//考虑到编码格式
            BufferedReader bufferedReader = new BufferedReader(read);
            String lineTxt = null;
            String[] strarray;
            while ((lineTxt = bufferedReader.readLine()) != null) {
                strarray = lineTxt.split("=", 2);//使用limit，最多分割成2个字符串

                switch (strarray[0]) {
                    case "versionCode":
                        versionCode = Integer.parseInt(strarray[1]);
                        break;
                    case "versionName":
                        versionName = strarray[1];
                        break;
                    case "packageName":
                        packageName = strarray[1];
                        break;
                    case "applicationLable":
                        appLable = strarray[1];
                        break;
                    case "size":
                        Size = Integer.parseInt(strarray[1]);
                        break;
                    default:
                        break;
                }
            }
            read.close();

        } catch (Exception e) {
            e.printStackTrace();
        }
    }


    /**
     * 从服务器下载app
     *
     * @param name appname;suffix 后缀
     */
    private void DownloadTxtFile(String name, String suffix) {
        /**
         * 连接到服务器
         */
        try {
            URL url = new URL(config.UrlFileInWeb + name + "/" + name + suffix);
            con = (HttpURLConnection) url.openConnection();
            con.setConnectTimeout(1000); //超时时间
            con.setDoOutput(true); //如果要输出，则必须加上此句
            con.setDoInput(true);
            con.connect();
        } catch (Exception e) {
            e.printStackTrace();
        }

        /**
         * 文件保存路劲 "/sdcard/update/"；
         */
        File tmpFile = new File(config.filePath);
        if (!tmpFile.exists()) {
            tmpFile.mkdir();//创建文件夹
        }
        final File file = new File(config.filePath + name + suffix);

        /**
         * 向SD卡写入文件
         */
        try {
            if (con.getResponseCode() == 200) {
                InputStream inputStream = con.getInputStream();
                ByteArrayOutputStream arrayOutputStream = new ByteArrayOutputStream(); //缓存
                byte[] buffer = new byte[1024 * 10];
                int len;
                while (true) {
                    len = inputStream.read(buffer);
                    if (len == -1) {
                        break;  //读取完
                    }
                    arrayOutputStream.write(buffer, 0, len);  //写入
                }
                arrayOutputStream.close();
                inputStream.close();
                con.disconnect();
                byte[] data = arrayOutputStream.toByteArray();
                FileOutputStream fileOutputStream = new FileOutputStream(file);
                fileOutputStream.write(data); //记得关闭输入流
                fileOutputStream.close();
            }
        } catch (MalformedURLException e) {
            e.printStackTrace();
        } catch (Exception e2) {
            e2.printStackTrace();
        }
    }


}
