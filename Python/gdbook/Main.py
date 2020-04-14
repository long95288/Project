
import pymysql
import random
import requests
import json
import re
import time

headers = [
    {
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36'
    },
    {
        'user-agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36'
    }
]
root = "http://gdslzstsg.superlib.libsou.com/"
download_root = "http://photoapps.yd.chaoxing.com"
MAX_DOWN_SIZE = 5
def getBooks(url):
    header = random.choice(headers)
    print("请求: {} \n".format(url))
    response = requests.get(url,header)
    if response.status_code == 200:
        response_data = response.content.decode("UTF-8")
        data = response_data.replace("jQuery172024434167592111566_1586496638266(","")
        data = data[0:len(data)-2]
        json_data = json.loads(data)
        # print(json_data["msg"]["list"])
        return json_data["msg"]["list"]

def saveBook(books):
    db = pymysql.connect("localhost","root","root2037","library")
    affect_count = -1
    cursor = db.cursor()
    sql = "INSERT INTO t_book (author,cover,isbn,name,num,path,pubdate,publisher,schoolName,sort,Summary,cataid) VALUES (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)"
    insert_book_list = []
    for book in books:
        author = book.get("author", "null")
        cover = book.get("cover","")
        isbn = book.get("isbn","")
        if cover != "":
            cover = "{}{}".format(download_root, cover)
        name = book.get("name","null")
        num = book.get("num","null")
        path = book.get("path","")
        if path != "":
            path = "{}{}".format(download_root,path)
        pubdate = book.get("pubdate","")
        publisher = book.get("publisher","null")
        schoolName = book.get("schoolName","null")
        sort = book.get("sort","0")
        summary = book.get("Summary","")
        cataid = book.get("cataid","12")
        insert_book = (author,cover,isbn,name,num,path,pubdate,publisher,schoolName,sort,summary,cataid)
        insert_book_list.append(insert_book)

    try:
        affect_count = cursor.executemany(sql,insert_book_list)
        db.commit()
        cursor.close()
    except EnvironmentError:
        print("执行sql:{}失败".format(sql))
        db.rollback()
    db.close()
    return affect_count

def saveBookCatalogue(books):
    db = pymysql.connect("localhost","root","root2037","library")
    affect_count = -1
    cursor = db.cursor()
    sql = "INSERT INTO t_book_catalogue (book_id,catalogue_id) VALUES (%s,%s)"
    insert_book_list = []

def testGetBook():
    url="http://unitvb.featurelib.libsou.com/book/list_jsonp?callback=jQuery172024434167592111566_1586496638266&schoolid=72&cpage=1&pageSize=150&_=1586496638384"
    books = getBooks(url)
    count = saveBook(books)
    print("插入{}数据".format(count))
    pass

def getAllBooks():
    root_url = "http://unitvb.featurelib.libsou.com/book/list_jsonp?callback=jQuery172024434167592111566_1586496638266&schoolid=72&cpage={}&pageSize=150&_=1586496638384"
    for i in range(1, 42):
        url = root_url.format(i)
        books = getBooks(url)
        file = "book_page_{}.json".format(i)
        with open(file, 'w',encoding="utf-8") as f:
            f.write(json.dumps(books, ensure_ascii=False))
            print("写入json完成")
            f.close()
        time.sleep(2)
    #
    print("获得全部图书信息")
    pass
def saveAllBook():
    for i in range(1,42):
        file = "book_page_{}.json".format(i)
        books = []
        with open(file,'r',encoding="utf-8") as f:
            json_data = f.read()
            f.close()
            books = json.loads(json_data)
        count = saveBook(books)
        print("添加:{}本书".format(count))



ariaUrl = "http://localhost:6800/jsonrpc"
"""
获得当前下载列表
"""
def getCurrentDownload():
    jsondata = {
        "jsonrpc": "2.0",
        "id": "QXJpYU5nXzE1NDgzODg5MzhfMC4xMTYyODI2OTExMzMxMzczOA=="
    }
    jsondata["method"] = "aria2.tellActive"
    response =requests.post(ariaUrl,json=jsondata)
    if response.status_code == 200:
        return response.json()
    else:
        return None

def downloadBook():

    download_url_list = [
        "http://photoapps.yd.chaoxing.com/MobileApp/GDSL/pdf/gddj/1315722.pdf"
    ]
    jsondata = {
        "jsonrpc":"2.0",
        "id":"QXJpYU5nXzE1NDgzODg5MzhfMC4xMTYyODI2OTExMzMxMzczOA=="
    }
    # 获得当前下载数
    current_download_size = len(getCurrentDownload().get("result"))
    while current_download_size >= MAX_DOWN_SIZE:
        print("下载任务已满...")
        time.sleep(3)
        current_download_size = len(getCurrentDownload().get("result"))

    # 查询下载完成情况
    # print(ret.json())
    # request["method"] = "aria2.addUri"
    # request["params"] = [[],{}]
    # request["params"][0] = download_url_list
    # respose = requests.post(ariaUrl,json=request)
    # print(respose.status_code)
    # print(respose.content)
    pass

if __name__ == '__main__':
    # downloadBook()
    # getAllBooks()
    saveAllBook()
    pass