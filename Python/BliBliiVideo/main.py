# encoding=utf-8
# coding=utf-8

import json
import csv
import requests
import time
import random

header = {
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) '
                  'Chrome/71.0.3578.98 Safari/537.36'
        }

# 读取文件
midf = open('mid.txt','r',encoding='utf-8')
mid = midf.readline()

# 生成初始化的请求url 
page1 = "https://space.bilibili.com/ajax/member/getSubmitVideos?mid="+str(mid)+\
                "&pagesize=30&tid=0&page=1&keyword=&order=pubdate"

# 写入文件
csvf = open('data.csv','w+',encoding='utf-8')
csvwriter = csv.writer(csvf)
csvwriter.writerow(('URL','title','length'))


# 写入数据函数
def saveData(jsondata):
    for item in jsondata:
        url = "https://www.bilibili.com/video/av"+str(item['aid'])
        title = item['title']
        length = item['length']
        csvwriter.writerow((url,title,length))

# 获得数据函数
def getTextData(url):
    res = requests.get(url,headers=header)
    return res.text

def start():
    jsondata = getTextData(page1)
    s = json.loads(jsondata)
    pages =  s['data']['pages']
    for page in range(1,pages):
        geturl = "https://space.bilibili.com/ajax/member/getSubmitVideos?mid="+str(mid)+\
                "&pagesize=30&tid=0&page="+\
                str(page)+"&keyword=&order=pubdate"
        jsondata = getTextData(geturl)
        s1 = json.loads(jsondata)
        saveData(s1['data']['vlist'])
        time.sleep(2+random.random())

# 开始爬取
start()