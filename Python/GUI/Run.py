import os
import sys
import threading

from PyQt5.QtCore import *
from PyQt5.QtGui import *
from PyQt5.QtWidgets import *

# 线程类
class AppThread(threading.Thread):
    def __init__(self,command):
        threading.Thread.__init__(self)
        self.command = command
        self.resolve = None

    def setResolve(self, resolve):
        self.resolve = resolve

    def run(self):
        print("运行程序")
        os.system(self.command)
        print("运行结束::::"+self.getName())
        print("运行的id::")
        self.threadEnd(self)


    def threadEnd(self,thread):
        self.resolve(thread)


class Form(QDialog):
    def __init__(self,parent = None):
        super(Form,self).__init__(parent)
        # 界面变量
        self.threads = []
        # 设置布局
        layout = QVBoxLayout()
        self.selectFilePath = None

        self.selectFileBtn = QPushButton("选择文件")
        self.appBtn1 = QPushButton("记事本")
        self.appBtn2 = QPushButton("VS Code")
        self.fileLabel = QLabel("选择的路径:")
        self.execute_btn = QPushButton("运行选择的程序")
        self.countBtn = QPushButton("当前运行的线程")

        self.selectFileBtn.clicked.connect(lambda :self.getfile())
        self.appBtn1.clicked.connect(lambda :self.executeApp(self.appBtn1))
        self.appBtn2.clicked.connect(lambda :self.executeApp(self.appBtn2))
        self.execute_btn.clicked.connect(self.executeByPath)
        self.countBtn.clicked.connect(lambda :self.countThreads())
        # 添加按钮
        layout.addWidget(self.selectFileBtn)
        layout.addWidget(self.appBtn1)
        layout.addWidget(self.appBtn2)
        layout.addWidget(self.fileLabel)
        layout.addWidget(self.execute_btn)
        layout.addWidget(self.countBtn)
        # 设置布局
        self.setLayout(layout)
        #
        self.setWindowTitle("启动器")
    def resolveThreadEnd(self,thread):
        print("线程:"+thread.getName()+"结束回调函数")
        self.threads.remove(thread)
        print("移除线程完成")

    def executeByPath(self):
        t = AppThread("\""+str(self.selectFilePath)+"\"")
        t.setResolve(self.resolveThreadEnd)
        t.start()
        self.threads.append(t)
    def executeApp(self, btn):
        if btn.text() == "记事本":
            print("打开记事本")
            # os.system("notepad")
            t = AppThread("notepad")
            t.setResolve(self.resolveThreadEnd)
            self.threads.append(t)
            t.start()

        elif btn.text() == "VS Code":
            print("打开VsCode")
            # os.system(r'"F:\Program Files\Microsoft VS Code\Code.exe"')
            t = AppThread(r'"F:\Program Files\Microsoft VS Code\Code.exe"')
            t.setResolve(self.resolveThreadEnd)
            self.threads.append(t)
            t.start()
            print(t)
            print(t.getName())
            # t2 = AppThread(r'"steam://rungameid/365670"')
            # t2.start()

    def getfile(self):
        path, _ = QFileDialog.getOpenFileName(self, '选择文件', '', "execute(*.exe)")
        print(path)
        self.fileLabel.setText(self.fileLabel.text() +path)
        self.selectFilePath = path
        # dlg = QFileDialog()
        # dlg.setFileMode(QFileDialog.AnyFile)
        # dlg.setFilter(QDir.Executable)
        #
        # if dlg.exec_():
        #     self.selectFilePath = dlg.selectFile()
        #     print("选择文件"+self.selectFilePath)

    def countThreads(self):
        print("当前运行的线程个数:")
        print(len(self.threads))
        print("每个线程的状况")
        for thread in self.threads:
            print("线程名:"+thread.getName())


if __name__ == '__main__':
    print("Hello World")
    app = QApplication(sys.argv)
    form = Form()
    form.show()
    sys.exit(app.exec_())

