from PyQt5.QtWidgets import *
from PyQt5.QtGui import QPixmap,QPainter,QBitmap,QCursor
from PyQt5.QtCore import Qt

from noveldownload.MainWind import Ui_Form
from noveldownload.NovelUtil import getNovelInfo
from noveldownload.DownloadThread import DownloadThread
from noveldownload.ContinueDownloadThread import ContinueDownloadThread
import sys

bg = "./image/bg2.jpg"
mask = "./image/mask.png"
"""
页面主类
"""
class MainWindow(QWidget,Ui_Form):
    def __init__(self):
        super(QWidget, self).__init__()
        self.novelChapterUrlList = None
        self.novelName = None
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
        painter.drawPixmap(self.rect(),self.bgImage)
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
        self.progressBar.setVisible(True)
        self.label_2.setVisible(True)
        self.analyze_btn.setEnabled(False)
        t = DownloadThread(self.novelName, self.novelChapterUrlList,self.handle_process)
        t.setDownloadEndCallBack(self.handle_download_end())
        t.start()
        self.download_btn.setEnabled(False)

    # 继续下载
    def handle_continue_download(self):
        url = self.url_text_line_edit.text()
        continueDownloadThread = ContinueDownloadThread(url)
        continueDownloadThread.setStatusCallBack(self.handle_status)
        continueDownloadThread.setDownloadEndCallBack(self.handle_download_end())
        continueDownloadThread.start()
        self.continue_download_btn.setEnabled(False)

    def handle_exit(self):
        QApplication.instance().quit()

    def handle_status(self, value):
        self.status_label.setText(str(value))

    # 下载进度回调函数
    def handle_process(self, value):
        self.progressBar.setProperty("value", value)
        self.status_label.setText("下载进度:"+str(value))

    def handle_download_end(self):
        self.analyze_btn.setEnabled(True)
        self.download_btn.setEnabled(False)
        self.continue_download_btn.setEnabled(True)
        self.novelChapterUrlList = None
        self.novelName = ""
        self.progressBar.setProperty("value", 0)
        self.status_label.setText("下载完成")


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
    # window.resize(379, 250)
    # window.setObjectName("MainWindow")
    # window.setStyleSheet("#MainWindow{border-image:url(./image/bg.jpg);}")
    window.show()
    sys.exit(app.exec_())

