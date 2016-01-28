package cn.com.agree.ab.amend.Manager;

import android.content.Context;
import android.util.AttributeSet;
import android.view.MotionEvent;
import android.view.VelocityTracker;
import android.view.View;
import android.view.ViewConfiguration;
import android.view.ViewGroup;
import android.widget.Scroller;

/**
 * Created by Administrator on 2015/10/15.
 */
public class ScrollLayout extends ViewGroup{
    private Scroller mScroller;
    private VelocityTracker mVelocityTracker;

    /**
     * 当前的屏幕位置
     */
    private int mCurScreen;

    /**
     * 设置默认屏幕的属性，0表示第一个屏幕
     */
    private int mDefaultScreen = 0;

    /**
     * 标识滚动操作已结束
     */
    private static final int TOUCH_STATE_REST = 0;
    /**
     * 标识正在执行滑动操作
     */
    private static final int TOUCH_STATE_SCROLLING = 1;

    /**
     * 标识滑动速率
     */
    private static final int SNAP_VELOCITY = 600;

    /**
     * 当前滑动状态
     */
    private int mTouchState = TOUCH_STATE_REST;

    /**
     * 在用户触发ontouch事件之前,我们认为用户能够使view滑动的距离（像素）
     */
    private int mTouchSlop;

    /**
     * 手指触碰屏幕的最后一次x坐标
     */
    private float mLastMotionX;

    /**
     * 手指触碰屏幕的最后一次y坐标
     */
    @SuppressWarnings("unused")
    private float mLastMotionY;

    public ScrollLayout(Context context) {
        super(context);
        mScroller = new Scroller(context);
        mCurScreen = mDefaultScreen;
        mTouchSlop = ViewConfiguration.get(getContext()).getScaledTouchSlop();
    }

    public ScrollLayout(Context context, AttributeSet attrs) {
        super(context, attrs);
        mScroller = new Scroller(context);
        mCurScreen = mDefaultScreen;
        mTouchSlop = ViewConfiguration.get(getContext()).getScaledTouchSlop();
    }

    public ScrollLayout(Context context, AttributeSet attrs, int defStyle) {
        super(context, attrs, defStyle);
        mScroller = new Scroller(context);
        mCurScreen = mDefaultScreen;
        mTouchSlop = ViewConfiguration.get(getContext()).getScaledTouchSlop();


    }

    @Override
    protected void onLayout(boolean changed, int l, int t, int r, int b) {
        if (changed) {
            int childLeft = 0;
            final int childCount = getChildCount();

            for (int i = 0; i < childCount; i++) {
                final View childView = getChildAt(i);
                if (childView.getVisibility() != View.GONE) {
                    final int childWidth = childView.getMeasuredWidth();
                    childView.layout(childLeft, 0, childLeft + childWidth,
                            childView.getMeasuredHeight());
                    childLeft += childWidth;
                }
            }
        }
    }

    @Override
    protected void onMeasure(int widthMeasureSpec, int heightMeasureSpec) {
        super.onMeasure(widthMeasureSpec, heightMeasureSpec);

        final int width = MeasureSpec.getSize(widthMeasureSpec);
        final int widthMode = MeasureSpec.getMode(widthMeasureSpec);
        if (widthMode != MeasureSpec.EXACTLY) {
            throw new IllegalStateException(
                    "ScrollLayout only canmCurScreen run at EXACTLY mode!");
        }

        final int heightMode = MeasureSpec.getMode(heightMeasureSpec);
        if (heightMode != MeasureSpec.EXACTLY) {
            throw new IllegalStateException(
                    "ScrollLayout only can run at EXACTLY mode!");
        }

        final int count = getChildCount();
        for (int i = 0; i < count; i++) {
            getChildAt(i).measure(widthMeasureSpec, heightMeasureSpec);
        }
        // 初始化视图的位置
        scrollTo(mCurScreen * width, 0);
    }

