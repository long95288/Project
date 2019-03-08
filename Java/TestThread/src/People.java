public class People extends Thread{
    StringBuffer str;
    People(String s,StringBuffer str)
    {
        setName(s); //给线程赋值一个名字
        this.str = str;
    }
    public void run(){
        for(int i=1;i<=3;i++)
        {
            str.append(getName()+",");
            System.out.println("我是"+getName()+",字符串为:"+str);
            try{
                sleep(200);
            }catch (InterruptedException e){


            }
        }
    }
}
