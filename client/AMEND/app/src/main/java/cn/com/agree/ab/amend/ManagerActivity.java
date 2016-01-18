package cn.com.agree.ab.amend;

import android.app.ActivityManager;
import android.app.AlertDialog;
import android.content.ActivityNotFoundException;
import android.content.BroadcastReceiver;
import android.content.ComponentName;
import android.content.Context;
import android.content.DialogInterface;
import android.content.Intent;
import android.content.IntentFilter;
import android.content.ServiceConnection;
import android.content.res.Configuration;

import android.graphics.Typeface;
import android.graphics.drawable.Drawable;
import android.os.Environment;
import android.os.Handler;
import android.os.IBinder;
import android.os.Message;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.text.TextUtils;
import android.text.method.PasswordTransformationMethod;
import android.view.Gravity;
import android.view.KeyEvent;
import android.view.Menu;
import android.view.MenuItem;
import android.view.MotionEvent;
import android.view.View;
import android.view.ViewGroup;
import android.view.WindowManager;
import android.view.inputmethod.InputMethodManager;
import android.widget.AdapterView;
import android.widget.EditText;
import android.widget.GridView;
import android.widget.ImageButton;
import android.widget.LinearLayout;
import android.widget.TextView;
import android.widget.Toast;


import java.io.File;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Calendar;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Timer;
import java.util.TimerTask;


import cn.com.agree.ab.amend.AppStore.ConfigerClass;
import cn.com.agree.ab.amend.AppStore.getApkInfoFromWeb;
import cn.com.agree.ab.amend.Manager.AppAdapter;
import cn.com.agree.ab.amend.Manager.AppInfo;
import cn.com.agree.ab.amend.Manager.AppManagerUtil;
import cn.com.agree.ab.amend.Manager.CallBack;
import cn.com.agree.ab.amend.Manager.CommunicationService;
import cn.com.agree.ab.amend.Manager.FileOperate;
import cn.com.agree.ab.amend.Manager.LocationClass;
import cn.com.agree.ab.amend.Manager.ScrollLayout;

public class ManagerActivity extends AppCompatActivity {
    static String whoLayout;
    int KEY_RIGHT = 2;
    int KEY_WRONG = 3;
    int KEY_OTHER = 4;
    private static String [] APPS;
    public static List<String> appFilter;
    int SHOW_ANOTHER_ACTIVITY = 1;

    //原Activitylogin变量
    EditText edt_pwd;
    TextView btn_login;
    LocationClass locationClass;
    CommunicationService communicationService;

