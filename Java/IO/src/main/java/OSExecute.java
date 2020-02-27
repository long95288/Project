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
            BufferedReader errReader = new BufferedReader(new InputStreamReader(process.getErrorStream(),"GBK"));
            BufferedReader results = new BufferedReader(new InputStreamReader(process.getInputStream(),"GBK"));
            String s = null;
            // 打印正常的结果
            while ((s=results.readLine())!=null){
                // 打印运行的结果
                System.out.println(s);
            }
            // 打印错误信息
            while((s=errReader.readLine())!=null){
                System.out.println(s);
            }
        }catch (IOException e){
            System.out.println(e.getMessage());
            e.printStackTrace();
        }catch (Exception e){
            System.out.println(e.getMessage());
        }
    }

}
