import java.nio.ByteBuffer;
import java.nio.CharBuffer;

/**
 * @author longquanxiao
 * @date 2020/1/17
 * 交换相邻字符
 */
public class UsingBuffers {
    private static void symmetricScramble(CharBuffer buffer){
       while (buffer.hasRemaining()){
           // 做标记
           buffer.mark();
           char c1 = buffer.get();
           char c2 = buffer.get();
           // 回到标记点
           buffer.reset();
           // 交换次序存储
           buffer.put(c2).put(c1);
       }
    }

    public static void main(String[] args) {
        char[] data = "UsingBuffers".toCharArray();
        // 一个字符占有两个字节
        ByteBuffer bb = ByteBuffer.allocate(data.length*2);
        CharBuffer cb = bb.asCharBuffer();
        cb.put(data);
        System.out.println(cb.rewind());
        // 交换
        symmetricScramble(cb);
        System.out.println(cb.rewind());
        // 交换
        symmetricScramble(cb);
        System.out.println(cb.rewind());
    }
}
