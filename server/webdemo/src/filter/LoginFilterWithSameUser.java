package filter;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Iterator;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.annotation.WebFilter;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpSession;

import org.apache.log4j.Logger;

/**
 * Servlet Filter implementation class LoginFilterWithSameUser
 */
@WebFilter("/LoginFilterWithSameUser")
public class LoginFilterWithSameUser implements Filter {

	// 已登陆列表
	public static ArrayList<HttpSession> sessionList = new ArrayList<HttpSession>();
	private static Logger logger = Logger.getLogger(LoginFilterWithSameUser.class);

	/**
	 * Default constructor.
	 */
	public LoginFilterWithSameUser() {
	}

	/**
	 * @see Filter#destroy()
	 */
	public void destroy() {
	}

	/**
	 * @see Filter#doFilter(ServletRequest, ServletResponse, FilterChain)
	 */
	public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
			throws IOException, ServletException {
		HttpServletRequest req = (HttpServletRequest) request;
		HttpSession session = req.getSession();
		String newName = (String) session.getAttribute("userName");
		logger.info(newName + " is login");
		if (newName != null) {// 当前session已登录
			Iterator<HttpSession> li = sessionList.iterator();
			while (li.hasNext()) {
				HttpSession tmpSession = li.next();
				try {
					String tmpName = (String) tmpSession.getAttribute("userName");
					if (tmpName == null) {
						li.remove();
					} else {
						if (tmpName.equals(newName)) {
							logger.info(newName + " is login again");
							if (tmpSession == session) {
								// 未退出就再次登陆
								logger.info(tmpSession.getId() + " is not change");
							} else {
								tmpSession.setAttribute("userName", null);
								logger.info(tmpSession.getId() + " is change");
							}
							li.remove();//
						} else {
							// 其他正常在线的用户
							logger.info(tmpName + " is not online");
						}
					}
				} catch (Exception e) {
					li.remove();
					logger.info(tmpSession.getId() + " is timeout");
				}
			}
			sessionList.add(session);
		}
		chain.doFilter(request, response);

	}

	/**
	 * @see Filter#init(FilterConfig)
	 */
	public void init(FilterConfig fConfig) throws ServletException {
	}

}
