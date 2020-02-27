import java.io.File;
import java.io.FileWriter;

/**
 * @author longquanxiao
 * @date 2020/1/13
 */
public class OSExecuteDemo {

    public static void main(String[] args) throws Exception {
        String basePath = "D://Temp/";
        String pythonFilename = "Hello2.py";
        File python = new File(basePath + pythonFilename);
        if (!python.exists()) python.createNewFile();

        FileWriter fileWriter = new FileWriter(python);
        String code = "\n" +
                "if __name__ == '__main__':\n" +
                "    print(\"This is python program\")\n" +
                "    print(\"这是python程序\")\n" +
                "\n" +
                "    print(fff)" +
                "";
        fileWriter.write(code);
        //
        fileWriter.close();
        // 写入完成之后运行
        OSExecute.command("python "+basePath+pythonFilename);

        // python 环境运行
        System.out.println("测试python环境");
         OSExecute.command("python D://Temp/hello.py");
        // java 环境运行
        // 1.编译
        OSExecute.command("javac -encoding UTF-8 D://Temp/Hello.java");
        // 2.运行
        OSExecute.command("java -classpath D://Temp/ Hello");
        System.out.println("测试C编译运行环境");
        // C++ 环境运行
        // 1.编译
        OSExecute.command("gcc D://Temp/TestC/test.c -o test");
        // 2.运行
        OSExecute.command("D://Temp/TestC/test.exe");
        // OSExecute.command("F://tools/blibli_image_download/start_get_blibli_image.bat");
        // C++环境
        System.out.println("测试C++编译运行环境");
        OSExecute.command("g++ D://Temp/HelloCpp.cpp -o HelloCpp");
        OSExecute.command("D://Temp/HelloCpp.exe");
//        OSExecute.command("node");
    }
}