    /**
     * 根据滑动的距离判断移动到第几个视图
     */
    public void snapToDestination() {
        final int screenWidth = getWidth();
        final int destScreen = (getScrollX() + screenWidth / 2) / screenWidth;
        snapToScreen(destScreen);
    }

    /**
     * 滚动到制定的视图
     *
     * @param whichScreen
     * 视图下标
     */
    public void snapToScreen(int whichScreen) {
        whichScreen = Math.max(0, Math.min(whichScreen, getChildCount() - 1));
        if (getScrollX() != (whichScreen * getWidth())) {

            final int delta = whichScreen * getWidth() - getScrollX();
            mScroller.startScroll(getScrollX(), 0, delta, 0, 1000);
            mCurScreen = whichScreen;
            invalidate();
        }
    }

    public void setToScreen(int whichScreen) {
        whichScreen = Math.max(0, Math.min(whichScreen, getChildCount() - 1));
        mCurScreen = whichScreen;
        scrollTo(whichScreen * getWidth(), 0);
    }

    public int getCurScreen() {
        return mCurScreen;
    }

    @Override
    public void computeScroll() {
        if (mScroller.computeScrollOffset()) {
            scrollTo(mScroller.getCurrX(), mScroller.getCurrY());
            postInvalidate();
        }
    }

    @Override
    public boolean onTouchEvent(MotionEvent event) {
        if (mVelocityTracker == null) {
            mVelocityTracker = VelocityTracker.obtain();
        }
        mVelocityTracker.addMovement(event);

        final int action = event.getAction();
        final float x = event.getX();

        switch (action) {
            case MotionEvent.ACTION_DOWN:
                if (!mScroller.isFinished()) {
                    mScroller.abortAnimation();
                }
                mLastMotionX = x;
                break;

            case MotionEvent.ACTION_MOVE:
                int deltaX = (int) (mLastMotionX - x);
                mLastMotionX = x;

                scrollBy(deltaX, 0);
                break;

            case MotionEvent.ACTION_UP:
                final VelocityTracker velocityTracker = mVelocityTracker;
                velocityTracker.computeCurrentVelocity(1000);
                int velocityX = (int) velocityTracker.getXVelocity();

                if (velocityX > SNAP_VELOCITY && mCurScreen > 0) {
                    // 向左移动
                    snapToScreen(mCurScreen - 1);
                } else if (velocityX < -SNAP_VELOCITY
                        && mCurScreen < getChildCount() - 1) {
                    // 向右移动
                    snapToScreen(mCurScreen + 1);
                } else {
                    snapToDestination();
                }
                if (mVelocityTracker != null) {
                    mVelocityTracker.recycle();
                    mVelocityTracker = null;
                }
                mTouchState = TOUCH_STATE_REST;
                break;
            case MotionEvent.ACTION_CANCEL:
                mTouchState = TOUCH_STATE_REST;
                break;
        }

        return true;
    }

    @Override
    public boolean onInterceptTouchEvent(MotionEvent ev) {
        final int action = ev.getAction();
        if ((action == MotionEvent.ACTION_MOVE)
                && (mTouchState != TOUCH_STATE_REST)) {
            return true;
        }

        final float x = ev.getX();
        final float y = ev.getY();

        switch (action) {
            case MotionEvent.ACTION_MOVE:
                final int xDiff = (int) Math.abs(mLastMotionX - x);
                if (xDiff > mTouchSlop) {
                    mTouchState = TOUCH_STATE_SCROLLING;
                }
                break;

            case MotionEvent.ACTION_DOWN:
                mLastMotionX = x;
                mLastMotionY = y;
                mTouchState = mScroller.isFinished() ? TOUCH_STATE_REST
                        : TOUCH_STATE_SCROLLING;
                break;

            case MotionEvent.ACTION_CANCEL:
            case MotionEvent.ACTION_UP:
                mTouchState = TOUCH_STATE_REST;
                break;
        }

        return mTouchState != TOUCH_STATE_REST;
    }

}
