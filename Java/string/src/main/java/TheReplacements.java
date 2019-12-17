/**
 * @author longquanxiao
 * @date 2019/12/16
 */
public class TheReplacements {
    public static void main(String[] args) throws Exception{
        String str = "a b c d e f g a b c d e f g";

        String str1 = str.replaceAll("a","a1");
        System.out.println(str1);

        String str2 = str.replaceFirst("b","b1");
        System.out.println(str2);
    }
}
