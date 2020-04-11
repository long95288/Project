import win32api
import win32gui
import win32con
import os
from PIL import Image


def setWallPaper(bmpFile):
    key = win32api.RegOpenKeyEx(win32con.HKEY_CURRENT_USER,
                                "Control Panel\\Desktop", 0, win32con.KEY_SET_VALUE)
    win32api.RegSetValueEx(key, "WallpaperStyle", 0, win32con.REG_SZ, "2")
    # 2拉伸桌面，0桌面居中
    win32api.RegSetValueEx(key, "TileWallPaper", 0, win32con.REG_SZ, "0")
    win32gui.SystemParametersInfo(win32con.SPI_SETDESKWALLPAPER, bmpFile, 1 + 2)

if __name__ == '__main__':
    print("正在设置桌面壁纸")

    # 文件路径
    img = Image.open("F:\\tools\\settingImage\\3fd38d86c2addb77fec9519ea40470a2b31711b2.jpg")
    # img.show()
    # bmpFile = img.convert("P")
    # bmpFile.show()
    # img.show()
    path = str(os.getcwd())+"\\wallpaper.bmp"
    # print(os.getcwd())
    # print("\n")
    #
    # img.save(path)
    # bmp = Image.open("bg.bmp")
    # bmpFile = bmp.load()
    path = "F:\\tools\\settingImage\\3fd38d86c2addb77fec9519ea40470a2b31711b2.jpg"
    setWallPaper(path)
    pass