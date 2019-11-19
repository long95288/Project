/**
 * @author longquanxiao
 * @date 2019/11/19
 */
public class WhiterStringBuilder {
    public String implicit(String[] fields){
        String result = "";
        for (int i = 0; i < fields.length; i++) {
            result += fields[i];
        }
        return result;
    }

    public String explicit(String[] fields){
        StringBuilder result = new StringBuilder();
        for (int i = 0; i < fields.length; i++) {
            result.append(fields[i]);
        }
        return result.toString();
    }
}
