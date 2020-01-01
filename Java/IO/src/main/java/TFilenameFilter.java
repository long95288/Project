import java.io.File;

/**
 * @author longquanxiao
 * @date 2019/12/31
 */
public class TFilenameFilter {
    public static void main(String[] args) throws Exception{
        File f = new File(".\\pom.xml");
        System.out.println(f.getParent());
        System.out.println(f.getName());
        System.out.println(f.getPath());
        System.out.println(f.getAbsolutePath());
        System.out.println(f.getCanonicalPath());
        System.out.println(f.length());
        System.out.println(f.lastModified());
        File f2 = new File("E:\\test\\test2\\test3\\test4\\test5.txt");
        if(!f2.exists()){
            boolean c2 = f2.createNewFile();
            System.out.println(c2);
//            try{
//                boolean c = f2.mkdirs();
//                System.out.println(c);
//            }catch (Exception e){
//                e.printStackTrace();
//            }

        }
    }
}
