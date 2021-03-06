"""
日志,写入日志文件
"""
import datetime
day = datetime.datetime.now().strftime("%Y%m%d")
log_file = day + "_log.txt"
def log(value,print_flag = True):
    logfile = open(log_file, 'a', encoding='utf-8')
    if logfile.writable():
        now_data = datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        log_message = "时间:{}:log:{}\n".format(now_data, value)
        if print_flag:
            print(log_message)
        logfile.write(log_message)
    try:
        logfile.close()
    except IOError:
        print("写入日志错误")
    else:
        return
