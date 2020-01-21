import java.io.*;
import java.util.zip.GZIPInputStream;
import java.util.zip.GZIPOutputStream;

/**
 * @author longquanxiao
 * @date 2020/1/21
 */
public class GZIPcompress {
    public static void main(String[] args) throws Exception {
        if (args.length == 0){
            System.out.println("请输入待压缩的文件的名称");
            System.exit(1);
        }
        BufferedReader in = new BufferedReader(new FileReader(args[0]));
        BufferedOutputStream out = new BufferedOutputStream(new GZIPOutputStream(new FileOutputStream("test.gz")));
        System.out.println("写入文件");
        int c;
        while ((c=in.read()) !=-1){
            out.write(c);
        }
        in.close();
        out.close();
        System.out.println("读入文件");
        BufferedReader in2 = new BufferedReader(
                new InputStreamReader(
                        new GZIPInputStream(
                                new FileInputStream("test.gz")
                        )
                )
        );
        String s;
        while((s=in2.readLine())!=null){
            System.out.println(s);
        }

    }
}
