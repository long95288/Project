package concurrency;

import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.ReentrantLock;

/**
 * @author longquanxiao
 * @date 2019/11/6
 */
public class AttemptLocking {
    private ReentrantLock lock = new ReentrantLock();
    public void untimed(){
        boolean capture = lock.tryLock();
        try{
            System.out.println("try Lock() " + capture);
        }finally {
            if(capture){
                lock.unlock();
            }
        }
    }
    public void timed(){
        boolean capture = false;
        try{
            capture = lock.tryLock(2, TimeUnit.SECONDS);
        }catch (InterruptedException e){
            throw new RuntimeException();
        }

        try{
            System.out.println("tryLock(2,TimeUnit.SECONDS): "+capture);
        }finally {
            if(capture) lock.unlock();
        }
    }

    public static void main(String[] args) {
        final AttemptLocking al = new AttemptLocking();
        al.untimed();
        al.timed();
        new Thread(){
            {
                setDaemon(true);
            }
            @Override
            public void run() {
                al.lock.lock();
                System.out.println("acquired");
            }
        }.start();
        Thread.yield();
        al.untimed();
        al.timed();
    }
}
