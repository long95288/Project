"""
继续下载小说线程
"""
import threading
from noveldownload.NovelUtil import *
import time
import random

timeout = 1

class ContinueDownloadThread(threading.Thread):
    def __init__(self, continueChapterUrl):
        threading.Thread.__init__(self)
        self.setDaemon(True)
        # 继续下载的URL
        self.continueChapterUrl = continueChapterUrl
        # 结束的标志
        self.endDownloadFlag = "/"+str(continueChapterUrl).split("/")[-2] + "/"
        # 状态回调
        self.statusCallBack = None
        # 结束回调
        self.downloadEndCallBack = None
        # 下载进程结束标志
        self.exit_flag = False

    # 状态回调
    def setStatusCallBack(self,method):
        self.statusCallBack = method

    def setDownloadEndCallBack(self,method):
        self.downloadEndCallBack = method

    # 结束进程
    def exitThread(self):
        self.exit_flag = True

    def run(self):
        filename = getNovelNameByChapterUrl(self.continueChapterUrl)
        filename = filename + ".txt"
        while not self.exit_flag:
            chapterTitle, chapterContent, nextChapterUrl = getChapterInfo(self.continueChapterUrl)
            if nextChapterUrl == self.endDownloadFlag or nextChapterUrl is None or nextChapterUrl == "":
                break
            else:
                # 写入文件
                saveContent = chapterTitle+"\n"+chapterContent
                log("保存:"+chapterTitle+":"+self.continueChapterUrl+"\n")
                # 设置URL
                saveNovelFile(filename=filename, content=saveContent)
                if self.statusCallBack is not None:
                    self.statusCallBack("保存:"+chapterTitle+"成功")
                self.continueChapterUrl = nextChapterUrl
                time.sleep(timeout+random.random())

        if self.downloadEndCallBack is not None:
            self.downloadEndCallBack()
            self.exitThread()
