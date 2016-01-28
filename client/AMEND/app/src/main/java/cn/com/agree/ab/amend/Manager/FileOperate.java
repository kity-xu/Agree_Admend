package cn.com.agree.ab.amend.Manager;

import java.io.File;

/**
 * Created by Administrator on 2015/11/11.
 */
public class FileOperate {
    String path;

    public FileOperate(String path){
        this.path = path;
    }


    public void delFalseApk(){
        deleteDirFilesLikeName(path,"1_");
        deleteDirFilesLikeName(path,"2_");
        deleteDirFilesLikeName(path,"3_");
        deleteDirFilesLikeName(path,"_pb");
    }
    //
    public void delAllApk(){
        deleteDirFilesLikeName(path,"apk");

    }

    //删除所有
    public void delAll(){
        delAllFile(path);
    }


    //删除文件名包含指定字符的文件
    private void deleteFilesLikeName(File file, String likeName){
        if(file.isFile()){			//是文件
        String temp = file.getName().substring(0,file.getName().lastIndexOf(""));
            if(temp.indexOf(likeName) != -1){
                file.delete();			}
        } else {			//是目录
        File[] files = file.listFiles();
            for(int i = 0; i < files.length; i++){
                deleteFilesLikeName(files[i], likeName);
            }
        }
    }

    //删除某目录下文件名字包含指定字符的文件
    private void deleteDirFilesLikeName(String dir, String likeName){
        File file = new File(dir);
        if(file.exists()){
            deleteFilesLikeName(file, likeName);
        } else {
            System.out.println("路径不存在");
        }
    }

    private void delFolder(String folderPath) {
        try {
            delAllFile(folderPath); //删除完里面所有内容
            String filePath = folderPath;
            filePath = filePath.toString();
            java.io.File myFilePath = new java.io.File(filePath);
            myFilePath.delete(); //删除空文件夹
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    //删除指定文件夹下所有文件
    //param path 文件夹完整绝对路径
    private boolean delAllFile(String path) {
        boolean flag = false;
        File file = new File(path);
        if (!file.exists()) {
            return flag;
        }
        if (!file.isDirectory()) {
            return flag;
        }
        String[] tempList = file.list();
        File temp = null;
        for (int i = 0; i < tempList.length; i++) {
            if (path.endsWith(File.separator)) {
                temp = new File(path + tempList[i]);
            } else {
                temp = new File(path + File.separator + tempList[i]);
            }
            if (temp.isFile()) {
                temp.delete();
            }
            if (temp.isDirectory()) {
                delAllFile(path + "/" + tempList[i]);//先删除文件夹里面的文件
                delFolder(path + "/" + tempList[i]);//再删除空文件夹
                flag = true;
            }
        }
        return flag;
    }
}
