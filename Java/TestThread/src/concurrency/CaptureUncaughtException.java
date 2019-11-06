package concurrency;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.ThreadFactory;

/**
 * @author longquanxiao
 * @date 2019/11/6
 */
class ExceptionThread2 implements Runnable {
    @Override
    public void run() {
        Thread t = Thread.currentThread();
        System.out.println("run() by "+ t);
        System.out.println("" +
                "eh="+t.getUncaughtExceptionHandler());
        throw new RuntimeException();
    }
}

/**
 * 自定义的线程未捕获的异常类
 * 用来处理线程逃逸的异常
 */
class MyUncaughtExceptionHandler implements Thread.UncaughtExceptionHandler {
    @Override
    public void uncaughtException(Thread t, Throwable e) {
        // 捕获的异常要在这里设置
        System.out.println("Caught " + e);
    }
}

/**
 * 线程工厂
 * 用来产生一个设置了未捕获异常处理的线程
 */
class HandlerThreadFactory implements ThreadFactory {
    @Override
    public Thread newThread(Runnable r) {
        System.out.println(this + " creating new Thread");
        Thread t = new Thread(r);
        System.out.println("created "+ t);
        // 给新建立的线程设置一个未捕获异常处理实列
        t.setUncaughtExceptionHandler(new MyUncaughtExceptionHandler());
        System.out.println("eh = "+ t.getUncaughtExceptionHandler());
        return t;
    }
}
public class CaptureUncaughtException {
    public static void main(String[] args) {
        ExecutorService exec = Executors.newCachedThreadPool(new HandlerThreadFactory());
        exec.execute(new ExceptionThread2());

    }
}
