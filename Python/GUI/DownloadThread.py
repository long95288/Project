
import threading
import requests
import re
from bs4 import BeautifulSoup
import random
import time
"""
下载线程,根据获得的列表下载数据
"""
"""
全局变量
"""
header = {
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) '
                  'Chrome/71.0.3578.98 Safari/537.36'
}
# 笔趣阁的根网址
rootUrl = "http://www.022003.com"
# 请求间隔 second
timeout = 1


# 去掉空格,换行标签
def clear_content(str):
    temp1_data = re.sub(r'.*?&nbsp;', '', str, count=0)
    temp2_data = re.sub(r'<br .*?>', '', temp1_data, count=0)
    return temp2_data

# 获得章节的内容
def get_chapter_content(url):
    response = requests.get(url, header)
    if response.status_code == 200:
        soup = BeautifulSoup(response.content.decode("utf-8"), "html.parser")
        # 章节名称
        chapter_title = soup.select('#wrapper > div.content_read > div.box_con > div.bookname > h1')
        # 章节内容
        chapter_content_temp = re.findall('<div id="content">(.*?)</div>', response.content.decode('utf-8'),
                                          re.S)
        chapter_content = clear_content(chapter_content_temp)
        return chapter_title, chapter_content
    else:
        return "章节:" + url, "获取失败"


class DownloadThread(threading.Thread):
    def __int__(self, novel_name,chapter_url_list):
        threading.Thread.__init__(self)
        self.process_callback = None
        self.novel_name = novel_name
        self.chapter_url_list = chapter_url_list
        # 线程完成回调函数
        self.end_callback =None
        # 下载进度回调函数


    def setProcessCallBack(self,callback1):
        self.process_callback = callback1
    def setEndCallBack(self,callback1):
        self.end_callback = callback1

    def run(self):
       print("start download ......")
       chapter_count = len(self.chapter_url_list)
       save_novel_file = open(str(self.novel_name), 'a', encoding="utf-8")
       for i in range(9, chapter_count):
           request_url = rootUrl + self.chapter_url_list[i].get("href")
           print("request:"+request_url)


           # 返回下载进度
           proccess = i / chapter_count
           self.process_callback(proccess)
       self.end_callback()
