"""
日志,写入日志文件
"""
def log(value):
    logfile = open('log.txt', 'a', encoding='utf-8')
    if logfile.writable():
        logfile.write(value)
    try:
        logfile.close()
    except IOError:
        print("写入日志错误")
    else:
        return
