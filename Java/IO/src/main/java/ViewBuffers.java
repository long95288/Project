import java.nio.*;
import java.sql.SQLOutput;

/**
 * @author longquanxiao
 * @date 2020/1/17
 */
public class ViewBuffers {
    public static void main(String[] args) {
        ByteBuffer bb = ByteBuffer.wrap(new byte[]{0,0,0,0,0,0,0,'a'});
        bb.rewind();
        System.out.print("Byte Buffer ");
        while (bb.hasRemaining()){
            // 打印数据的位置和数据的值
            System.out.print(bb.position() + "-> " + bb.get() + ",");
        }
        System.out.println();
        // 字符视图缓冲区
        CharBuffer cb = ((ByteBuffer) bb.rewind()).asCharBuffer();
        System.out.print("Char Buffer ");
        // 打印字符
        while (cb.hasRemaining()){
            System.out.print(cb.position() + " -> " + cb.get()+ ",");
        }
        System.out.println();
        //
        FloatBuffer fb = ((ByteBuffer)bb.rewind()).asFloatBuffer();
        System.out.print("Float Buffer ");
        while (fb.hasRemaining()){
            System.out.print(fb.position() + "-> "+ fb.get()+",");
        }
        System.out.println();
        IntBuffer ib = ((ByteBuffer)bb.rewind()).asIntBuffer();
        System.out.print("Int Buffer");
        while (ib.hasRemaining()){
            System.out.print(ib.position() + "-> "+ib.get()+",");
        }
        System.out.println();
        LongBuffer lb = ((ByteBuffer)bb.rewind()).asLongBuffer();
        System.out.print("Long Buffer ");
        while (lb.hasRemaining()){
            System.out.print(lb.position()+ "-> "+ lb.get()+",");
        }
        System.out.println();
        ShortBuffer sb = ((ByteBuffer)bb.rewind()).asShortBuffer();
        System.out.print("Short Buffer ");
        while (sb.hasRemaining()){
            System.out.print(sb.position() + "-> "+ sb.get()+ ",");
        }
        System.out.println();
        DoubleBuffer db = ((ByteBuffer)bb.rewind()).asDoubleBuffer();
        System.out.print("Double Buffer ");
        while (db.hasRemaining()){
            System.out.print(db.position() + "-> "+db.get()+ ",");
        }
        System.out.println();
    }
}
