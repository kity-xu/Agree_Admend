package cn.com.agree.ab.amend.Manager;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.ImageView;
import android.widget.TextView;

import java.util.ArrayList;
import java.util.List;

import cn.com.agree.ab.amend.R;

/**
 * Created by Administrator on 2015/10/16.
 */
public class AppAdapter extends BaseAdapter {
    private ArrayList<AppInfo> mList = new ArrayList<>();
    private Context mContext;
    public static final int APP_PAGE_SIZE = 12;

    public AppAdapter(Context context, List<AppInfo> list, int page) {
        mContext = context;
        int i,iEnd;
        if(list.size()>=APP_PAGE_SIZE){
            i = page * APP_PAGE_SIZE;
            iEnd = i + APP_PAGE_SIZE;
        }else {
            i = 0;
            iEnd = APP_PAGE_SIZE;
        }
        while ((i<list.size()) && (i<iEnd)) {
            mList.add(list.get(i));
            i++;
        }
    }
    public int getCount() {
        // TODO Auto-generated method stub
        return mList.size();
    }

    public Object getItem(int position) {
        // TODO Auto-generated method stub
        return mList.get(position);
    }

    public long getItemId(int position) {
        // TODO Auto-generated method stub
        return position;
    }

    public View getView(int position, View convertView, ViewGroup parent) {
        // TODO Auto-generated method stub
        AppInfo appInfo = mList.get(position);
        AppItem appItem;
        if (convertView == null) {
            View v = LayoutInflater.from(mContext).inflate(R.layout.grid_item, null);

            appItem = new AppItem();
            appItem.mAppIcon = (ImageView)v.findViewById(R.id.ivAppIcon);
            appItem.mAppName = (TextView)v.findViewById(R.id.tvAppName);

            v.setTag(appItem);
            convertView = v;
        } else {
            appItem = (AppItem)convertView.getTag();
        }
        // set the icon
        appItem.mAppIcon.setImageDrawable(appInfo.appIcon);
        // set the app name
        appItem.mAppName.setText(appInfo.appName);

        return convertView;
    }

    /**
     * 每个应用显示的内容，包括图标和名称
     * @author Yao.GUET
     *
     */
    class AppItem {
        ImageView mAppIcon;
        TextView mAppName;
    }
}
