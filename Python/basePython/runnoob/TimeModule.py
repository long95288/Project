
import time

if __name__ == '__main__':
    ticks = time.time()
    print("时间戳为:{}".format(ticks))
    local_time = time.localtime(time.time())
    print("本地时间为:{}".format(local_time))
    # 获得可读的时间
    readable_time = time.asctime(time.localtime(time.time()))
    print("本地时间:{}".format(readable_time))
    # 格式化时间
    print(time.strftime("%Y - %m - %d %H:%M:%S",time.localtime()))
    print(time.strftime("%a %b %d %H:%M:%S %Y",time.localtime()))
    # 字符串转时间戳
    str = "Sat Mar 28 22:24:23 2016"
    print(time.mktime(time.strptime(str,"%a %b %d %H:%M:%S %Y")))
    #
