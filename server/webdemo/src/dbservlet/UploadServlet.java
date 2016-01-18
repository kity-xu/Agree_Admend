package dbservlet;

import java.io.*;
import java.text.DecimalFormat;
import java.util.*;

import javax.servlet.RequestDispatcher;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import org.apache.commons.fileupload.FileItem;
import org.apache.commons.fileupload.ProgressListener;
import org.apache.commons.fileupload.disk.DiskFileItemFactory;
import org.apache.commons.fileupload.servlet.ServletFileUpload;
import org.apache.log4j.Logger;

public class UploadServlet extends HttpServlet {

	/**
	 * 
	 */
	private static final long serialVersionUID = -2530384139978297534L;
	private static Logger logger = Logger.getLogger(UploadServlet.class);
	private boolean isMultipart;
	private String filePath;
	private int maxFileSize = 300 * 1000 * 1024;
	private int maxMemSize = 1000 * 1024;
	private String TempFilePath = "f:/temp2/";
	private File file;

	public void init() {
		// Get the file location where it would be stored.
		filePath = getServletContext().getInitParameter("file-upload");
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, java.io.IOException {
		// Check that we have a file upload request
		isMultipart = ServletFileUpload.isMultipartContent(request);
		if (!isMultipart) {
			return;
		} else {
			request.getSession().setAttribute("IS_UPLOAD_BEGIN", 1);
			request.getSession().setAttribute("UPLOAD_PERCENTAGE", 0);
		}
		DiskFileItemFactory factory = new DiskFileItemFactory();
		// maximum size that will be stored in memory
		factory.setSizeThreshold(maxMemSize);
		// Location to save data that is larger than maxMemSize.
		factory.setRepository(new File(TempFilePath));

		// Create a new file upload handler
		ServletFileUpload upload = new ServletFileUpload(factory);
		// maximum file size to be uploaded.
		upload.setSizeMax(maxFileSize);

		class UploadProgressListener implements ProgressListener {
			private HttpServletRequest request;
			private DecimalFormat df = new DecimalFormat("#00.0");

			UploadProgressListener(HttpServletRequest request) {
				this.request = request;
			}

			public void update(long bytesRead, long bytesTotal, int items) {
				double percent = (double) bytesRead * 100 / (double) bytesTotal;
				request.getSession().setAttribute("UPLOAD_PERCENTAGE", df.format(percent));
			}
		}
		upload.setProgressListener(new UploadProgressListener(request));
		try {
			// Parse the request to get file items.
			List<?> fileItems = upload.parseRequest(request);

			// Process the uploaded file items
			Iterator<?> i = fileItems.iterator();
			while (i.hasNext()) {
				FileItem fi = (FileItem) i.next();
				if (!fi.isFormField()) {
					fi.getFieldName();
					String fileName = fi.getName();
					fi.getContentType();
					fi.isInMemory();
					fi.getSize();
					// Write the file
					if (fileName.lastIndexOf("\\") >= 0) {
						file = new File(filePath + fileName.substring(fileName.lastIndexOf("\\")));
					} else {
						file = new File(filePath + fileName.substring(fileName.lastIndexOf("\\") + 1));
					}
					fi.write(file);
					request.setAttribute("filename", fileName);
					logger.info("fileName:" + fileName);
				}
			}
			request.getSession().setAttribute("IS_UPLOAD_BEGIN", 0);
			RequestDispatcher rd1 = request.getRequestDispatcher("appfile/appinput.jsp");
			rd1.forward(request, response);
		} catch (Exception e) {
			logger.info(e);
		}
	}

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, java.io.IOException {
		logger.info("doGet");
		doPost(request, response);
	}
}
