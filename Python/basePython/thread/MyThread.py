import threading

class MyThread(threading.Thread):
    def __init__(self,name,list):
        threading.Thread.__init__(self)
        self.name = name
        self.list = list
        self.runCallBack = None

    def setRunCallBack(self,runCallBack):
        self.runCallBack = runCallBack

    def run(self):
        print("run ......")
        for i in range(0,len(self.list)):
            print("i "+str(i)+"="+ self.list[i])

        print("运行完成...回调函数")
        self.runCallBack("回调值")
        print("回调完成.........")
