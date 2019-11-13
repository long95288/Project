from GUI.MainWind import Ui_Form
from PyQt5.QtWidgets import *
from PyQt5.QtGui import *
from PyQt5.QtCore import *

import sys
import requests
import re
from bs4 import BeautifulSoup
import random
import time

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
"""
页面主类
"""
class MainWindow(QWidget,Ui_Form):
    def __init__(self,parent=None):
        super(MainWindow,self).__init__(parent)
        self.chapter_list_url = []
        self.setupUi(self)
        self.connect_btn()

    # 设置按钮连接
    def connect_btn(self):
        self.analyze_btn.clicked.connect(lambda : self.handle_analyze())
        self.download_btn.clicked.connect(self.handle_download)

    # 分析URL
    def handle_analyze(self):
        self.status_label.setText("分析URL")
        print("分析URL")
        url = self.url_text_line_edit.text()
        response = requests.get(url, header)
        if response.status_code == 200:
            response_data = BeautifulSoup(response.content.decode('utf-8'),"html.parser")
            novel_name = response_data.select("#info > h1")[0].string
            print("小说名:"+novel_name)
            self.novel_name_edit.setText(novel_name)
            self.chapter_list_url = response_data.select("dd >a")
            # print("============章节列表============")
            print(self.chapter_list_url)
            # # 第一个章节从9开始
            # print(self.chapter_list_url[9].get("href"))
            if len(self.chapter_list_url) <= 0:
                print("获得章节列表错误")
                self.status_label.setText("获得章节列表错误!!!")
            else:
                chapter_count = len(self.chapter_list_url) - 9
                self.chapter_count.setText(str(chapter_count))
                self.analyze_btn.setEnabled(False)
                self.download_btn.setEnabled(True)
        else:
            message = "请求:"+url+"失败，code:"+response.status_code
            print(message)
    # 下载
    def handle_download(self):
        self.download_btn.setEnabled(False)
        save_novel_name = self.novel_name_edit.text() + ".txt"
        print("下载小说名:"+self.novel_name_edit.text())
        save_novel_file = open(save_novel_name,'a', encoding="utf-8")
        chapter_len = len(self.chapter_list_url)

        for i in range(9, chapter_len):
            # 获得数据
            request_url = rootUrl + self.chapter_list_url[i].get("href")
            print(i+"request:"+request_url)
            self.status_label.setText("请求:"+request_url)
            title, content = self.get_chapter_content(request_url)
            self.status_label.setText("保存:"+title)
            save_novel_file.write(title+"\n"+content)
            self.status_label.setText("保存:"+title+"成功！")
            time.sleep(timeout+random.random())
        self.download_btn.setEnabled(True)

    def get_chapter_content(self, url):
        response = requests.get(url, header)
        if response.status_code == 200:
            soup = BeautifulSoup(response.content.decode("utf-8"),"html.parser")
            # 章节名称
            chapter_title = soup.select('#wrapper > div.content_read > div.box_con > div.bookname > h1')
            # 章节内容
            chapter_content_temp = re.findall('<div id="content">(.*?)</div>', response.content.decode('utf-8'),
                                  re.S)
            chapter_content = self.clear_content(chapter_content_temp)
            return chapter_title,chapter_content
        else:
            return "章节:"+url, "获取失败"
    # 去掉空格,换行标签
    def clear_content(self,str):
        temp1_data = re.sub(r'.*?&nbsp;', '', str, count=0)
        temp2_data = re.sub(r'<br .*?>', '', temp1_data, count=0)
        return temp2_data
if __name__ == '__main__':
    app = QApplication(sys.argv)
    win = MainWindow()
    win.show()

    sys.exit(app.exec_())
