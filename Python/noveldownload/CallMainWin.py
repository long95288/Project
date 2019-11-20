from PyQt5.QtWidgets import *
from PyQt5.QtGui import QPixmap,QPainter,QBitmap,QCursor
from PyQt5.QtCore import Qt

from noveldownload.MainWind import Ui_Form
from noveldownload.NovelUtil import getNovelInfo
from noveldownload.NovelUtil import getHistoryUrl

from noveldownload.DownloadThread import DownloadThread

import sys

bg = "./image/bg2.jpg"
mask = "./image/mask.png"
"""
页面主类
"""
class MainWindow(QWidget,Ui_Form):
    def __init__(self):
        super(QWidget, self).__init__()
        # 小说相关的数据
        self.novelName = None
        # 章节数
        self.novelChapterCount = 0
        # 已经下载的章节数
        self.novelChapterDownloadCount = 0
        # 章节数据
        self.novelChapterUrlList = None
        # 下载进程
        self.downloadThread = None

        self.setupUi(self)
        self.bgImage = QPixmap(bg)
        self.maskImage = QBitmap(mask)
        self.setContainStyle()
        self.connect_btn()

    # 设置容器的样式
    def setContainStyle(self):
        self.setMask(self.maskImage)
        self.resize(self.maskImage.size())
        self.progressBar.setVisible(False)
        self.label_2.setVisible(False)
        style = """
        QPushButton{
            border-radius:5px;
            background-color: rgb(255, 255, 255,0.5);
            color:black;
        }
        """
        self.download_btn.setStyleSheet(style)
        self.download_btn.setWindowOpacity(0.5)
        edit_style = """
        QLineEdit{
        background-color: rgb(255, 255, 255,0.5);
        border-radius:5px;
        }
        """
        self.url_text_line_edit.setStyleSheet(edit_style)
        self.novel_name_edit.setStyleSheet(edit_style)
        analyze_style = """
        QPushButton{
            border-radius:5px;
            background-color: rgb(37, 175, 243,0.5);
            color:white;
        }
        """
        self.analyze_btn.setStyleSheet(analyze_style)
        self.continue_download_btn.setStyleSheet(style)
        self.stop_download_btn.setStyleSheet(style)
        # 设置退出样式
        exit_btn_style ="""
        QPushButton{
            border-image: url(./image/close.png);
        }
        """
        self.exit_btn.setText("")
        self.exit_btn.setStyleSheet(exit_btn_style)
        process_bar_style = """
        QProgressBar{
            border-radius:5px;
            background-color: rgb(255, 255, 255,0.5);
        }
        """
        self.progressBar.setStyleSheet(process_bar_style)

    def paintEvent(self, event):
        painter = QPainter(self)
        painter.drawPixmap(self.rect(), self.bgImage)
        # painter.drawPixmap(0,0,self.height(),self.width(),self.bgImage)

    # 设置按钮连接
    def connect_btn(self):
        # debug
        self.download_btn.setEnabled(True)
        # develop
        self.analyze_btn.clicked.connect(lambda: self.handle_analyze())
        self.download_btn.clicked.connect(self.handle_download)
        self.continue_download_btn.clicked.connect(self.handle_continue_download)
        self.exit_btn.clicked.connect(self.handle_exit)
        self.stop_download_btn.clicked.connect(self.handle_stop_download)

    # 分析按钮
    def handle_analyze(self):
        url = self.url_text_line_edit.text()
        novelName, novelChapterCount, novelChapterUrlList = getNovelInfo(url)
        self.novelChapterCount = novelChapterCount
        self.novel_name_edit.setText(str(novelName))
        self.novelName = self.novel_name_edit.text() + ".txt"
        # 如果设置的值是数字的话就直接内存报错
        self.chapter_count.setText(str(novelChapterCount))
        self.novelChapterUrlList = novelChapterUrlList
        # print(self.novelChapterUrlList)
        # 下载进程
        self.novelChapterDownloadCount = novelChapterCount - len(novelChapterUrlList)

        progressValue = self.novelChapterDownloadCount*100 / novelChapterCount
        self.progressBar.setProperty("value", progressValue)
        if progressValue > 0:
            self.status_label.setText("继续下载:"+str(progressValue))
            self.progressBar.setVisible(True)
            self.label_2.setVisible(True)

        self.download_btn.setEnabled(True)

    # 下载按钮
    def handle_download(self):
        # 显示进度条
        self.progressBar.setVisible(True)
        # 显示进度条的标签
        self.label_2.setVisible(True)
        # 关闭分析按钮
        self.analyze_btn.setEnabled(False)
        # 关闭下载按钮
        self.download_btn.setEnabled(False)
        # 启动下载进程
        self.downloadThread = DownloadThread(self.novelName,
                                             self.novelChapterUrlList,
                                             self.handle_process,
                                             self.handle_download_end)
        # self.downloadThread.setDownloadEndCallBack(self.handle_download_end)
        self.downloadThread.start()

    # 继续下载
    def handle_continue_download(self):
        # 继续下载
        self.continue_download_btn.setEnabled(False)
        # 1.获得最后的URL
        url = getHistoryUrl()
        # 2.设置URL
        self.url_text_line_edit.setText(url)
        # 3.调用分析
        self.handle_analyze()

    # 暂停下载
    def handle_stop_download(self):
        if self.downloadThread is not None:
            self.downloadThread.exitDownloadThread()

    def handle_exit(self):
        QApplication.instance().quit()

    def handle_status(self, value):
        self.status_label.setText(str(value))

    # 下载进度回调函数
    def handle_process(self, value):
        # 设置下载进度
        # (下载的数量)/总量
        count = self.novelChapterCount
        self.novelChapterDownloadCount += value
        showValue = self.novelChapterDownloadCount * 100 / count
        self.progressBar.setProperty("value", showValue)
        self.status_label.setText("下载中..."+str(showValue))

    def handle_download_end(self, message):
        self.analyze_btn.setEnabled(True)
        self.download_btn.setEnabled(False)
        self.continue_download_btn.setEnabled(True)
        self.novelChapterUrlList = None
        self.novelName = ""
        self.progressBar.setProperty("value", 0)
        print("dddd")
        print(message)
        # self.status_label.setText(str(message))
        return

    def mousePressEvent(self, event):
        if event.button() == Qt.LeftButton:
            self.m_drag = True
            self.m_DragPosition = event.globalPos() - self.pos()
            event.accept()
            self.setCursor(QCursor(Qt.OpenHandCursor))

    def mouseMoveEvent(self, QMouseEvent):
        if Qt.LeftButton and self.m_drag:
            self.move(QMouseEvent.globalPos() - self.m_DragPosition)
            QMouseEvent.accept()

    def mouseReleaseEvent(self, QMouseEvent):
        self.m_drag = False
        self.setCursor(QCursor(Qt.ArrowCursor))


if __name__ == '__main__':
    app = QApplication(sys.argv)
    window = MainWindow()
    window.show()
    sys.exit(app.exec_())

