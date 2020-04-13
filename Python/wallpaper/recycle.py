import win32api
import win32gui
import win32con
import os
import datetime
import random
import time


day = datetime.datetime.now().strftime('%Y%m%d')
log_filename = day+"_log.txt"
def log(value,print_flag = True):
    log_path = os.getcwd() + "\\log"
    if not os.path.exists(log_path) :
        os.mkdir(log_path)
    log_file_path = log_path + "\\" + log_filename
    logfile = open(log_file_path, 'a', encoding='utf-8')
    if logfile.writable():
        now_data = datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        log_message = "时间:{} : log : {}\n".format(now_data, value)
        if print_flag:
            print(log_message)
        logfile.write(log_message)
    try:
        logfile.close()
    except IOError:
        print("写入日志错误{}".format(IOError))
    else:
        return


def setWallPaper(wallpaper):
    try:
        key = win32api.RegOpenKeyEx(win32con.HKEY_CURRENT_USER,
                                    "Control Panel\\Desktop", 0, win32con.KEY_SET_VALUE)
        win32api.RegSetValueEx(key, "WallpaperStyle", 0, win32con.REG_SZ, "2")
        win32api.RegSetValueEx(key, "TileWallPaper", 0, win32con.REG_SZ, "0")
        win32gui.SystemParametersInfo(win32con.SPI_SETDESKWALLPAPER, wallpaper, 1 + 2)
        return True
    except EnvironmentError:
        log("except:{}错误".format(EnvironmentError))
        return False


"""
过滤图片文件
"""
def image_filter(file_list):
    image_list = []
    # 图片后缀
    image_suffix = ["jpg","png","bmp"]
    for file in file_list:
        if file.split(".")[-1] in image_suffix:
            image_list.append(file)
    return image_list


if __name__ == '__main__':
    image_dir = os.getcwd() + "\images"
    if os.path.exists(image_dir):
        file_list = os.listdir(image_dir)
        print(file_list)
        image_list = image_filter(file_list)
        print(image_list)
        if len(image_list) > 0:
            while True:
                wallpaper_path = image_dir + "\\" + random.choice(image_list)
                if os.path.exists(wallpaper_path):
                    if setWallPaper(wallpaper_path):
                        log("设置壁纸成功:path = {}".format(wallpaper_path))
                    else:
                        log("设置壁纸失败,详情查看日志。path = {}".format(wallpaper_path))
                        break
                else:
                    log("无效路径:{}".format(wallpaper_path))
                    break
                time.sleep(10)

        else:
            log("文件夹中没有图片文件:path = {}".format(image_dir))
    else:
        log("不存在路径:{}".format(image_dir))
