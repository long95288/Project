"""
爬虫小说工具
"""
import sys
from PyQt5.QtWidgets import *
from PyQt5.QtGui import *
from PyQt5.QtCore import *

import requests
import re
import time
from bs4 import BeautifulSoup
import random

"""
全局变量设置
"""
# 请求头
header = {
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) '
                  'Chrome/71.0.3578.98 Safari/537.36'
}
#
class MainWindow(QWidget):
    def __init__(self, parent =None):
        super(MainWindow, self).__init__(parent)
        # 笔趣阁的URL地址
        self.rootURL = "http://www.022003.com"
        # 章节
        self.novel_chapters = []
        self.initGUI()
    """
    初始化页面函数
    """
    def initGUI(self):
        # 基本信息
        self.setWindowTitle("小说下载器")
        self.resize(300, 400)
        # 布局
        layout = QFormLayout()
        # 地址后缀
        self.url_text = QLineEdit()
        layout.addRow(QLabel("小说地址:"),self.url_text)
        # 操作按钮
        # 开始分析按钮
        self.start_analyze_btn = QPushButton()
        self.start_analyze_btn.setText("分析URL")
        self.start_analyze_btn.clicked.connect(self.startAnalyzeURL)
        layout.addRow(QLabel("分析URL:"), self.start_analyze_btn)

        # 开始爬取
        start_btn = QPushButton()
        start_btn.setText("开始爬取")
        start_btn.setEnabled(False)
        start_btn.clicked.connect(self.startGetNovel)
        layout.addRow(QLabel(""))
        # 提示信息
        self.novel_name = QLabel("未知")
        self.novel_total_chapter = QLabel("未知")
        self.novel_end_flag = QLabel("未知")
        self.novel_status = QLabel("未分析")
        # 设置布局
        layout.addWidget(start_btn)
        layout.addRow(QLabel("小说名"), self.novel_name)
        layout.addRow(QLabel("总章节:"), self.novel_total_chapter)
        layout.addRow(QLabel("结束标志"),self.novel_end_flag)
        layout.addRow(QLabel("当前状态:"),self.novel_status)
        self.setLayout(layout)
    # 分析url
    def startAnalyzeURL(self):
        url = self.url_text.text()
        response = requests.get(url, headers=header)
        if response.status_code == 200:
            # 分析页面数据
            soup = BeautifulSoup(response.content.decode('utf-8'), "html.parser")
            print(soup)
            self.novel_name.setText(soup.select('#info > h1')[0].string)
        else:
            message = "请求"+url+"失败:"+response.status_code
            print(message)
            self.novel_status.setText(message)
        pass
    def startGetNovel(self):
        print("开始爬取数据")
        print("爬取的URL:"+self.url_text.text())
        self.novel_total_chapter.setText(self.url_text.text())


if __name__ == '__main__':
    app = QApplication(sys.argv)
    win = MainWindow()
    win.show()
    sys.exit(app.exec_())

