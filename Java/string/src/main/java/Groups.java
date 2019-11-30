import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * @author longquanxiao
 * @date 2019/11/30
 */
public class Groups {
    public static final String POEM =
            "Twas brillig, and the slithy toves\n"+
                    "Did gyre an gimble in the wabe.\n"+
                    "All mimsy were the borogoves\n"+
                    "And the mome raths outgrabe.\n"+
                    "Beware th Jabberwork, my son .\n"+
                    "The jaws that bite,the claws that catch.\n"+
                    "The frumious Bandersnatch.";
    public static void main(String[] args) {
        Matcher m = Pattern.compile("(?m)(\\S+)\\s+((\\S+)\\s+(\\S+))$")
                .matcher(POEM);
        while (m.find()){
            for (int i = 0; i < m.groupCount(); i++) {
                System.out.println("["+ m.group(i)+" ]");
            }
            System.out.println();
        }
    }
}
