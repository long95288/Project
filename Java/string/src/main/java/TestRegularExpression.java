import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * @author longquanxiao
 * @date 2019/11/29
 */
public class TestRegularExpression {
    public static void main(String[] args) {
        if (args.length < 2){
            System.out.println("Usage: \njava TestRegularExpression characterSequence regularExpression");
            System.exit(0);
        }
        System.out.println("Input:\""+args[0]+"\"");
        for (String arg: args){
            System.out.println("Regular expression:\""+arg+"\"");
            Pattern p = Pattern.compile(arg);
            Matcher m = p.matcher(args[0]);
            while (m.find()){
                System.out.println("Match \""+ m.group()+"\" at positions "+
                        m.start() + "-" +(m.end() -1));
            }
        }
    }
}
