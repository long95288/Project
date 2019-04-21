# encoding=utf-8
# coding=utf-8

# author:long
# description:学习使用json数据的解析

import json
import csv
# 读取数据

f = open('jsondata.txt', 'r', encoding='utf-8')

jsondata =''
temp = f.readline()
while temp:
    jsondata +=temp
    temp = f.readline()

# print(jsondata)

midf = open('index.txt', 'r', encoding='utf-8')
mid = midf.readline()

geturl = "https://space.bilibili.com/ajax/member/getSubmitVideos?mid="+str(mid)+\
                "&pagesize=30&tid=0&page="+\
                str(page)+"&keyword=&order=pubdate"

s = json.loads(jsondata)

# print(s['data']['vlist'])

print("read success")

pages = s['data']['pages']
if pages > 1:
    for page in range(2, pages):
        geturl =
        # 获得数据


        print(page)



csvdata = open('data.csv', 'w+')
csvwriter = csv.writer(csvdata)
csvwriter.writerow(('URL', 'title', 'length'))


def saveData(jsondata):
    for item in jsondata:
        url ="https://www.bilibili.com/video/av"+str(item['aid'])
        title = item['title']
        length = item['length']
        csvwriter.writerow((url, title, length))
        # print(url)
        # print(title)
        # print(length)


saveData(s['data']['vlist'])

