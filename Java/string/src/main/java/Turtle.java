import java.io.PrintStream;
import java.util.Formatter;

/**
 * @author longquanxiao
 * @date 2019/11/27
 */
public class Turtle {
    private String name;
    private Formatter f;

    public Turtle(String name, Formatter f) {
        this.name = name;
        this.f = f;
    }
    public void move(int x,int y){
        f.format("%s The Turtle is a (%d,%d)\n",name,x,y);
    }

    public static void main(String[] args) {
        PrintStream outAlias = System.out;
        Turtle tommy = new Turtle("Tommy",new Formatter(System.out));
        Turtle terry = new Turtle("Terry",new Formatter(outAlias));

        tommy.move(0,0);
        terry.move(1,1);
        tommy.move(2,2);
        terry.move(3,3);
        tommy.move(4,4);
    }
}
