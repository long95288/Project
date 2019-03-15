public class hfc {

    public int hfc(int m,int n)
    {
        int r =1;
        r = m % n;
        while ( r != 0)
        {
            m = n;
            n = r;
            r = m % n;
        }
        return n;
    }
    public static void main(String[] args)
    {
        hfc c = new hfc();
        int hfc1 = c.hfc(377,246);
        System.out.println("377,246的最大公因子为:"+hfc1);
        System.out.println("1000,150的最大公因子为:"+c.hfc(1000,150));
    }
}
