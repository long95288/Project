import java.lang.reflect.Method;

public class Reflect {
    public static void main(String[] args) {
        String s = "Hello World";
        try{
            // Method 的方法是String类中substring(beginIndex)中的方法
            Method m = String.class.getMethod("substring",int.class);
            // invoke 调用,等价于s.substring(6)
            String sub = (String) m.invoke(s, 6);
            System.out.println(sub);
            String sub2 = s.substring(6);
            System.out.println(sub2);
            
            Method parseInt = Integer.class.getMethod("parseInt", String.class);
            // invoke
            Integer n = (Integer) parseInt.invoke(null, "12346");
            System.out.println("n=" + n);
        }catch(Exception e){
            e.printStackTrace();
        }
       
    }
}