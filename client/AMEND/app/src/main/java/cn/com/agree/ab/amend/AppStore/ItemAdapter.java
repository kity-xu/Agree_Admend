package cn.com.agree.ab.amend.AppStore;

import android.app.Activity;
import android.app.AlertDialog;
import android.content.Context;
import android.content.DialogInterface;
import android.content.Intent;
import android.graphics.Bitmap;
import android.graphics.BitmapFactory;
import android.graphics.Color;
import android.net.Uri;
import android.os.Message;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.AdapterView;
import android.widget.BaseAdapter;
import android.os.Handler;
import android.widget.ListView;


import java.io.File;
import java.util.ArrayList;
import java.util.HashMap;

import cn.com.agree.ab.amend.R;

/**
 * Created by Administrator on 2015/10/29.
 */
public class ItemAdapter extends BaseAdapter {
    ArrayList<HashMap<String, Object>> Itemlist;
    ItemInfo info;
    private LayoutInflater infater = null;
    private Context context;
    private Activity activity;
    private ListView lv;
    ArrayList<String> listDownItem = new ArrayList<>();
    ViewHolder holderBack=null;


    public ItemAdapter(Activity activity, ListView lv, ArrayList<HashMap<String, Object>> Itemlist) {
        this.activity = activity;
        this.context = activity;
        this.lv = lv;
        infater = (LayoutInflater) context.getSystemService(Context.LAYOUT_INFLATER_SERVICE);
        this.Itemlist = Itemlist;
    }
    @Override
    public int getCount() {
        // TODO Auto-generated method stub
        return Itemlist.size();
    }
    @Override
    public Object getItem(int position) {
        // TODO Auto-generated method stub
        return Itemlist.get(position);
    }
    @Override
    public long getItemId(int position) {
        // TODO Auto-generated method stub
       // return 0;
        return position;
    }


        @Override
    public View getView(final int position, View convertview, final ViewGroup arg2) {
        View view;
        final ViewHolder holder;
        OnClickOperate onClickOperate=null;
        boolean isExist=false;
        boolean isInstalled=false;
        info = new ItemInfo(position,Itemlist);

        if (convertview == null || convertview.getTag() == null) {
            view = infater.inflate(R.layout.appstore_item, arg2, false);
            holder = new ViewHolder(view);
            view.setTag(holder);
        }
        else{
            view = convertview ;
            holder = (ViewHolder) convertview.getTag();
        }

        holder.setAllInvisible();
        setItemView(holder);

        Handler downhandler = new Handler(){
            public void handleMessage(Message msg) {
                if(msg.what == 0) {  //下载前
                    String lable = (String)msg.obj;
                    if(!listDownItem.contains(lable)) {
                        listDownItem.add(lable);
                    }
                    String fLable = lable +"=false";
                    if(listDownItem.contains(fLable)){
                      listDownItem.remove(fLable);
                    }
                    holder.progressBar.setMax(msg.arg2);
                    updateSingleRow(msg.arg1);
                }

                if (msg.what == 1) {  //下载中
                    holderBack = (ViewHolder) msg.obj;
                    if (holderBack != null) {
                        holderBack.progressBar.setProgress(msg.arg1);
                        holderBack.txt_process.setText(msg.arg2 + "%");
                    }
                }

                if(msg.what == 2){   //下载完成
                    String lable = (String)msg.obj;
                    if(listDownItem.contains(lable)){
                        listDownItem.remove(lable);
                    }
                    updateSingleRow(msg.arg1);
                }

                if(msg.what == 3){  //下载出错
                    String lable = (String)msg.obj;
                    String flable = lable +"=false";

                    if(listDownItem.contains(lable)){
                        listDownItem.remove(lable);
                    }
                    if(!listDownItem.contains(flable)){
                        listDownItem.add(flable);
                    }
                    updateSingleRow(msg.arg1);
                }
            }
        };

            Log.d("正在下载的应用-----------",listDownItem.toString());
        if(listDownItem.contains(info.getApplable())){
            holder.progressBar.setVisibility(View.VISIBLE);
            holder.txt_process.setVisibility(View.VISIBLE);
            holder.txt_state.setText("下载中...");

            holder.btn_download.setTextColor(Color.parseColor("#9932CC"));
            holder.txt_continue.setTextColor(Color.parseColor("#9932CC"));
            holder.btn_download.setText("下载中");
            holder.txt_continue.setText("下载中");
            holder.btn_download.setEnabled(false);
            holder.txt_continue.setEnabled(false);
        } else {
            if(listDownItem.contains(info.getApplable()+"=false")){
                holder.txt_process.setText("");
                holder.txt_state.setText("下载出错");

                holder.btn_download.setTextColor(Color.parseColor("#696969"));
                holder.txt_continue.setTextColor(Color.parseColor("#696969"));
                holder.btn_download.setText("下载");
                holder.txt_continue.setText("继续");
                holder.btn_download.setEnabled(true);
                holder.txt_continue.setEnabled(true);
            }else {
                holder.progressBar.setVisibility(View.INVISIBLE);
                holder.txt_process.setVisibility(View.INVISIBLE);
                //  holder.txt_state.setText("下载完成");
            }
        }

        //设置下载按钮的点击事件
        holder.btn_download.setOnClickListener(new OnClickOperate(activity,position, info, downhandler, holder));
        holder.txt_install.setOnClickListener(new OnClickOperate(activity, position, info, downhandler, holder));
        holder.txt_uninstall.setOnClickListener(new OnClickOperate(activity, position,info, downhandler,holder));
        holder.txt_cancel.setOnClickListener(new OnClickOperate(activity, position,info, downhandler,holder));
        holder.txt_update.setOnClickListener(new OnClickOperate(activity, position,info, downhandler,holder));
        holder.txt_continue.setOnClickListener(new OnClickOperate(activity, position,info, downhandler,holder));

        final String xxxpackage=info.getPackageName();
        lv.setOnItemLongClickListener(new AdapterView.OnItemLongClickListener() {
            @Override
            public boolean onItemLongClick(AdapterView<?> parent, View view, int position, long id) {
                showUpdataDialog(xxxpackage);
                return false;
            }
        });
        return view;
    }

