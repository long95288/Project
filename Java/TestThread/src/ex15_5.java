public class ex15_5 {
    public static void main(String[] args)
    {
        ComputerSum computer = new ComputerSum();
        Thread threadone;
        threadone= new Thread(computer);
        threadone.setName("张三");
        threadone.start();
    }
}
