public class ComputerSum implements Runnable{

    int i=0,sum=0;

    public void run() {
       Thread thread = Thread.currentThread();
       System.out.println(thread.getName()+"开始计算");
       while (i<=10)
       {
           sum+=i;
           System.out.println(" "+sum);
           if (i==5){
               System.out.println(thread.getName()+"完成任务! i="+i);
               Thread Two = new Thread(this);
               Two.setName("李四");
               Two.start();
               i++;
               return;
           }
           i++;
           try{Thread.sleep(300);}
           catch (InterruptedException e){}
       }
    }
}
