public class Blank implements Runnable {
    private int number=0;
    public void setMoney(int n)
    {
        number=n;
    }
    @Override
    public void run() {
        while (true)
        {
            String name = Thread.currentThread().getName();
            if (name.equals("one"))
            {
                if (number<=160)
                {
                    System.out.println(name+"进入死亡状态");
                    return;
                }
               number+=10;
                System.out.println("我是"+name+"现在number"+number);
            }

            if (Thread.currentThread().getName().equals("Two"));
            {
                if (number<=0)
                {
                    System.out.println(name+"进入死亡状态");
                    return;
                }
                number-=100;
                System.out.println("我是"+name+"现在number="+number);

            }

            try {
                Thread.sleep(800);
            }catch (InterruptedException e)
            {
            }
        }
    }
}
