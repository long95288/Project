import threading
import time
from noveldownload.NovelUtil import saveNovelFile
from noveldownload.NovelUtil import getChapterInfo
from noveldownload.NovelUtil import log
import random
"""
下载线程
"""
timeout = 1
class DownloadThread(threading.Thread):
    def __init__(self,novelName=None,novelChapterUrlList=None,processCallBack=None,downloadEndCallBack=None):
        threading.Thread.__init__(self)
        self.setDaemon(True)
        self.novelName = novelName
        self.novelChapterUrlList = novelChapterUrlList
        self.processCallBack = processCallBack
        self.downloadEndCallBack = downloadEndCallBack
        self.exitFlag = False

    def setDownloadEndCallBack(self, method):
        self.downloadEndCallBack = method
    #
    # def setExitFlag(self):
    #     self.exitFlag = True
    def exitCallBack(self):
        if self.downloadEndCallBack is not None:
            print("回调。。。")
            message = "下载完成!!"
            if self.exitFlag:
                message = "退出下载"
            self.downloadEndCallBack(message=message)
    # 结束线程
    def exitDownloadThread(self):
        print("停止下载... 在线程中")
        mutex = threading.Lock()
        mutex.acquire()
        self.exitFlag = True
        print("设置完成...")
        mutex.release()
        # 异步调用
        # t = threading.Thread(target=self.setExitFlag(),args=())
        # t.setDaemon(True)
        # t.start()

    def run(self):
        chapterCount = len(self.novelChapterUrlList)
        for i in range(0, chapterCount):
            mutex = threading.Lock()
            mutex.acquire()
            if self.exitFlag:
                print("退出.....")
                break
            else:
                # 请求数据和写入数据
                # code
                chapterUrl = self.novelChapterUrlList[i]
                chapterTitle,chapterContent,_ = getChapterInfo(chapterUrl)
                saveContent = chapterTitle+"\n"+chapterContent
                saveNovelFile(filename=self.novelName,content=saveContent)
                log("保存:" + chapterTitle + "#" + chapterUrl + "\n")

                # 返回进度
                if self.processCallBack != None:
                    # 返回下载的章节的index
                    self.processCallBack(i)
            # 放锁
            mutex.release()
            # 等待
            time.sleep(timeout + random.random())
        # print("退出循环")
        # 如果设置了进程结束回调函数则调用该函数
        # if self.downloadEndCallBack is not None:
        #     print("回调。。。")
        #     message = "下载完成!!"
        #     if self.exitFlag:
        #         message = "退出下载"
        #     self.downloadEndCallBack(message=message)
        # t = threading.Thread(target=self.exitCallBack(),args=())
        # t.setDaemon(True)
        # t.start()
        # print("线程退出.....")