    void setItemView(ViewHolder holder){
        boolean isInstalled = false,isExist = false;
        ApplicationState state = new ApplicationState(activity);
        //名字
        holder.txt_name.setText(info.getApplable());
        //大小
        float size = (float)info.getFilesize()/(1024*1024);
        java.text.DecimalFormat df = new java.text.DecimalFormat("#.00");
        holder.txt_size.setText(df.format(size) + " MB");
        //图片
        Bitmap bm = BitmapFactory.decodeFile(info.getImgpath());
        holder.img_logo.setImageBitmap(bm);
        //包名、状态
        if(state.appIsInstalled(info.getPackageName())) {
            holder.txt_state.setText("已安装");
            isInstalled = true;
        }else {
            holder.txt_state.setText("未安装");
        }

        //判断各应用状态，并设置状态
        isExist = state.appIsExist(info.getFilename());
        File file = new File(new ConfigerClass(context).getLocalFilePath(),info.getFilename()+".apk");
        //根据app状态设置相应的点击事件
        if(!isExist&&!isInstalled) {
            holder.btn_download.setVisibility(View.VISIBLE);
            holder.btn_download.setText("下载");
            holder.btn_download.setTextColor(Color.parseColor("#696969"));
        }

        if(!isInstalled&&isExist){
            if(state.getFileSize(file)==info.getFilesize()) {
                holder.txt_install.setVisibility(View.VISIBLE);
            } else if (state.getFileSize(file) < info.getFilesize()){
                holder.txt_continue.setVisibility(View.VISIBLE);
                holder.txt_continue.setText("继续");
                holder.txt_continue.setTextColor(Color.parseColor("#696969"));
            } else {
                holder.txt_state.setText("安装包出错");
                holder.btn_download.setVisibility(View.VISIBLE);
                holder.btn_download.setText("下载");
                holder.btn_download.setTextColor(Color.parseColor("#696969"));
            }
        }
        if (isInstalled){
            if(state.getVersionCode(info.getPackageName()) < info.getVersionCode()){
                //更新
                holder.txt_update.setVisibility(View.VISIBLE);
            }
            else {
                //无更新
                holder.txt_uninstall.setVisibility(View.VISIBLE);
            }
        }

    }


    protected void showUpdataDialog(final String packageName) {
        AlertDialog.Builder builer = new AlertDialog.Builder(context);
        builer.setTitle("卸载应用");
        builer.setMessage("你是否要卸载该应用？");
        //当点确定按钮时从服务器上下载 新的apk 然后安装   װ
        builer.setPositiveButton("确定", new DialogInterface.OnClickListener() {
            public void onClick(DialogInterface dialog, int which) {
                try {
                    Uri packageURI = Uri.parse("package:"+packageName);
                    Intent uninstallIntent = new Intent(Intent.ACTION_DELETE, packageURI);
                    context.startActivity(uninstallIntent);
                }catch (Exception e){
                    e.printStackTrace();
                }
            }
        });
        builer.setNegativeButton("取消", new DialogInterface.OnClickListener() {
            public void onClick(DialogInterface dialog, int which) {
                // TODO Auto-generated method stub
                //do sth
            }
        });
        builer.create().show();
    }

    /**
     * ListView单条更新
     * @param
     */
    private void updateSingleRow(int position){

        if (lv != null) {
            int start = lv.getFirstVisiblePosition();
            for (int i = start, j = lv.getLastVisiblePosition(); i <= j; i++) {
                View view = lv.getChildAt(i - start);
                getView(i, view, lv);
            }
        }
    }


    //当有应用安装或删除时，AppstoreActivity调用此方法 更新listview
    public void ItemChange(){
        notifyDataSetChanged();
    }

    //AppstoreActivity 退出是调用，清除下载标识listDownItem
    public void ClearDownItem(){
        if(listDownItem != null)
            listDownItem.clear();
    }
}
