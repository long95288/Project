import java.math.BigInteger;
import java.util.Formatter;
import java.util.HashMap;

/**
 * @author longquanxiao
 * @date 2019/12/18
 * 各种类型使用formatter 进行格式化输出
 */
public class Conversion {
    public static void main(String[] args) {
        /*
         * Formatter
         */
        Formatter f = new Formatter(System.out);
//        String[] types =  {
//                "d: %d\n", // 0
//                "c: %c\n", // 1
//                "b: %b\n", // 2
//                "s: %s\n", // 3
//                "f: %f\n", // 4
//                "e: %e\n", // 5
//                "x: %x\n", // 6
//                "h: %h\n"  // 7
//        };
        HashMap<String,String> types = new HashMap<>(8);
        types.put("d","d: %d\n");
        types.put("c","c: %c\n");
        types.put("b","b: %b\n");
        types.put("s","s: %s\n");
        types.put("f","f: %f\n");
        types.put("e","e: %e\n");
        types.put("x","x: %x\n");
        types.put("h","h: %h\n");
        /*
         * 字符类型的格式化输出
         * 字符类型可以使用的格式化输出有
         * 1、s String
         * 2、c Unicode字符
         * 3、b Boolean值
         * 4、h 散列码(16进制)
         */

        char u = 'a';
        System.out.println("u = 'a'");
        f.format("s: %s\n",u);
        f.format("c: %c\n",u);
        f.format("b: %b\n",u);
        f.format("h: %h\n",u);

        /*
        * 整数的格式化输出
        * 1、d 整数型 十进制
        * 2、c
        * 3、b
        * 4、s
        * 5、x
        * */
        int v = 121;
        System.out.println("v = 121");
        f.format("d: %d\n",v);
        f.format("c: %c\n",v);
        f.format("b: %b\n",v);
        f.format("s: %s\n",v);
        f.format("x: %x\n",v);
        f.format("h: %h\n",v);

        /*
        * 大数
        * d b s x h
        * */
        BigInteger w = new BigInteger("50000000000000000000");
        System.out.println("w = BigInteger(\"50000000000000000000\")");
        f.format("d: %d\n",w);
        f.format("b: %b\n",w);
        f.format("s: %s\n",w);
        f.format("x: %x\n",w);
        f.format("h: %h\n",w);

        /*
        * 浮点数类型
        * */
        double x = 179.543;
        System.out.println("x = 179.543");
        f.format(types.get("b"),x);
        f.format(types.get("f"),x);
        f.format(types.get("s"),x);
        f.format(types.get("e"),x);
        f.format(types.get("h"),x);

        /*
        * 对象类型
        * */
        Conversion y = new Conversion();
        System.out.println("y = new Conversion()");
        f.format(types.get("b"),y);
        f.format(types.get("s"),y);
        f.format(types.get("h"),y);

        /*
        * 布尔类型
        * */
        boolean z = false;
        System.out.println("z = false");
        f.format(types.get("b"),z);
        f.format(types.get("s"),z);
        f.format(types.get("h"),z);

        String noneStr = null;
        System.out.println("noneStr = null");
        f.format(types.get("b"),noneStr);


    }
}
