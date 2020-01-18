import java.io.FileOutputStream;
import java.io.IOException;
import java.nio.channels.FileLock;
import java.util.concurrent.TimeUnit;

/**
 * @author longquanxiao
 * @date 2020/1/18
 */
public class FileLocking {
    public static void main(String[] args) throws Exception {
        FileOutputStream fos = new FileOutputStream("file.txt");
        FileLock fl = fos.getChannel().tryLock();
        if (fl != null ){
            System.out.println("Locked File");
            TimeUnit.MILLISECONDS.sleep(100);
            System.out.println("Released Lock");
        }
        fos.close();
    }
}
