from GUI.DownloadThread import *



def callBack1():
    print("call Back")

if __name__ == '__main__':
    list = ["1","2","3","4"]
    name = "gadgaga"
    t = DownloadThread(name,list)
    t.setProcessCallBack(callBack1())
    t.start()