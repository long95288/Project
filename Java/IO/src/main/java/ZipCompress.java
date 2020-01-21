import java.io.*;
import java.util.Enumeration;
import java.util.zip.*;

/**
 * @author longquanxiao
 * @date 2020/1/21
 */
public class ZipCompress {
    public static void main(String[] args) throws IOException {
        FileOutputStream f = new FileOutputStream("test.zip");
        CheckedOutputStream csum = new CheckedOutputStream(f,new Adler32());
        ZipOutputStream zos = new ZipOutputStream(csum);
        BufferedOutputStream out = new BufferedOutputStream(zos);
        zos.setComment("测试zip压缩");
        for (String arg :
                args) {
            System.out.println("写入文件: "+arg);
            BufferedReader in = new BufferedReader(new FileReader(arg));
            zos.putNextEntry(new ZipEntry(arg));
            int c;
            while ((c=in.read())!=-1){
                // 写入压缩文件
                out.write(c);
            }
            in.close();
            out.flush();
        }
        out.close();
        // 校验和
        System.out.println("Checksum: "+ csum.getChecksum().getValue());
        // 解压文件
        System.out.println("解压文件");
        FileInputStream fi = new FileInputStream("test.zip");
        CheckedInputStream csumi = new CheckedInputStream(fi,new Adler32());
        ZipInputStream in2 = new ZipInputStream(csumi);
        BufferedInputStream bis = new BufferedInputStream(in2);
        ZipEntry ze;
        while((ze=in2.getNextEntry())!=null){
            //
            System.out.println("读取Entry:"+ze);
            int x;
            while ((x=bis.read())!=-1){
                System.out.write(x);
            }
        }
        //
        if(args.length == 1){
            System.out.println("校验和："+csumi.getChecksum().getValue());
        }
        //
        bis.close();
        ZipFile zf = new ZipFile("test.zip");
        Enumeration e = zf.entries();
        while (e.hasMoreElements()){
            ZipEntry ze2 = (ZipEntry)e.nextElement();
            System.out.println("File: "+ze2);
        }
    }
}
