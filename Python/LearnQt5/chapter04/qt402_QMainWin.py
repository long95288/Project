"""
主窗口
"""
import sys
from PyQt5.QtWidgets import QMainWindow,QApplication,QDesktopWidget,QPushButton,QVBoxLayout,QWidget
from PyQt5.QtGui import *

class MainWindow(QMainWindow):
    def __init__(self,parent=None):
        super(MainWindow,self).__init__(parent)
        # 设置主窗口的属性
        self.resize(400,200)
        self.status = self.statusBar()
        self.status.showMessage("这是状态栏的提示",5000)
        self.setWindowTitle("主窗口")
        # 应用定位到中间
        self.center()
        # 关闭主窗口
        self.button1 = QPushButton("关闭窗口")
        self.button1.clicked.connect(self.handleBtn1Clicked)

        # 设置布局
        layout = QVBoxLayout()
        layout.addWidget(self.button1)

        # window 中主要的应用框架
        main_frame = QWidget()
        main_frame.setLayout(layout)
        self.setCentralWidget(main_frame)

    """
    处理按钮点击事件
    """
    def handleBtn1Clicked(self):
        # 发送的对象
        sender = self.sender()
        print(sender)
        print(sender.text() + "被按下")
        # 应用实例
        qApp = QApplication.instance()
        # 退出应用
        qApp.quit()
    """
    屏幕居中显示
    """
    def center(self):
        # 屏幕的图标
        screen = QDesktopWidget().screenGeometry()

        # 应用的大小坐标
        size = self.geometry()
        self.move((screen.width() - size.width())/2, (screen.height() - size.height())/2)


if __name__ == '__main__':
    # 应用框架
    app = QApplication(sys.argv)
    app.setWindowIcon(QIcon("./image/record.png"))
    form = MainWindow()
    form.show()
    sys.exit(app.exec_())
