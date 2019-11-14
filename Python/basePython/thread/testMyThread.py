"""
测试 自定义的线程
"""
from basePython.thread.MyThread import MyThread


# 线程完成调用的函数
def threadRunCallBack(value):
    print("callback value="+value)


if __name__ == '__main__':
    threadName = "testThread"
    list = ["12","34","56"]
    print(len(list))
    t = MyThread(threadName,list=list)
    t.setRunCallBack(threadRunCallBack)
    t.start()
    print("线程运行............")