    private ScrollLayout mScrollLayout;
    private static final float APP_PAGE_SIZE = 12.0f;
    private Context mContext;
    private static int ID_TIME = 9;
    LinearLayout layout;
    TextView txt_date,txt_time;
    String year,mouth,day,hour,minute,date;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_login1);


        ImageButton bt_set = (ImageButton)findViewById(R.id.imageButton);
        bt_set.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(ManagerActivity.this, SettingActivity.class);
                startActivity(intent);
            }
        });
        initLogin();
    }

    void initLogin(){
        whoLayout = "login";
        edt_pwd = (EditText) findViewById(R.id.id_edt_pwd);
        edt_pwd.setTransformationMethod(PasswordTransformationMethod.getInstance());
        btn_login = (TextView) findViewById(R.id.id_btn_login);

        startService();
        btn_login.setOnClickListener(new btnLoginClass());
    }
    void startService(){
        ConfigerClass config = new ConfigerClass(this);
        Intent intent = new Intent(this, CommunicationService.class);
        intent.putExtra("URL", config.Url);
        startService(intent);
        bindService(intent, conn, Context.BIND_AUTO_CREATE);
        locationClass = new LocationClass();
        locationClass.ConnLocationService(this);
    }


    public class btnLoginClass implements View.OnClickListener{
        int res = 0;
        public void onClick(View v) {
            try {
                String pwd = edt_pwd.getText().toString();
                communicationService.checkpwd(pwd, new CallBack() {
                    public void ChkPwdRes() {
                        int res = 0;
                        res = communicationService.getChekResult();
                        communicationService.cleanCheckResult();

                        if (res == 1) {
                            res = 0;
                            Message msg = mHandler.obtainMessage(KEY_RIGHT);
                            mHandler.sendMessageDelayed(msg, 0);
                        } else if (res == 2) {
                            res = 0;
                            Message msg = mHandler.obtainMessage(KEY_WRONG);
                            mHandler.sendMessageDelayed(msg,0);
                        } else if (res == 0) {
                            Message msg = mHandler.obtainMessage(KEY_OTHER);
                            mHandler.sendMessageDelayed(msg,0);
                        } else {
                            Toast.makeText(ManagerActivity.this, "密码错误！", Toast.LENGTH_SHORT).show();
                        }
                    }
                });
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
    ServiceConnection conn = new ServiceConnection() {
        @Override
        public void onServiceDisconnected(ComponentName name) {
            communicationService = null;
        }
        @Override
        public void onServiceConnected(ComponentName name, IBinder service) {
            communicationService = ((CommunicationService.ConnBinder) service).getService();

        }
    };
    /////////////////////-----------------end Activitylogin---------------------////////////////////////////




    void getApplicationLabel(){
        ArrayList<HashMap<String, Object>> Itemlist = new getApkInfoFromWeb(this,APPS).getAppsInfo();
        String [] appLicationLabel = new String[Itemlist.size()];
        for(int i=0; i < Itemlist.size(); i++) {
            appLicationLabel[i] = (String) Itemlist.get(i).get("appLable");
        }
        appFilter = Arrays.asList(appLicationLabel);
    }

    public void create() {
        whoLayout = "manager";
        mContext = this;
        mScrollLayout = (ScrollLayout)findViewById(R.id.ScrollLayoutTest);
        creatreLayout();

        txt_time = (TextView)findViewById(8);
        txt_date = (TextView)findViewById(ID_TIME);

        Integer cacheTime = 1000 * 60;
        final Timer timer = new Timer();
        timer.schedule(new TimerTask() {
            @Override
            public void run() {
                Timer();
            }
        }, 500, cacheTime);

        ImageButton btnStore = (ImageButton)findViewById(11);
        btnStore.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(ManagerActivity.this, AppstoreActivity.class);
                intent.putExtra("apps",APPS);
                startActivity(intent);
            }
        });
        btnStore.setOnLongClickListener(new View.OnLongClickListener() {
            @Override
            public boolean onLongClick(View v) {
                final AlertDialog.Builder builder = new AlertDialog.Builder(ManagerActivity.this);
                builder.setTitle("选择操作");
                //    指定下拉列表的显示数据
                final String[] cities = {"清除下载出错的apk", "清除已下载的apk", "清除所有"};
                final FileOperate fileOperate = new FileOperate(Environment.getExternalStorageDirectory().getPath() + "/update/");
                //    设置一个下拉的列表选择项
                builder.setItems(cities, new DialogInterface.OnClickListener() {
                    @Override
                    public void onClick(DialogInterface dialog, int which) {
                        if (which == 0) {
                            fileOperate.delFalseApk();
                        } else if (which == 1) {
                            fileOperate.delAllApk();
                        } else if (which == 2) {
                            fileOperate.delAll();
                        }
                        Toast.makeText(ManagerActivity.this, " 已清除", Toast.LENGTH_SHORT).show();
                    }
                });
                builder.show();

                return true;
            }
        });

    }

    /**
     * 获取系统所有的应用程序，并根据APP_PAGE_SIZE生成相应的GridView页面
     */

    private void creatreLayout(){
        final List<AppInfo> apps = AppManagerUtil.getInstalledApps(ManagerActivity.this, appFilter);
        Drawable[] drawables = new Drawable[apps.size()];
        AppInfo[] ais = apps.toArray(new AppInfo[0]);
        for (int i = 0; i < ais.length; i++) {
            drawables[i] = ais[i].appIcon;
        }

        final int PageCount = (int)Math.ceil(apps.size() / APP_PAGE_SIZE);
        for (int i=0; i<PageCount; i++) {
            layout = new LinearLayout(this);
            LinearLayout.LayoutParams params  = new LinearLayout.LayoutParams(ViewGroup.LayoutParams.MATCH_PARENT, ViewGroup.LayoutParams.WRAP_CONTENT);
            layout.setOrientation(LinearLayout.VERTICAL);


            LinearLayout timeLayout = new LinearLayout(this);
            LinearLayout.LayoutParams timeParams = new LinearLayout.LayoutParams(ViewGroup.LayoutParams.MATCH_PARENT, ViewGroup.LayoutParams.WRAP_CONTENT);
            timeLayout.setOrientation(LinearLayout.HORIZONTAL);
             //timeLayout.setBackgroundColor(getResources().getColor(R.color.white));
            timeParams.weight=2;
            timeLayout.setLayoutParams(timeParams);
             //timeLayout.setBackgroundColor(getResources().getColor(R.color.yellow));


            LinearLayout date_out = new LinearLayout(this);
            LinearLayout.LayoutParams date_params = new LinearLayout.LayoutParams(ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.MATCH_PARENT);
            date_out.setOrientation(LinearLayout.VERTICAL);
            date_params.weight = 7;


            TextView tx2 = new TextView(this);
            LinearLayout.LayoutParams txt2Params = new LinearLayout.LayoutParams(760, 226);// 设置组件参数
            tx2.setVisibility(View.INVISIBLE);
            txt2Params.topMargin = 140;
            txt2Params.leftMargin = 300;
            tx2.setLayoutParams(txt2Params);

            if(i == 0){
                Typeface tf2=Typeface.createFromAsset(getAssets(), "fonts/4b03b.TTF");
                tx2.setTypeface(tf2); //设置字体

                tx2.setVisibility(View.VISIBLE);
                tx2.setId(8);
                tx2.setTextColor(getResources().getColor(R.color.white));
                tx2.setTextSize(140);
                tx2.setGravity(Gravity.CENTER_HORIZONTAL);
            }
            date_out.addView(tx2,txt2Params);


            TextView tx1 = new TextView(this);
            LinearLayout.LayoutParams txtParams = new LinearLayout.LayoutParams(600, 60);// 设置组件参数
            tx1.setVisibility(View.INVISIBLE);
            txtParams.topMargin = 30;
            txtParams.leftMargin = 360;
            tx1.setLayoutParams(txtParams);

            if(i == 0) {
                Typeface tf1=Typeface.createFromAsset(getAssets(), "fonts/4b03b.TTF");
                tx1.setTypeface(tf1); //设置字体

                tx1.setVisibility(View.VISIBLE);
                tx1.setId(ID_TIME);
                tx1.setTextColor(getResources().getColor(R.color.white));
                tx1.setTextSize(28);
                tx1.setGravity(Gravity.CENTER_HORIZONTAL);
            }
            tx1.setEllipsize(TextUtils.TruncateAt.END);
            tx1.setGravity(Gravity.CENTER);

            date_out.addView(tx1, txtParams);


            date_out.setLayoutParams(date_params);
            timeLayout.addView(date_out, date_params);
            //tx1.setBackgroundColor(getResources().getColor(R.color.white));

            LinearLayout btnlayout = new LinearLayout(this);
            LinearLayout.LayoutParams btnlayParams = new LinearLayout.LayoutParams(ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.MATCH_PARENT);
            btnlayParams.weight=3;
             // btnlayout.setBackgroundColor(getResources().getColor(R.color.orchid));

            ImageButton btn = new ImageButton(this);
            LinearLayout.LayoutParams btnParam = new LinearLayout.LayoutParams(173,173);
            btn.setBackground(getResources().getDrawable(R.drawable.shop));
            btn.setVisibility(View.INVISIBLE);
            btnParam.topMargin = 200;

            if(i == 0){
                btn.setVisibility(View.VISIBLE);
                btn.setId(11);
            }
            btn.setLayoutParams(btnParam);
            btnlayout.addView(btn,btnParam);

            btnlayout.setLayoutParams(btnlayParams);
            timeLayout.addView(btnlayout,btnlayParams);

            layout.addView(timeLayout, timeParams);

            GridView gv = new GridView(this);
            LinearLayout.LayoutParams gridParams = new LinearLayout.LayoutParams(ViewGroup.LayoutParams.WRAP_CONTENT, ViewGroup.LayoutParams.WRAP_CONTENT);
            gv.setAdapter(new AppAdapter(this, apps, i));
            gv.setNumColumns(6);
            gv.setOnItemClickListener(listener);
            gridParams.weight=3;
            gv.setLayoutParams(gridParams);

            layout.addView(gv,gridParams);
            mScrollLayout.addView(layout,params);
        }
    }
    void Timer() {
        Calendar c = Calendar.getInstance();

        year = String.valueOf(c.get(Calendar.YEAR));
        //mouth = String.valueOf(c.get(Calendar.MONTH) + 1);
        int mm = c.get(Calendar.MONTH)+1;
        int dd = c.get(Calendar.DAY_OF_MONTH);
        //hour = String.valueOf(c.get(Calendar.HOUR_OF_DAY));
        int hh = c.get(Calendar.HOUR_OF_DAY);
        int min = c.get(Calendar.MINUTE);
        date = String.valueOf(c.get(Calendar.DAY_OF_WEEK));
        switch (date) {
            case "1":
                date = "星期日";
                break;
            case "2":
                date = "星期一";
                break;
            case "3":
                date = "星期二";
                break;
            case "4":
                date = "星期三";
                break;
            case "5":
                date = "星期四";
                break;
            case "6":
                date = "星期五";
                break;
            case "7":
                date = "星期六";
                break;
            default:
                break;
        }

        if (hh < 10 && 0 <= hh) {
            hour = "0" + hh;
        } else {
            hour = String.valueOf(hh);
        }
        if (min < 10 && 0 <= min) {
            minute = "0" + min;
        } else {
            minute = String.valueOf(min);
        }

        if (mm < 10 && 0 <= mm){
            mouth = "0" + mm;
        } else {
            mouth = String.valueOf(mm);
        }
        if (dd < 10 && 0 <= dd) {
            day = "0" + dd;
        } else {
            day = String.valueOf(dd);
        }
        Message message = new Message();
        message.what = 0;
        mHandler.sendMessage(message);
    }

    /**
     * gridView 的onItemLick响应事件
     */
    public AdapterView.OnItemClickListener listener = new AdapterView.OnItemClickListener() {

        public void onItemClick(AdapterView<?> parent, View view, int position,
                                long id) {
            // TODO Auto-generated method stub
            AppInfo appInfo = (AppInfo)parent.getItemAtPosition(position);
            Intent mainIntent = mContext.getPackageManager()
                    .getLaunchIntentForPackage(appInfo.packageName);
            mainIntent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK);
            try {
                mContext.startActivity(mainIntent);
            } catch (ActivityNotFoundException noFound) {
                Toast.makeText(mContext, "Package not found!", Toast.LENGTH_SHORT).show();
            }
        }

    };

    // this is a timer under 3 parts
    @Override
    public boolean dispatchTouchEvent(MotionEvent ev) {
        resetTime();
        return super.dispatchTouchEvent(ev);
    }
    private void resetTime() {
        mHandler.removeMessages(SHOW_ANOTHER_ACTIVITY);
        Message msg = mHandler.obtainMessage(SHOW_ANOTHER_ACTIVITY);
        mHandler.sendMessageDelayed(msg, 1000 * 30);//無操作30秒钟后退到登陆界面
    }

    private Handler mHandler = new Handler() {
        @Override
        public void handleMessage(Message msg) {
            // TODO Auto-generated method stub
            super.handleMessage(msg);
            if("manager".equals(whoLayout)&&isForeground(ManagerActivity.this, "cn.com.agree.ab.amend.ManagerActivity") && msg.what==SHOW_ANOTHER_ACTIVITY)
            {
                finish();
                System.exit(0);
            }
            if(msg.what == 0){
                txt_time.setText(hour+":"+minute);
                txt_date.setText(date +"    " + day + "-" + mouth + "-" + year);
            }
            if (msg.what == KEY_RIGHT) {
                InputMethodManager imm = (InputMethodManager) getSystemService(Context.INPUT_METHOD_SERVICE);
                imm.hideSoftInputFromWindow(edt_pwd.getWindowToken(), 0);       //关闭软键盘

                //登陆后 加载新布局
                APPS = communicationService.getapps();
                getApplicationLabel();
                setContentView(R.layout.activity_manager);
                create();
            }
            if(msg.what == KEY_WRONG){
                new AlertDialog.Builder(ManagerActivity.this)
                        .setTitle("警告")
                        .setMessage("密码错误")
                        .setPositiveButton("确定", null)
                        .show();
            }
            if(msg.what == KEY_OTHER){
                new AlertDialog.Builder(ManagerActivity.this)
                        .setTitle("警告")
                        .setMessage("连接失败,请检查网络")
                        .setPositiveButton("确定", null)
                        .show();
            }
        }
    };

    /**
     * 判断某个界面是否在前台
     *
     * @param context
     * @param className 某个界面名称
     */
    private boolean isForeground(Context context, String className) {
        if (context == null || TextUtils.isEmpty(className)) {
            return false;
        }
        ActivityManager am = (ActivityManager) context.getSystemService(Context.ACTIVITY_SERVICE);
        List<ActivityManager.RunningTaskInfo> list = am.getRunningTasks(1);
        if (list != null && list.size() > 0) {
            ComponentName cpn = list.get(0).topActivity;
            if (className.equals(cpn.getClassName())) {
                return true;
            }
        }
        return false;
    }



    //广播接受器 刷新页面布局
    public BroadcastReceiver mRefreshBroadcastReceiver = new BroadcastReceiver() {
        @Override
        public void onReceive(Context context, Intent intent) {
            String action = intent.getAction();
            if (action.equals("cn.agree.AppAddorDelete.action.broadcast")) {
               try {
                   layout.removeAllViews();
                   APPS = communicationService.getapps();
                   setContentView(R.layout.activity_manager);
                   getApplicationLabel();
                   create();
               }catch (Exception e){
                   e.printStackTrace();
               }
            }
        }
    };

    @Override
    public void onConfigurationChanged(Configuration newConfig) {
        super.onConfigurationChanged(newConfig);
        if (newConfig.orientation == Configuration.ORIENTATION_LANDSCAPE) {
            Toast.makeText(ManagerActivity.this, "现在是横屏", Toast.LENGTH_SHORT).show();
        } else if (newConfig.orientation == Configuration.ORIENTATION_PORTRAIT){
            Toast.makeText(ManagerActivity.this, "现在是竖屏", Toast.LENGTH_SHORT).show();
        }
    }


    protected Map<String, String> createExtras() {
        Map<String, String> map = new HashMap<String, String>();
        map.put("user", "xulang");
        map.put("pwd", null);
        return map;
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        if (requestCode == 0x03) {
            super.onActivityResult(requestCode, resultCode, data);
        }
    }

    // interception key-back   operation:nothing
     long i=0,fistpress = 0,secondpress = 0;
    public boolean onKeyDown(int keyCode, KeyEvent event) {
        if (keyCode == KeyEvent.KEYCODE_BACK && event.getRepeatCount() == 0) {
            fistpress = secondpress;
            secondpress = System.currentTimeMillis();
            if(secondpress - fistpress > 1000){
                i = 0;
            }else {
                if(++i > 4){
                    finish();
                    System.exit(0);;
                }
            }
            return true;
        }
        return super.onKeyDown(keyCode, event);
    }


    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_amend, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();
        //noinspection SimplifiableIfStatement
        if (id == R.id.action_settings) {
            return true;
        }
        return super.onOptionsItemSelected(item);
    }


    /*************************************************************************/
    @Override
    protected void onStart() {
        super.onStart();
        //动态注册广播 并接受消息
        IntentFilter intentFilter = new IntentFilter();
        intentFilter.addAction("cn.agree.AppAddorDelete.action.broadcast");
        registerReceiver(mRefreshBroadcastReceiver, intentFilter);
    }

    //Activity从后台重新回到前台时被调用
    @Override
    protected void onRestart() {
        super.onRestart();
    }

    //Activity创建或者从被覆盖、后台重新回到前台时被调用
    @Override
    protected void onResume() {
        super.onResume();
    }

    //Activity被覆盖到下面或者锁屏时被调用
    @Override
    protected void onPause() {
        super.onPause();
    }

    //退出当前Activity或者跳转到新Activity时被调用
    @Override
    protected void onStop() {
        super.onStop();
    }

    //退出当前Activity时被调用,调用之后Activity就结束了
    @Override
    protected void onDestroy() {
        android.os.Process.killProcess(android.os.Process.myPid());
        super.onDestroy();
        unregisterReceiver(mRefreshBroadcastReceiver);
    }

    /**
     * Activity被系统杀死时被调用.
     * 例如:屏幕方向改变时,Activity被销毁再重建;当前Activity处于后台,系统资源紧张将其杀死.
     * 另外,当跳转到其他Activity或者按Home键回到主屏时该方法也会被调用,系统是为了保存当前View组件的状态.
     * 在onPause之前被调用.
     */
    @Override
    protected void onSaveInstanceState(Bundle outState) {
        outState.putInt("param", 0);
        super.onSaveInstanceState(outState);
    }


}
