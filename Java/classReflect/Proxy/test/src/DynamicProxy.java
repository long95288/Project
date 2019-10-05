import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;

/**
 * 动态代理类
 * @author longquanxiao
 */
public class DynamicProxy implements InvocationHandler{

    /**
     * 维护的的对象
     */
    private Object subject;

    public DynamicProxy(){}
    /**
     * 构造函数，输入要代理的对象
     * @param subject
     */
    public DynamicProxy(Object subject){
        this.subject = subject;
    }
    /**
     * @param subject the subject to set
     */
    public void setSubject(Object subject) {
        this.subject = subject;
    }
    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable{
        //代理对象前的操作
        System.out.println("before rent house");
        System.out.println("Method:"+method);

        // 调用原有的函数
        method.invoke(subject, args);

        // 调用原有的函数之后
        System.out.println("after rent house");

        return null;
    }
}
