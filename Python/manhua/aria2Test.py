import os

if __name__ == '__main__':
    url = "https://m.lfkmhw.com//shenshimonhua/11050_11.html"
    dir = "D:\\tmp"
    out = "tmp.html"
    cmd = "aria2c.exe \"{}\" --dir={} --out={} --stop=1 --allow-overwrite=true".format(url,dir,out)
    re = os.system(cmd)
    if re == 0:
        print("download success")
    else:
        print("download false")
