public class ex15_1 {

    public static void main(String args[])
    {
        SpearkHello hello;
        SpeakNihao nihao;
        hello = new SpearkHello();  // 创建进程
        nihao = new SpeakNihao();  // 创建进程
        hello.start();  // 启动进程
        nihao.start();  // 启动进程
        for(int i=0;i<20;i++){
            System.out.print("主进程"+i);
        }
        System.out.println("程序运行结束");

    }
}
