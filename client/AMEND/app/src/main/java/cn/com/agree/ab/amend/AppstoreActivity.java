package cn.com.agree.ab.amend;

import android.content.IntentFilter;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.Intent;
import android.util.Log;
import android.view.KeyEvent;
import android.view.View;
import android.widget.AbsListView;
import android.widget.ListView;
import android.widget.Toast;

import java.util.ArrayList;
import java.util.HashMap;


import cn.com.agree.ab.amend.AppStore.ConfigerClass;
import cn.com.agree.ab.amend.AppStore.ItemAdapter;
import cn.com.agree.ab.amend.AppStore.getApkInfoFromWeb;

public class AppstoreActivity extends AppCompatActivity {
    ConfigerClass config;
    static  Boolean addOrDelete = false;
    String [] apps ;
    static ArrayList<HashMap<String, Object>> Itemlist;
    static ListView lv;
    static ItemAdapter adapter;
    private boolean is_divPage;// 是否进行分页操作

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.appstore_list);
        lv = (ListView)findViewById(R.id.listviewApp);

        config= new ConfigerClass(this);
        //String localLabel = this.getApplicationInfo().loadLabel(this.getPackageManager()).toString(); //获取本应用的label
        apps= (String [])getIntent().getSerializableExtra("apps");
        Itemlist = new getApkInfoFromWeb(this,apps).getAppsInfo();

        adapter= new ItemAdapter(this,lv,Itemlist);
        lv.setAdapter(adapter);
        lv.setOnScrollListener(new listviewScroll());
    }


    class listviewScroll implements AbsListView.OnScrollListener {
        @Override
        public void onScrollStateChanged(AbsListView view, int scrollState) {
            /**
             * 如果等到该分页（is_divPage = true）的时候，并且滑动停止（这个时候手已经离开了屏幕），自动加载更多。
             */
            if (is_divPage && scrollState == AbsListView.OnScrollListener.SCROLL_STATE_IDLE) {
               // Toast.makeText(AppstoreActivity.this, "正在获取更多数据", Toast.LENGTH_SHORT).show();

               // adapter.notifyDataSetChanged();

            } else if (!is_divPage && scrollState == AbsListView.OnScrollListener.SCROLL_STATE_IDLE) {

            }
        }

        @Override
        public void onScroll(AbsListView view, int firstVisibleItem,
                             int visibleItemCount, int totalItemCount) {
            is_divPage = (firstVisibleItem + visibleItemCount == totalItemCount);
        }
    }




    //广播接收器，接受添加、删除app的系统广播,更新listview
    public static class getBroadcast extends BroadcastReceiver {
        public void onReceive(Context context, Intent intent) {
            if (intent.getAction().equals("android.intent.action.PACKAGE_ADDED")) {
                addOrDelete = true;
                try{
                    //lv.setAdapter(adapter);
                    adapter.ItemChange();
                }catch (Exception e){
                    e.printStackTrace();
                }
            }
            if (intent.getAction().equals("android.intent.action.PACKAGE_REMOVED")) {
                addOrDelete = true;
                try{
                    //lv.setAdapter(adapter);
                    adapter.ItemChange();
                }catch (Exception e){
                    e.printStackTrace();
                }
            }
        }
    }

    @Override
    public boolean onKeyDown(int keyCode, KeyEvent event) {
        if (keyCode == KeyEvent.KEYCODE_BACK && event.getRepeatCount() == 0) {
            Intent intent = new Intent();
            intent.setAction("cn.agree.AppActivityBack.action.broadcast");      //按返回键时发送下载线程结束广播；
            if(addOrDelete) {
                intent.setAction("cn.agree.AppAddorDelete.action.broadcast");   //当有应用添加或删除时通知GirdView更新（ManagerActivity重新布局）;
                intent.putExtra("Key", "flush");  // 要发送的内容
               // finish();
            }
            AppstoreActivity.this.sendBroadcast(intent);//  发送 一个无序广播
            adapter.ClearDownItem();
        }
        return super.onKeyDown(keyCode, event);
    }

}
