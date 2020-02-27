import com.sun.xml.internal.ws.policy.privateutil.PolicyUtils;

import java.io.*;
import java.util.Random;
import java.util.UUID;

/**
 * @author longquanxiao
 * @date 2020/2/27
 */
public class PythonExecutor extends TestOJExecutor{

    /**
     * 文件保存的位置
     */
    private static final String BASE_PATH = "D://Temp/";

    @Override
    public String run(String code) {
        // 1.写入文件
        String filename = "python"+ UUID.randomUUID().toString().replaceAll("-","") + ".py";
        String filepath = BASE_PATH  + filename;
        File pythonFile = new File(filepath);
        try{
            if (!pythonFile.exists()){
                pythonFile.createNewFile();
            }
            FileWriter fileWriter = new FileWriter(pythonFile);
            fileWriter.write(code);
            fileWriter.flush();
            System.out.println("flush ...");
            fileWriter.close();
            // 2.运行文件
            String command = "python " + filepath;
            System.out.println(command);
            Process process = new ProcessBuilder(command.split(" ")).start();
            // 获得错误流
            BufferedReader errReader = new BufferedReader(new InputStreamReader(process.getErrorStream(),"GBK"));
            // 获得结果流
            BufferedReader resultReader = new BufferedReader(new InputStreamReader(process.getInputStream(),"GBK"));
            // 获得输入流
            System.out.println("获得输入流...");

            String print = null;
            while ((print = resultReader.readLine()) != null){
                System.out.println(print);
            }

            while((print = errReader.readLine())!= null){
                System.out.println(print);
            }
            //
        }catch (IOException e){
            //
            e.printStackTrace();
            return null;
        }

        // 3.返回结果
        return null;
    }

    public static void main(String[] args) {
        String code = "\n" +
                "import time\n" +
                "if __name__ == '__main__':\n" +
                "    print(\"This is python program\")\n" +
                "    print(\"这是python程序\")\n" +
                "    print(fff)";
        new PythonExecutor().run(code);
    }
}
