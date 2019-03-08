public class ex15_3 {
    public static void main(String[] args)
    {
        // 两个线程的目标对象相同
        Blank bank=new Blank();
        bank.setMoney(300);
        Thread one,two;
        one=new Thread(bank);
        one.setName("One");
        two = new Thread(bank);
        two.setName("Two");
        one.start();
        two.start();
    }
}
