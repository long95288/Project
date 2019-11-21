/**
 * 可以使用
 * System.out.format()
 * or System.out.printf()
 * 格式化输出
 * @author longquanxiao
 * @date 2019/11/21
 */
public class SimpleFormat {
    public static void main(String[] args) {
        int x = 5;
        double y = 5.332543;
        // 传统方式
        System.out.println("Row 1 : ["+x+" "+y+"]");
        // 格式化方式
        System.out.format("Row 1:[%d %f]\n",x,y);
        System.out.printf("Row 1:[%d %f]\n",x,y);
    }
}
