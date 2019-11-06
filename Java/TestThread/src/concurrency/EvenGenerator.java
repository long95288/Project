package concurrency;

/**
 * @author longquanxiao
 * @date 2019/11/6
 */
public class EvenGenerator extends IntGenerator{
    private int currentEvenValue = 0;
    @Override
    public int next() {
        // 获得下一个偶数
        ++currentEvenValue;
        ++currentEvenValue;
        return currentEvenValue;
    }

    public static void main(String[] args) {
        EvenChecker.test(new EvenGenerator());
    }
}
