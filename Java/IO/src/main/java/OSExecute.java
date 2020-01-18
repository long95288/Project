import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * @author longquanxiao
 * @date 2020/1/13
 */
public class OSExecute {
    public static void command(String command){
        boolean err = false;
        try{
            Process process = new ProcessBuilder(command.split(" ")).start();
            BufferedReader results = new BufferedReader(new InputStreamReader(process.getInputStream()));
            String s = null;
            while ((s=results.readLine())!=null){
                // 打印运行的结果
                System.out.println(s);
            }
        }catch (IOException e){
            e.printStackTrace();
        }
    }

}
