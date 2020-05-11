import java.util.Scanner;

/**
 * @author longquanxiao
 * @date 2020/4/29
 */
public class TestInputJava {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        String input = in.nextLine();
        if (null != input && "" != input){
            String[] data = input.split(" ");
            System.out.println(Integer.parseInt(data[0]) + Integer.parseInt(data[1]));
        }
    }
}
