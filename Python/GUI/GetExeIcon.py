"""
获得应用软件的图标
"""
import sys
import os
import win32api
import win32con
import win32ui
import win32gui
# from PyQt5.QtCore import *
from PyQt5.QtWidgets import *
from PyQt5.QtGui import *


class MyWindow(QWidget):
    def __init__(self):
        super(MyWindow,self).__init__()
        self.setWindowTitle("获得应用的图标")
        self.resize(400, 300)
        layout = QVBoxLayout()
        # 选择文件按钮
        self.selectBtn = QPushButton()
        self.selectBtn.setText("选择运行的文件")
        self.selectBtn.clicked.connect(self.handle_select_file)
        layout.addWidget(self.selectBtn)

        # fileSelectDialog = QFileDialog()
        self.setLayout(layout)
    def handle_select_file(self):
        print("选择文件")
        filename,fileType = QFileDialog.getOpenFileName(self,"选择运行的文件","","exe(*.exe)")
        print("select file"+filename)
        large,small = win32gui.ExtractIconEx(filename, 0)
     #    image = self.iconToQImage(large[0])
  #       print(image)
        # win32gui.DestroyIcon(small[0])
        # hdc = win32ui.CreateDCFromHandle(win32gui.GetDC(0))
        # hbmp = win32ui.CreateBitmap()
        # hbmp.CreateCompatibleBitmap(hdc, 32, 32)
        # hdc = hdc.CreateCompatibleDC()
        # hdc.SelectObject(hbmp)
        # hdc.DrawIcon((0, 0), large[0])
        # hdc.DeleteDC()
        # handler = hbmp.GetHandle()
        # print(handler)
        # QPixmap().loadFromData(hbmp.GetHandle())
        # hbmp.SaveBitmapFile(hdc, "save.bmp")
        #hh = self.bitMapFromHICon(large[0])
        #print(hh)


    def bitMapFromHICon(self,hIcon):
        hdc = win32ui.CreateDCFromHandle(win32gui.GetDC(0))
        hbmp = win32ui.CreateBitmap()
        hbmp.CreateCompatibleBitmap(hdc, 32, 32)
        hdc = hdc.CreateCompatibleDC()
        hdc.SelectObject(hbmp)
        hdc.DrawIcon((0, 0), hIcon)
        hdc.DeleteDC()
        return hbmp.GetHandle()

    def iconToQImage(self,hIcon):
        hdc = win32ui.CreateDCFromHandle(win32gui.GetDC(0))
        hbmp = win32ui.CreateBitmap()
        hbmp.CreateCompatibleBitmap(hdc, hIcon.width, hIcon.height)
        hdc = hdc.CreateCompatibleDC()
        hdc.SelectObject(hbmp)
        win32gui.DrawIconEx(hdc.GetHandleOutput(), 0, 0, hIcon.hIcon, hIcon.width, hIcon.height, 0, None, 0x0003)
        bitmapbits = hbmp.GetBitmapBits(True)
        image = QImage(bitmapbits, hIcon.width, hIcon.height, QImage.Format_ARGB32_Premultiplied)
        return image


# win32api.MessageBox(win32con.NULL,"第一个窗口",'hello',win32con.MB_OK)
if __name__ == '__main__':
    app = QApplication(sys.argv)
    win = MyWindow()
    win.show()

    sys.exit(app.exec_())