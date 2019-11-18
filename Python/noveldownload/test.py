from noveldownload.NovelUtil import getNovelInfo
from noveldownload.NovelUtil import getChapterInfo
from noveldownload.DownloadThread import DownloadThread

import nose

# 测试获得小说信息
def testGetNovelInfo():
    url = "http://www.022003.com/9_9072/"
    name,count,list = getNovelInfo(url)
    print("小说名:"+name)
    print("章节数:"+str(count))
    print("章节列表:")
    print(list)

def testGetChapterInfo():
    url = "http://www.022003.com/9_9072/4802415.html"
    title,content, nextUrl = getChapterInfo(url)
    print("章节标题:"+title)
    print("下一个章节:"+ nextUrl)
    print("章节内容:\n")
    print(content)
def testLog():
    logMessage = "测试写入日志文件:"
    from noveldownload.NovelUtil import log as log
    log(logMessage)

def testSaveNovelFile():
    filename = "测试保存小说文件.txt"
    saveContent = "保存的内容"
    from noveldownload.NovelUtil import saveNovelFile
    saveNovelFile(filename,saveContent)

def threadCallBack(value):
    print("callback value "+ str(value))

def testDownloadThread():
    from noveldownload.DownloadThread import DownloadThread
    t = DownloadThread(novelName="fefe",novelChapterUrlList="ff",processCallBack=threadCallBack)
    t.start()

def testContinueDownloadThread():
    from noveldownload.ContinueDownloadThread import ContinueDownloadThread
    url = "http://www.022003.com/9_9072/19607848.html"
    t = ContinueDownloadThread(continueChapterUrl=url)
    t.start()
def testGetNovelNameByChapterUrl():
    from noveldownload.NovelUtil import getNovelNameByChapterUrl
    # url = "http://www.022003.com/9_9072/19607848.html"
    url = "http://www.022003.com/8_8293/4057865.html"
    novelName = getNovelNameByChapterUrl(url)
    print(novelName)

