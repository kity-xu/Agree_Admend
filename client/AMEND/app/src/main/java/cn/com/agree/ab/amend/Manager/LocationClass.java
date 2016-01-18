package cn.com.agree.ab.amend.Manager;

import android.app.Activity;
import android.app.Service;
import android.content.Context;
import android.content.Intent;
import android.location.Criteria;
import android.location.GpsStatus;
import android.location.Location;
import android.location.LocationListener;
import android.location.LocationManager;
import android.location.LocationProvider;
import android.os.Bundle;
import android.provider.Settings;
import android.util.Log;
import android.widget.TextView;
import android.widget.Toast;


/**
 * Created by Administrator on 2015/10/9.
 */
public class LocationClass {
    private TextView textView;
    private LocationManager lm;
    private static final String TAG = "LocationClass";
    private Location location;
    private Location Curr_Location;//获得刷新后的位置
    private int viewflag = 0;


    public void Destroy() {
//         TODO Auto-generated method stub
        viewflag = 0;
        lm.removeUpdates(locationListener);
    }

    //为activity启动服务
    public void ConnLocationService(Activity activity) {
        lm = (LocationManager) activity.getSystemService(Context.LOCATION_SERVICE);
        // 判断GPS是否正常启动
        if (!lm.isProviderEnabled(LocationManager.GPS_PROVIDER)) {
            Toast.makeText(activity, "请开启GPS导航...", Toast.LENGTH_SHORT).show();
            // 返回开启GPS导航设置界面
            Intent intent = new Intent(Settings.ACTION_LOCATION_SOURCE_SETTINGS);
            activity.startActivityForResult(intent, 0);
            return;
        }
        // 为获取地理位置信息时设置查询条件
        String bestProvider = lm.getBestProvider(getCriteria(), true);
        // 获取位置信息
        // 如果不设置查询要求，getLastKnownLocation方法传人的参数为LocationManager.GPS_PROVIDER
        location = lm.getLastKnownLocation(bestProvider);
        Curr_Location = location;
        if (viewflag == 1)
            updateView(location);
        // 监听状态
        lm.addGpsStatusListener(listener);
        // 绑定监听，有4个参数
        // 参数1，设备：有GPS_PROVIDER和NETWORK_PROVIDER两种
        // 参数2，位置信息更新周期，单位毫秒
        // 参数3，位置变化最小距离：当位置距离变化超过此值时，将更新位置信息
        // 参数4，监听
        // 备注：参数2和3，如果参数3不为0，则以参数3为准；参数3为0，则通过时间来定时更新；两者为0，则随时刷新
        // 1秒更新一次，或最小位移变化超过1米更新一次；
        // 注意：此处更新准确度非常低，推荐在service里面启动一个Thread，在run中sleep(10000);然后执行handler.sendMessage(),更新位置
        lm.requestLocationUpdates(LocationManager.GPS_PROVIDER, 1000, 0, locationListener);
    }

    //为service启动服务
    public void ConnLocationService(Service service) {
        lm = (LocationManager) service.getSystemService(Context.LOCATION_SERVICE);
        if (!lm.isProviderEnabled(LocationManager.GPS_PROVIDER)) {
            Toast.makeText(service, "请开启GPS导航...", Toast.LENGTH_SHORT).show();
        }
        String bestProvider = lm.getBestProvider(getCriteria(), true);
        location = lm.getLastKnownLocation(bestProvider);
        lm.addGpsStatusListener(listener);
        lm.requestLocationUpdates(LocationManager.GPS_PROVIDER, 1000, 0, locationListener);
    }

