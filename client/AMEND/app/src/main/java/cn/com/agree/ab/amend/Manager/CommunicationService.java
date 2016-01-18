package cn.com.agree.ab.amend.Manager;


import android.app.Service;
import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.Intent;
import android.content.IntentFilter;
import android.location.Location;
import android.net.wifi.WifiInfo;
import android.net.wifi.WifiManager;
import android.os.Binder;
import android.os.IBinder;
import android.util.Log;

import com.alibaba.fastjson.JSONObject;

import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;


public class CommunicationService extends Service {
    public CommunicationService() {
    }

    private static final String TAG = "CommunicationService";
    private LocationClass locationClass;
    private boolean ThreadExit = false;
    private int Timeout = 10;

    int chkresult = 0;
    private String pwd;
    JSONObject resultObj = new JSONObject();
    static private String URL;

    public void onCreate() {
        super.onCreate();

        locationClass = new LocationClass();
        locationClass.ConnLocationService(this);
        initTimeTick();
        sendmsg();
        Log.i(TAG, "CommunicationCreate");
    }

    private int initClientHttp(final String msg) {
        final URL[] url = {null};
        try {
            ThreadExit = true;
            url[0] = new URL(URL);
            Log.i("URL---communication：", URL);
            HttpURLConnection urlConnection = (HttpURLConnection) url[0].openConnection();
            urlConnection.setDoOutput(true);
            urlConnection.setDoInput(true);
            urlConnection.setUseCaches(false);
            urlConnection.setInstanceFollowRedirects(true);

            urlConnection.setRequestProperty("Content-Type", "application/manager");
            urlConnection.connect();
            DataOutputStream outputStream = new DataOutputStream(urlConnection.getOutputStream());

            outputStream.writeBytes(msg);
            outputStream.flush();
            outputStream.close();

            BufferedReader reader = new BufferedReader(new InputStreamReader(urlConnection.getInputStream()));
            String line = reader.readLine();
            if (line != null) {
                Log.i("line", line);
            }
            JSONObject jsonrecive = JSONObject.parseObject(line);
            String method = jsonrecive.getString("Key");

            if ("right".equals(method)) {
                resultObj = jsonrecive;
                ThreadExit = false;
                return 1;
            }
            if ("wrong".equals(method)) {
                ThreadExit = false;
                return 2;
            }
            if ("Record".equals(method)) {
                ThreadExit = false;
                return 9;
            }
            reader.close();
            urlConnection.disconnect();
            Log.i("接收", line);
        } catch (Exception e) {
            ThreadExit = false;
            e.printStackTrace();
        }
        ThreadExit = false;
        return 0;
    }

    //获取坐标对象
    private Location getlocation() {
        Log.i(TAG, "getlocation");
        return locationClass.getLocation();
    }

    //获取MAC
    public String getLocalMacAddress() {
        WifiManager wifi = (WifiManager) getSystemService(Context.WIFI_SERVICE);
        WifiInfo info = wifi.getConnectionInfo();
        return info.getMacAddress();
    }

    public void checkpwd(String pwd, final CallBack callBack) {
        this.pwd = pwd;
        String msg = getmsg("CheckPwd", 1);

        Log.i(TAG, msg);
        final String finalMsg = msg;
        new Thread(new Runnable() {
            @Override
            public void run() {
                chkresult = initClientHttp(finalMsg);
                callBack.ChkPwdRes();
            }
        }).start();
    }

    public int getChekResult() {
        return chkresult;
    }

    public void cleanCheckResult() {
        chkresult = 0;
    }

    //打包msg
    private String getmsg(String method, int lenth) {
        JSONObject jsonObject = new JSONObject();
        JSONObject LocationJsonObject = new JSONObject();
        JSONObject DataJsonObject = new JSONObject();
        Log.i(TAG, "getmsg");
        Location location = getlocation();
        String MAC = getLocalMacAddress().replace(";", "").replace(":", "");
        jsonObject.clear();
        jsonObject.put("MAC", MAC);
        if (location != null) {
            LocationJsonObject.put("Longitude", location.getLongitude());
            LocationJsonObject.put("Latitude", location.getLatitude());
        } else {

            LocationJsonObject.put("Longitude", 0.0);
            LocationJsonObject.put("Latitude", 0.0);
        }
        jsonObject.put("Location", LocationJsonObject);
        jsonObject.put("Key", method);
        DataJsonObject.put("Lenth", lenth);
        if (lenth == 1) {
            DataJsonObject.put("Data1", pwd);
            pwd = "";
        }
        jsonObject.put("Data", DataJsonObject);
        return jsonObject.toString();
    }


    //广播接收器

    private void initTimeTick() {
        IntentFilter filter = new IntentFilter();
        filter.addAction(Intent.ACTION_TIME_TICK);
        registerReceiver(receiver, filter);
    }

    //每分钟
    private final BroadcastReceiver receiver = new BroadcastReceiver() {
        @Override
        public void onReceive(Context context, Intent intent) {
            String action = intent.getAction();
            if (action.equals(Intent.ACTION_TIME_TICK)) {
                Log.i(TAG, "time tick is on");
                sendmsg();
            }
        }
    };

    private void sendmsg() {
        Log.i(TAG, String.valueOf(Timeout));
        if (Timeout <= 0) {
            //TODO restart manager activi
        }
        String msg = getmsg("0", 0);
        Log.i(TAG, msg);
        if (!ThreadExit) {
            Log.i(TAG, "conn");
            final String finalMsg;
            finalMsg = msg;
            new Thread(new Runnable() {
                @Override
                public void run() {
                    if (initClientHttp(finalMsg) == 9) {
                        Timeout = 10;//记录坐标成功
                    } else {
                        Timeout--;
                    }
                }
            }).start();
        } else {
            Timeout--;
        }
    }

    public String[] getapps() {
        JSONObject Data = (JSONObject) resultObj.get("Data");
        int lenth= (int) Data.get("long");
        String[] res=new String[lenth];
        for(int i=0;i<lenth;i++){
            res[i]= (String) Data.get("data"+i);
        }
        return res;
    }

    @Override
    public void onStart(Intent intent, int startId){
       super.onStart(intent,startId);
        try {
            String tmp = intent.getStringExtra("URL");
            if(!"".equals(tmp))
            URL = tmp;
        }catch (Exception e){
            e.printStackTrace();
        }

    }

    @Override
    public IBinder onBind(Intent intent) {
        return new ConnBinder();
    }


    public class ConnBinder extends Binder {
        public CommunicationService getService() {
            return CommunicationService.this;
        }
    }

    public void onDestroy() {
        locationClass.Destroy();
        super.onDestroy();
    }
}
