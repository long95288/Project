"""
小说下载工具类
"""
import requests
from bs4 import BeautifulSoup
import re
"""
全局变量
"""
header = {
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) '
                  'Chrome/71.0.3578.98 Safari/537.36'
}
# 笔趣阁的根网址
rootUrl = "http://www.022003.com"
# 日志文件
logfile = open('log.txt','a',encoding='utf-8')
"""
根据小说的URL获得小说的信息
params:
    url: 小说首页的url
return:
    novelName: 小说名
    novelChapterCount: 小说章节数
    novelChapterUrlList: 小说每个章节的url地址
可能会返回空值
"""
def getNovelInfo(url):
    # 小说名
    novelName = None
    # 小说章节数
    novelChapterCount = None
    # 小说章节地址数据
    novelChapterUrlList = []
    response = requests.get(url, header)
    if response.status_code == 200:
        response_data = BeautifulSoup(response.content.decode('utf-8'), "html.parser")
        novelName = response_data.select("#info > h1")[0].string
        listTemp = response_data.select("dd >a")
        # 获得每个章节的地址
        for i in range(9,len(listTemp)):
            novelChapterUrlList.append(rootUrl+listTemp[i].get("href"))

        novelChapterCount = len(novelChapterUrlList)
    return novelName,novelChapterCount,novelChapterUrlList

"""
去掉空格标签和换行标签
"""
def clearSpaceTabAndBrTab(str):
    # 去除空格标签
    noSpace_data = re.sub(r'.*?&nbsp;', '', str, count=0)
    # 去除换行标签
    clearData = re.sub(r'<br .*?>', '', noSpace_data, count=0)
    return clearData

"""
获得章节数据
params:
    url: 章节的url
return:
    chapterTitle: 章节标题
    chapterContent: 章节内容
"""
def getChapterInfo(url):
    chapterTitle = ""
    chapterContent = ""
    response = requests.get(url,header)
    if response.status_code == 200:
        soup = BeautifulSoup(response.content.decode("utf-8"),"html.parser")
        # 章节标题
        chapterTitle = soup.select('#wrapper > div.content_read > div.box_con > div.bookname > h1')[0].string
        # 章节内容
        tempContent = re.findall('<div id="content">(.*?)</div>', response.content.decode('utf-8'),
                                  re.S)[0]
        chapterContent = clearSpaceTabAndBrTab(tempContent)
    return chapterTitle, chapterContent

"""
将content里面的内容写入文件中
params:
    filename: 文件名
    content: 写入内容,添加到后面
"""
def saveNovelFile(filename,content):
    # 小说文件
    try:
        novel_file = open(filename,'a',encoding="utf-8")
        novel_file.write(content+"\n")
    except IOError:
        log("写入:"+filename+"失败")
    else:
        novel_file.close()

"""
日志,写入日志文件
"""
def log(value):
    if logfile.writable():
        logfile.write(value)
