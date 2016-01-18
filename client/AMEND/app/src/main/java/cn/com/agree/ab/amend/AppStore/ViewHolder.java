package cn.com.agree.ab.amend.AppStore;

import android.view.View;
import android.widget.ImageView;
import android.widget.ProgressBar;
import android.widget.TextView;

import cn.com.agree.ab.amend.R;

/**
 * Created by Administrator on 2015/10/29.
 */
public class ViewHolder {
    ProgressBar progressBar;
    TextView btn_download;
    TextView txt_process;
    TextView txt_install;
    TextView txt_uninstall;
    TextView txt_cancel;
    TextView txt_update;
    TextView txt_size;
    TextView txt_name;
    TextView txt_continue;
    ImageView img_logo;
    TextView txt_state;


    public ViewHolder(View view) {
        //进度条、下载按钮、下载进度的TextView
        this.progressBar = (ProgressBar) view.findViewById(R.id.pb_downloading);
        this.btn_download = (TextView) view.findViewById(R.id.id_download_weixin);
        this.txt_process = (TextView)view.findViewById(R.id.id_down_sudu);

        //安装、卸载、取消、更新的TextView
        this.txt_install = (TextView) view.findViewById(R.id.id_install);
        this.txt_uninstall = (TextView)view.findViewById(R.id.id_uninstall);
        this.txt_cancel = (TextView) view.findViewById(R.id.id_cancel);
        this.txt_update = (TextView) view.findViewById(R.id.id_update);
        this.txt_continue = (TextView)view.findViewById(R.id.id_continue);

        // 加载app大小
        txt_size = (TextView)view.findViewById(R.id.id_down_daxiao);

        //加载app名字
        txt_name = (TextView)view.findViewById(R.id.id_down_name);

        //加载app图片
        img_logo = (ImageView)view.findViewById(R.id.id_down_logo);

        //加载app状态（是否已安装）
        txt_state = (TextView)view.findViewById(R.id.id_state);
    }

    //点击事件view全部设为不可见
    public void setAllInvisible(){
        btn_download.setVisibility(View.INVISIBLE);
        txt_install.setVisibility(View.INVISIBLE);
        txt_cancel.setVisibility(View.INVISIBLE);
        txt_uninstall.setVisibility(View.INVISIBLE);
        txt_update.setVisibility(View.INVISIBLE);
        txt_continue.setVisibility(View.INVISIBLE);
    }

}
