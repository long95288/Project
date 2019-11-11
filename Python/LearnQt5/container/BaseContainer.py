"""
基础控件合集
"""
import sys
from PyQt5.QtWidgets import *
from PyQt5.QtGui import *
from PyQt5.QtCore import *


"""
页面主类
"""
class MainWindow(QMainWindow):
    def __init__(self,parent = None):
        super(MainWindow, self).__init__(parent)
        self.resize(400,200)
        self.status = self.statusBar()
        self.setWindowTitle("MainWindow")
        self.status.showMessage("这是状态栏消息", 5000)
        # 提示气泡
        toolTipFont = QFont('SansSerif',10)
        QToolTip.setFont(toolTipFont)
        self.setToolTip("提示气泡<h2>标题2内容</h2>")
        # 标签QLabel
        label1 = QLabel(self)
        label1.setText("标签内容")
        label1.setAutoFillBackground(True)

        label2 = QLabel(self)
        label2.setText("第二个标签的内容")
        label2.setAutoFillBackground(True)
        # 设置标签组件在页面内容的坐标: x,y,weight,height
        label2_position = QRect(40,60,60,60)
        label2.setGeometry(label2_position)


"""
运行类
"""
if __name__ == '__main__':
    app = QApplication(sys.argv)
    # 应用图标
    icon = QIcon("./resource/account.png")
    app.setWindowIcon(icon)
    mainWindow = MainWindow()
    mainWindow.show()

    # mainWindow2 = MainWindow()
    # mainWindow2.show()
    # 退出应用
    sys.exit(app.exec_())
