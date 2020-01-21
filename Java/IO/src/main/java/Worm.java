import java.io.*;
import java.util.Random;

/**
 * @author longquanxiao
 * @date 2020/1/21
 */
class Data implements Serializable{
    private int n;
    public Data(int n){this.n = n;}
    @Override
    public String toString() {
        return Integer.toString(n);
    }
}
public class Worm implements Serializable {
    private static Random rand = new Random(47);
    private Data[] d = {
            new Data(rand.nextInt(10)),
            new Data(rand.nextInt(10)),
            new Data(rand.nextInt(10))
    };
    private Worm next;
    private char c;
    public Worm(int i,char x){
        System.out.println("Worm constructor: "+ i);
        this.c = x;
        if (--i> 0){
            // 自动构建下一个蠕虫
            next = new Worm(i,(char)(x+1));
        }
    }
    public Worm(){
        System.out.println("Defaut constructor");
    }

    @Override
    public String toString() {
        StringBuilder result = new StringBuilder(":");
        result.append(c);
        result.append("(");
        for (Data dat :
                d) {
            result.append(dat);
        }
        result.append(")");
        if (next!=null){
            result.append(next);
        }
        return result.toString();
    }

    public static void main(String[] args) throws ClassNotFoundException, IOException {
        Worm w = new Worm(6,'a');
        System.out.println("w = "+w);
        //
        ObjectOutputStream out = new ObjectOutputStream(new FileOutputStream("worm.out"));
        // 对象持久化
        out.writeObject("Worm write");
        out.writeObject(w);
        out.close();

        // 反序列化
        ObjectInputStream in = new ObjectInputStream(new FileInputStream("worm.out"));
        // 读取第一个对象
        String s = (String) in.readObject();
        Worm w2 = (Worm) in.readObject();
        System.out.println("w2 = "+w2);

        //
        ByteArrayOutputStream bout = new ByteArrayOutputStream();
        ObjectOutputStream out2 = new ObjectOutputStream(bout);
        out2.writeObject("Worm storage\n");
        out2.writeObject(w);
        out2.flush();
        //
        ObjectInputStream in2 = new ObjectInputStream(new ByteArrayInputStream(bout.toByteArray()));
        s = (String)in2.readObject();
        Worm worm3 =(Worm)in2.readObject();
        System.out.println(s+ " w3 = "+ worm3);

    }
}
