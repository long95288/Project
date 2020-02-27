/**
 * @author longquanxiao
 * @date 2020/2/22
 * 运行器
 */
public interface MyExecutor {
    /**
     * 编译
     * @return
     */
    boolean compile();

    /**
     * 运行
     * @param
     * @return 运行的信息
     */
    boolean run(String path);

}
