package concurrency;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/**
 * 测试未捕获异常，
 * @author longquanxiao
 * @date 2019/11/6
 */
public class ExceptionThread implements Runnable{
    @Override
    public void run() {
        // 向外抛出异常
        throw new RuntimeException();
    }

    public static void main(String[] args) {
        ExecutorService exec = Executors.newCachedThreadPool();
        try {
            exec.execute(new ExceptionThread());
        }catch (RuntimeException e){
            // 不能捕获线程抛出的异常
            System.out.println(e.getMessage());
            System.out.println("异常被处理");
        }
    }
}
