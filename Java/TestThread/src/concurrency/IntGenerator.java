package concurrency;

/**
 * @author longquanxiao
 * @date 2019/11/6
 */
public abstract class IntGenerator {
    /**
     * 是否被取消的标志
     */
    private volatile boolean canceled = false;

    /**
     * 获得下一个整数
     * @return
     */
    public abstract int next();

    /**
     * 取消
     */
    public void cancel(){canceled = true;}

    /**
     * 判断是否被取消了
     * @return
     */
    public boolean isCanceled() {
        return canceled;
    }
}
