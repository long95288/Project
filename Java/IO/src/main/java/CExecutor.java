import java.io.*;
import java.util.UUID;

/**
 * @author longquanxiao
 * @date 2020/2/27
 */
public class CExecutor extends TestOJExecutor{

    private static final String BASE_PATH = "D://Temp/";
    @Override
    public String run(String code) {
        String filename = "C"+ UUID.randomUUID().toString().replaceAll("-","")+".c";
        String filepath = BASE_PATH + filename;
        File cFile = new File(filepath);
        try{
            if(!cFile.exists()){cFile.createNewFile();}
//            FileWriter fileWriter = new FileWriter(cFile);
            // 写入为GBK才能显示中文....
            BufferedWriter fileWriter = new BufferedWriter(new OutputStreamWriter(new FileOutputStream(cFile),"GBK"));
            fileWriter.write(code);
            fileWriter.flush();
            fileWriter.close();
            // 编译文件
            String compileCommand = "gcc "+filepath+" -o "+BASE_PATH+filename.split("\\.")[0];
            System.out.println("compile: "+compileCommand);
            Process process = new ProcessBuilder(compileCommand.split(" ")).start();
            // 获得错误流
            BufferedReader errReader = new BufferedReader(new InputStreamReader(process.getErrorStream(),"GBK"));
            // 获得结果流
            BufferedReader resultReader = new BufferedReader(new InputStreamReader(process.getInputStream(),"GBK"));

            String print;
            while ((print = errReader.readLine())!= null){
                System.out.println(print);
            }
            while ((print = resultReader.readLine()) != null){
                System.out.println(print);
            }
            // 运行文件
            String runCommand = BASE_PATH+filename.split("\\.")[0]+".exe";
            System.out.println("run: "+runCommand);
            process = new ProcessBuilder(runCommand.split(" ")).start();
            // 获得错误流
            errReader = new BufferedReader(new InputStreamReader(process.getErrorStream(),"GBK"));
            // 获得结果流
            resultReader = new BufferedReader(new InputStreamReader(process.getInputStream(),"GBK"));

            while ((print = errReader.readLine())!= null){
                System.out.println(print);
            }
            while ((print = resultReader.readLine()) != null){
                System.out.println(print);
            }
        }catch (IOException e){

        }finally {

        }
        return null;
    }

    public static void main(String[] args) {
        String code = "#include<stdio.h>\n" +
                "int main(){\n" +
                "  printf(\"this is C program \\n\");\n" +
                "  //\n" +
                "  //\n" +
                "  printf(\"这是C程序 \\n\");\n" +
                "  return 0;\n" +
                "}";
        new CExecutor().run(code);
    }
}
