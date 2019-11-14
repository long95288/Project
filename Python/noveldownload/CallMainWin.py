from PyQt5.QtWidgets import *
from noveldownload.MainWind import Ui_Form
from noveldownload.NovelUtil import getNovelInfo
from noveldownload.DownloadThread import DownloadThread
import sys


"""
页面主类
"""
class MainWindow(QWidget,Ui_Form):
    def __init__(self):
        super(QWidget, self).__init__()
        self.novelChapterUrlList = None
        self.novelName = None
        self.setupUi(self)
        self.connect_btn()
    # 设置按钮连接
    def connect_btn(self):
        # debug
        self.download_btn.setEnabled(True)
        # develop
        self.analyze_btn.clicked.connect(lambda: self.handle_analyze())
        self.download_btn.clicked.connect(self.handle_download)

    # 分析按钮
    def handle_analyze(self):
        url = self.url_text_line_edit.text()
        novelName,novelChapterCount,novelChapterUrlList = getNovelInfo(url)
        self.novel_name_edit.setText(novelName)
        self.novelName = str(novelName) + ".txt"
        # 如果设置的值是数字的话就直接内存报错
        self.chapter_count.setText(str(novelChapterCount))
        self.novelChapterUrlList = novelChapterUrlList
        self.download_btn.setEnabled(True)

    # 下载按钮
    def handle_download(self):
        self.analyze_btn.setEnabled(False)
        t = DownloadThread(self.novelName, self.novelChapterUrlList,self.handle_process)
        t.setDownloadEndCallBack(self.handle_download_end())
        t.start()
        self.download_btn.setEnabled(False)

    # 下载进度回调函数
    def handle_process(self, value):
        self.progressBar.setProperty("value", value)
        self.status_label.setText("下载进度:"+str(value))

    def handle_download_end(self):
        self.analyze_btn.setEnabled(True)
        self.download_btn.setEnabled(False)
        self.novelChapterUrlList = None
        self.novelName = ""
        self.progressBar.setProperty("value", 0)
        self.status_label.setText("下载完成")


if __name__ == '__main__':
    app = QApplication(sys.argv)
    window = MainWindow()
    window.show()

    sys.exit(app.exec_())

