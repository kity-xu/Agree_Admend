package dbservlet;

import java.io.BufferedInputStream;
import java.io.BufferedOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStreamWriter;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.zip.ZipEntry;
import java.util.zip.ZipFile;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.log4j.Logger;

import apkutil.ApkInfo;
import apkutil.ApkPath;
import apkutil.ApkUtil;
import bean.AppPermission;
import bean.Page;

public class AppServlet extends MyServlet {
	/**
	 * 
	 */

	private static final long serialVersionUID = 1L;
	private static Logger logger = Logger.getLogger(AppServlet.class);

	public void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		request.setCharacterEncoding("UTF-8");
		response.setCharacterEncoding("UTF-8");
		String methodName = request.getParameter("methodName");
		int method = Integer.parseInt(methodName);
		try {
			switch (method) {
			case 0:
				if (setAppPath(request, response)) {
					super.insert(request, response, Servlet.AppServlet);
				}
				break;
			case 1:
				super.difPage(request, response, Servlet.AppServlet);
				break;
			case 2:
				super.delete(request, response, Servlet.AppServlet);
				delete(request, response);
				break;
			case 3:
				super.updated(request, response, Servlet.AppServlet);
				break;
			case 4:
				super.update(request, response, Servlet.AppServlet);
				break;
			case 5:
				super.dispatch(request, response, Servlet.AppServlet);
				break;
			}
		} catch (ClassNotFoundException e) {
			// e.printStackTrace();
			logger.info(e);
		} catch (SQLException e) {
			// e.printStackTrace();
			logger.info(e);
		}
	}

	public void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		logger.info("doGet");
		doPost(request, response);
	}

	public ArrayList<AppPermission> select(String attr1, String attr2) throws ClassNotFoundException, SQLException {
		return super.appSelect(attr1, attr2);
	}

	public Page setPage(HttpServletRequest request, HttpServletResponse response)
			throws ClassNotFoundException, SQLException {
		return super.setPage(request, response, Servlet.AppServlet);
	}

	//TODO 部署app位置不一定是服务器本地
	public boolean setAppPath(HttpServletRequest request, HttpServletResponse response) {
		String appName = request.getParameter("appName");
		logger.info("应用名为：" + appName);
		String appPath = request.getParameter("appPath");
		logger.info("文件为：" + appPath);
		ApkPath apkpath = new ApkPath();
		String filePath = apkpath.getbasePath() + "\\" + appName;
		try {
			apkpath.setInfoPath(filePath + "\\" + appName + ".txt");
			apkpath.setIconPath(filePath + "\\" + appName + ".jpg");
			apkpath.setApkpath(appPath);
			copy(apkpath.getApkpath(), filePath, "\\" + appName + ".apk");
			ApkInfo apkInfo = new ApkUtil().getApkInfo(apkpath.getApkpath());
			File file = new File(apkpath.getInfoPath());
			long Lenth = new File(apkpath.getApkpath()).length();
			apkInfo.setLenth(Lenth);
			OutputStreamWriter out = null;
			out = new OutputStreamWriter(new FileOutputStream(file), "GBK");
			out.write(apkInfo.getresult());
			out.close();
			extractFileFromApk(apkpath.getApkpath(), apkInfo.getApplicationIcon(), apkpath.getIconPath());
			return true;
		} catch (Exception e) {
			// e.printStackTrace();
			logger.info(e);
		}
		return false;
	}

	public static InputStream extractFileFromApk(String apkpath, String fileName) {
		try {
			@SuppressWarnings("resource")
			ZipFile zFile = new ZipFile(apkpath);
			ZipEntry entry = zFile.getEntry(fileName);
			entry.getComment();
			entry.getCompressedSize();
			entry.getCrc();
			entry.isDirectory();
			entry.getSize();
			entry.getMethod();
			InputStream stream = zFile.getInputStream(entry);
			return stream;
		} catch (IOException e) {
			// e.printStackTrace();
			logger.info(e);
		}
		return null;
	}

	public static void extractFileFromApk(String apkpath, String fileName, String outputPath) throws Exception {
		InputStream is = extractFileFromApk(apkpath, fileName);
		File file = new File(outputPath);
		BufferedOutputStream bos = new BufferedOutputStream(new FileOutputStream(file), 1024);
		byte[] b = new byte[1024];
		BufferedInputStream bis = new BufferedInputStream(is, 1024);
		while (bis.read(b) != -1) {
			bos.write(b);
		}
		bos.flush();
		is.close();
		bis.close();
		bos.close();
	}

	private void copy(String from, String to, String name) throws IOException {
		File file = new File(to);
		// 如果文件夹不存在则创建
		if (!file.exists() && !file.isDirectory()) {
			file.mkdirs();
		} else {
			logger.info("文件" + file.toString() + "已经存在");
		}
		byte[] b = new byte[4096];
		try {
			FileInputStream input = new FileInputStream(from);
			FileOutputStream output = new FileOutputStream(to + name);
			int n = 0;
			n = input.read(b);
			while (n != -1) {
				output.write(b, 0, n);
				n = input.read(b);
			}
			output.flush();
			input.close();
			output.close();
		} catch (IOException e) {
			logger.info(e);
		}
	}

	private void delete(HttpServletRequest request, HttpServletResponse response) {
		String appName = request.getParameter("appName");
		ApkPath base = new ApkPath();
		File f = new File(base.getbasePath() + "\\" + appName + "\\" + appName + ".apk");
		if (f.exists()) {
			f.delete();
		}
		f = new File(base.getbasePath() + "\\" + appName + "\\" + appName + ".txt");
		if (f.exists()) {
			f.delete();
		}
		f = new File(base.getbasePath() + "\\" + appName + "\\" + appName + ".jpg");
		if (f.exists()) {
			f.delete();
		}
	}
}