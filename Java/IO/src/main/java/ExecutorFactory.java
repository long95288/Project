/**
 * 运行器工厂
 * @author longquanxiao
 * @date 2020/2/22
 */
public class ExecutorFactory {
    /**
     * 根据类型创建相应的运行器
     * @param type 文件类型
     * @return 该文件对应的运行器
     */
    public static MyExecutor getExecutor(ExecutorType type){
        MyExecutor executor = null;
        switch (type) {
            case C:
                System.out.println("C");
                break;
            case PY:
                System.out.println("Python");
                break;
            case CPP:
                System.out.println("CPP");
                break;
            case JAVA:
                System.out.println("java");
                break;

                default:
                    break;
        }
        return executor;
    }

    public static void main(String[] args) {
        ExecutorFactory.getExecutor(ExecutorType.C);
        ExecutorFactory.getExecutor(ExecutorType.PY);
    }
}
