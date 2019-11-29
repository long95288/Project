/**
 * @author longquanxiao
 * @date 2019/11/29
 * 匹配字符序列
 */
public class Rudolph {
    public static void main(String[] args) {
        for (String pattern: new String[]{"Rudolph","[rR]udolph","[rR][aeiou][a-z]0l.*","R.*"}){
            System.out.println("Rudolph".matches(pattern));
        }
    }
}
