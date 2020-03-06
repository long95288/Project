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
            fileWriter.close();
            // 2.运行文件
            String command = "python " + filepath;
            System.out.println(command);
            Process process = new ProcessBuilder(command.split(" ")).start();
            // 获得错误流
            BufferedReader errReader = new BufferedReader(new InputStreamReader(process.getErrorStream(),"GBK"));
            // 获得结果流
            BufferedReader resultReader = new BufferedReader(new InputStreamReader(process.getInputStream(),"GBK"));
            // 获得输出流 控制台的输入
            BufferedWriter writer = new BufferedWriter(new OutputStreamWriter(process.getOutputStream(),"GBK"));
            String inputData = "这是向控制台输入的数据1....";
            // 换行符是必须的!!!!,否则控制台无法知道已经完成输入
            writer.write(inputData + "\n");
            inputData = "这是向控制台输入的数据2...";
            writer.write(inputData+"\n");
            writer.flush();
            writer.close();

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
                "    print(\"这是控制台输出的数据....\")\n" +
                "    # 读取控制台的输入数据\n" +
                "    print(\"读取第一组数据\")\n" +
                "    fff = input()\n"+
                "    print(fff)\n" +
                "    print(\"读取第二组数据\")\n" +
                "    fff2 = input()\n" +
                "    print(fff2)\n" +
                "    print(fff3)\n";
        new PythonExecutor().run(code);
    }
}
