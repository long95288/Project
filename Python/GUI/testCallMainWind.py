from GUI.CallMainWind import DownloadThread



def callBack():
    print("call Back")

if __name__ == '__main__':
    t = DownloadThread(novel_name="",chapter_list_url="", end_callback=callBack)
    t.start()