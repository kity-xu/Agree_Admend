package filter;

import java.io.IOException;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.annotation.WebFilter;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

import org.apache.log4j.Logger;

/**
 * Servlet Filter implementation class LoginFilter
 */
@WebFilter(description = "检测是否登录", urlPatterns = { "/LoginFilter" })
public class LoginFilterWithJump implements Filter {
	private static Logger logger = Logger.getLogger(LoginFilterWithJump.class);

	/**
	 * Default constructor.
	 */
	public LoginFilterWithJump() {
	}

	/**
	 * @see Filter#destroy()
	 */
	public void destroy() {
	}

	/**
	 * @see Filter#doFilter(ServletRequest, ServletResponse, FilterChain)
	 */
	@Override
	public void doFilter(ServletRequest req, ServletResponse res, FilterChain chain)
			throws IOException, ServletException {
		HttpServletRequest request = (HttpServletRequest) req;
		HttpServletResponse response = (HttpServletResponse) res;
		HttpSession session = request.getSession();
		if (session.getAttribute("userName") == null) {
			logger.info(session.getId() + "未登录,并跳转");
			response.setHeader("Refresh", "5;URL=/webdemo/login.jsp");
			response.sendError(401, "未登录,5秒后进入登录页面");
		} else {
			chain.doFilter(req, res);
		}

	}

	/**
	 * @see Filter#init(FilterConfig)
	 */
	public void init(FilterConfig fConfig) throws ServletException {
	}

}