from PyQt5.QtWidgets import *
from noveldownload.MainWind import Ui_Form

import sys


"""
页面主类
"""
class MainWindow(QWidget,Ui_Form):
    def __init__(self):
        super(QWidget, self).__init__()
        self.setupUi(self)
        self.connect_btn()

    # 设置按钮连接
    def connect_btn(self):
        self.analyze_btn.clicked.connect(lambda: self.handle_analyze())
        self.download_btn.clicked.connect(self.handle_download)

    # 分析按钮
    def handle_analyze(self):
        pass

    # 下载按钮
    def handle_download(self):
        pass


if __name__ == '__main__':
    app = QApplication(sys.argv)
    window = MainWindow()
    window.show()

    sys.exit(app.exec_())

