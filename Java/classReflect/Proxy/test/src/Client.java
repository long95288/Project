import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Proxy;

public class Client{

    public static void main(String[] args){
        Subject realSubject = new RealSubject();

        try{
            // 声明的代理类的处理者
            DynamicProxy handler2 = new DynamicProxy();
            handler2.setSubject(realSubject);

            // 获得代理类
            Subject subject = (Subject)Proxy.newProxyInstance(handler2.getClass().getClassLoader(),
                                                            realSubject.getClass().getInterfaces(),
                                                            handler2);
            System.out.println(subject.getClass().getName());
            subject.rent();
            subject.hello("world");
        }catch(Exception e){
            e.printStackTrace();
        }

    }
}
