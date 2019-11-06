package concurrency;

/**
 * next 加上同步锁之后数据变得可控了
 * @author longquanxiao
 * @date 2019/11/6
 */
public class SynchronizedEvenGenerator extends IntGenerator{
    private int currentEvenValue = 0;

    @Override
    public synchronized int next() {
        ++currentEvenValue;
        Thread.yield();
        ++currentEvenValue;
        return currentEvenValue;
    }

    public static void main(String[] args) {
        EvenChecker.test(new SynchronizedEvenGenerator());
    }
}
