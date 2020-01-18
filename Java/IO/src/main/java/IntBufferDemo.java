import java.nio.ByteBuffer;
import java.nio.IntBuffer;

/**
 * @author longquanxiao
 * @date 2020/1/17
 * 整型数据直接通过视图写入通道
 */
public class IntBufferDemo {
    /**
     * 缓冲区的大小
     */
    private static final int BSIZE = 1024;
    public static void main(String[] args) {
        // 获得字节缓冲区
        ByteBuffer bb = ByteBuffer.allocate(BSIZE);
        // 获得整型的视图缓冲器
        IntBuffer ib = bb.asIntBuffer();
        // 将整型数组写入缓冲器中
        ib.put(new int[]{11,42,47,99,143,811,1016});
        // 测试通过视图获得数据
        // 输出99
        System.out.println(ib.get(3));
        // 写入
        ib.flip();
        // 全部读出来
        while (ib.hasRemaining()){
            int i = ib.get();
            System.out.println(i);
        }
    }
}
