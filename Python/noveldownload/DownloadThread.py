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
    def __init__(self,novelName=None,novelChapterUrlList=None,processCallBack=None):
        threading.Thread.__init__(self)
        self.novelName = novelName
        self.novelChapterUrlList = novelChapterUrlList
        self.processCallBack = processCallBack
        self.downloadEndCallBack = None

    def setDownloadEndCallBack(self,method):
        self.downloadEndCallBack = method

    def run(self):
        chapterCount = len(self.novelChapterUrlList)
        for i in range(0, chapterCount):
            print("request:"+str(i)+"="+self.novelChapterUrlList[i])
            # 请求数据和写入数据
            # code
            chapterUrl = self.novelChapterUrlList[i]
            log("download:"+chapterUrl+"\n")
            chapterTitle,chapterContent = getChapterInfo(chapterUrl)
            saveContent = chapterTitle+"\n"+chapterContent
            saveNovelFile(filename=self.novelName,content=saveContent)
            time.sleep(timeout+random.random())
            # 返回进度
            if self.processCallBack != None:
                processValue = (i/chapterCount)*100
                self.processCallBack(processValue)

        # 如果设置了进程结束回调函数则调用该函数
        if self.downloadEndCallBack != None:
            self.downloadEndCallBack()

