from GUI.DownloadThread import DownloadThread

def endcallback():
    print("end call back....")
    return
def processback(process):
    print("process:"+process)
    return

if __name__ == '__main__':
    list = ["ddd","afafafa"]
    t = DownloadThread("haha.txt",chapter_url_list=list)
    t.start()
