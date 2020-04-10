import datetime
import pymysql
import time
import requests

day = datetime.datetime.now().strftime('%Y%m%d')
log_file = day+"_log.txt"
rpc_url = "http://localhost:6800/jsonrpc"

def log(value,print_flag = True):
    logfile = open(log_file, 'a', encoding='utf-8')
    if logfile.writable():
        now_data = datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        log_message = "时间:{} : log : {}\n".format(now_data, value)
        if print_flag:
            print(log_message)
        logfile.write(log_message)
    try:
        logfile.close()
    except IOError:
        print("写入日志错误")
    else:
        return

def queryUndownloadBook(size):
    db = pymysql.connect("localhost","root","root2037","book")
    cursor = db.cursor()
    sql = "SELECT DISTINCT id,path FROM t_book WHERE status = 0 LIMIT 0,{} ".format(size)
    results = []
    try:
        cursor.execute(sql)
        # 获得记录
        results = cursor.fetchall()
        cursor.close()
    except EnvironmentError:
        log("无法获取数据")
    # 关闭数据库
    db.close()

    # 返回结果
    return results

def download_book(path):

    postdata = {
        "jsonrpc": "2.0",
        "id": "QXJpYU5nXzE1NDgzODg5MzhfMC4xMTYyODI2OTExMzMxMzczOA=="
    }
    rpc_request = postdata
    rpc_request["method"] = "aria2.addUri"
    rpc_request["params"] = [[path]]
    response = requests.post(url=rpc_url,json=rpc_request)
    if response.status_code == 200:
        result = response.json().get("result",[])
        print("gid: {}".format(result))
        return result
        # if len(result) > 0:
        #     gid = result[0]
        #     print("gid: {}".format(gid))
        #     return gid
        # else:
        #     log("添加下载:{}失败".format(path))
    else:
        log("无法调用aria2")


def download_status(gid):
    postdata = {
        "jsonrpc": "2.0",
        "id": "QXJpYU5nXzE1NDgzODg5MzhfMC4xMTYyODI2OTExMzMxMzczOA=="
    }
    rpc_request = postdata
    rpc_request["method"] = "aria2.tellStatus"
    rpc_request["params"] = [gid]
    response = requests.post(url=rpc_url,json=rpc_request)
    if response.status_code == 200:
        result = response.json().get("result","")
        if result != "":
            status = result.get("status")
            if status != "":
                return status
    return None


def updateBookStatus(id):
    db = pymysql.connect("localhost", "root", "root2037", "book")
    cursor = db.cursor()
    sql = "UPDATE t_book SET status = 1 WHERE id = {}".format(id)
    affect_count = 0
    try:
        affect_count = cursor.execute(sql)
        db.commit()
        cursor.close()
    except EnvironmentError:
        log("执行sql: {} 失败".format(sql))
        db.rollback()
    db.close()
    return affect_count

if __name__ == '__main__':
    log("开始下载书籍")
    timeout = 2
    log("获得数据库数据")
    undownload_books = queryUndownloadBook(10)
    # print(undownload_books)
    error = False
    while len(undownload_books) > 0 and not error:
        for book in undownload_books:
            id = book[0]
            path = book[1]
            log("开始下载: {}".format(path))
            gid = download_book(path)
            status = download_status(gid)
            while True and not error:
                if status == "active":
                    time.sleep(3)
                    print("下载中.....\n")
                    status = download_status(gid)
                if status == "complete":
                    # 一本书下载完成,更新数据库
                    log("下载: {} 完成".format(path))
                    updateBookStatus(id)
                    log("更新数据库.....")
                    break
                elif status == "waiting":
                    log("下载队列已满")
                    time.sleep(4)
                elif status == "paused":
                    log("暂停下载")
                    break
                elif status == "error":
                    log("下载错误")
                    error = True
                    break
                elif status == "removed":
                    log("已经从下载队列中移除")
                    break

            if error:
                break
        log("获得数据库数据")
        undownload_books = queryUndownloadBook(10)