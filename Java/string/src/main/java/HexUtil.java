/**
 * @author longquanxiao
 * @date 2019/12/19
 */
public class HexUtil {
    /**
     * 16进制显示工具
     * @param data
     * @return
     */
    public static String format(byte[] data) {
        StringBuilder result = new StringBuilder();
        int n = 0;
        for (byte b :
                data) {
            if (n % 16 == 0) {
                result.append(String.format("%05X: ", n));
            }
            result.append(String.format("%02X: ", b));
            n++;
        }
        if (n % 16 == 0) result.append("\n");
        return result.toString();
    }

    public static void main(String[] args) {
        if(args.length == 0){
            System.out.println("");
        }
    }
}