    // 位置监听
    private LocationListener locationListener = new LocationListener() {
        /**
         * 位置信息变化时触发
         */
        public void onLocationChanged(Location location) {
            Curr_Location = location;
            if (viewflag == 1)
                updateView(location);
//            Log.i(TAG, "时间：" + location.getTime());
//            Log.i(TAG, "经度：" + location.getLongitude());
//            Log.i(TAG, "纬度：" + location.getLatitude());
//            Log.i(TAG, "海拔：" + location.getAltitude());
        }

        /**
         * GPS状态变化时触发
         */
        public void onStatusChanged(String provider, int status, Bundle extras) {
            switch (status) {
                // GPS状态为可见时
                case LocationProvider.AVAILABLE:
                    Log.i(TAG, "当前GPS状态为可见状态");
                    break;
                // GPS状态为服务区外时
                case LocationProvider.OUT_OF_SERVICE:
                    Log.i(TAG, "当前GPS状态为服务区外状态");
                    break;
                // GPS状态为暂停服务时
                case LocationProvider.TEMPORARILY_UNAVAILABLE:
                    Log.i(TAG, "当前GPS状态为暂停服务状态");
                    break;
            }
        }

        /**
         * GPS开启时触发
         */
        public void onProviderEnabled(String provider) {
            Location location = lm.getLastKnownLocation(provider);
            Curr_Location = location;
            if (viewflag == 1)
                updateView(location);
        }

        /**
         * GPS禁用时触发
         */
        public void onProviderDisabled(String provider) {
            if (viewflag == 1)
                updateView(null);
        }
    };

    // 状态监听
    GpsStatus.Listener listener = new GpsStatus.Listener() {
        public void onGpsStatusChanged(int event) {
            switch (event) {
                // 第一次定位
                case GpsStatus.GPS_EVENT_FIRST_FIX:
//                    TODO 解决必须移动超过一定范围才能第一次刷新textview
                    Curr_Location = location;
                    if (viewflag == 1)
                        updateView(location);
                    Log.i(TAG, "第一次定位");
                    break;
                // 卫星状态改变
//                case GpsStatus.GPS_EVENT_SATELLITE_STATUS:
//                    if (viewflag == 1)
//                        updateView(location);
//                    Log.i(TAG, "卫星状态改变");
//                    // 获取当前状态
//                    GpsStatus gpsStatus = lm.getGpsStatus(null);
//                    // 获取卫星颗数的默认最大值
//                    int maxSatellites = gpsStatus.getMaxSatellites();
//                    // 创建一个迭代器保存所有卫星
//                    Iterator<GpsSatellite> iters = gpsStatus.getSatellites()
//                            .iterator();
//                    int count = 0;
//                    while (iters.hasNext() && count <= maxSatellites) {
//                        GpsSatellite s = iters.next();
//                        count++;
//                    }
//                    System.out.println("搜索到：" + count + "颗卫星");
//                    break;
                // 定位启动
                case GpsStatus.GPS_EVENT_STARTED:
                    Log.i(TAG, "定位启动");
                    Curr_Location = location;
                    if (viewflag == 1)
                        updateView(location);
                    break;
                // 定位结束
                case GpsStatus.GPS_EVENT_STOPPED:
                    Curr_Location = location;
                    Log.i(TAG, "定位结束");
                    if (viewflag == 1)
                        updateView(location);
                    break;
            }
        }
    };

    /**
     * 实时更新文本内容
     */
    private void updateView(Location location) {
        if (location != null) {
            textView.setText("设备位置信息\n经度：");
            textView.append(String.valueOf(location.getLongitude()));
            textView.append("\n纬度：");
            textView.append(String.valueOf(location.getLatitude()));
            textView.append("\n时间：");
            textView.append(String.valueOf(location.getTime()));
            textView.append("\n海拔：");
            textView.append(String.valueOf(location.getAltitude()));
        } else {
            textView.setText("正在定位，请查看是否已经打开gps服务");
        }
    }

    /**
     * 返回查询条件
     *
     * @return criteria
     */
    private Criteria getCriteria() {
        Criteria criteria = new Criteria();
        // 设置定位精确度 Criteria.ACCURACY_COARSE比较粗略，Criteria.ACCURACY_FINE则比较精细
        criteria.setAccuracy(Criteria.ACCURACY_FINE);
        // 设置是否要求速度
        criteria.setSpeedRequired(false);
        // 设置是否允许运营商收费
        criteria.setCostAllowed(false);
        // 设置是否需要方位信息
        criteria.setBearingRequired(false);
        // 设置是否需要海拔信息
        criteria.setAltitudeRequired(false);
        // 设置对电源的需求
        criteria.setPowerRequirement(Criteria.POWER_LOW);
        return criteria;
    }

    public Location getLocation() {
        return Curr_Location;
    }

    public void setText(TextView intextView) {
        textView = intextView;
        viewflag = 1;
    }
}
