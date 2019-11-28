import java.util.Arrays;

/**
 * @author longquanxiao
 * @date 2019/11/28
 * 正则切割
 */
public class Splitting {
    public static String knights =
            "Then,when you have found the shrubbery, you must"+
                    "cut down the mightiest tree in the forest"+
                    "with ...a herring";
    public static void split(String regex){
        System.out.println(
                Arrays.toString(knights.split(regex))
        );
    }

    public static void main(String[] args) {
        split(" ");
        split("\\W+");
        split("n\\W+");
    }
}
