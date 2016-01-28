package cn.com.agree.ab.amend.AppStore.Operate;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;

/**
 * Created by Administrator on 2015/10/28.
 */
public class StreamTools {
    public static String streamToStr(InputStream is){
        String value = null;
        try {
            ByteArrayOutputStream baos = new ByteArrayOutputStream();
            // 定义读取的长度
            int len = 0;
            // 定义缓冲区
            byte b[] = new byte[1024];
            // 循环读取
            while ((len = is.read(b)) != -1) {
                baos.write(b, 0, len);
            }
            baos.close();
            is.close();
            value = new String(baos.toByteArray());
        } catch (IOException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
        return value;
    }
}
