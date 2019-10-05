import java.lang.reflect.Field;

public class Main{
    public static void main(String[] args) {
        Person p = new Person("xiaoming");
         try{
        Class pClass = p.getClass();
        Field f = pClass.getDeclaredField("name");
        f.setAccessible(true);
        Object value = f.get(p);
        System.out.println(value);
        f.set(p,"张三");
        System.out.println(p.getName());
        System.out.println("Hello world");
        Class stdClass = Student.class;
       
        // 获得公共字段
        System.out.println(stdClass.getField("score"));
       
        System.out.println(stdClass.getField("name")); 
        
        System.out.println(stdClass.getDeclaredField("grade"));
        }catch(NoSuchFieldException e) {
            e.printStackTrace();
        }catch(IllegalAccessException e){
            e.printStackTrace();
        }
    }
}
class Student extends Person {
    public int score;
    private int grade;
    public int name;
}
class Person {
    private String name;
    public Person(){}
    public Person(String name) {
        this.name = name;
    }
    String getName(){
        return name;
    }
}