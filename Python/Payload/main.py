# encoding=utf-8
# coding=utf-8

import json
import csv
import requests
import time
import random

header = {
    'content-type': 'application/json'
}
headers = {
    'Origin': 'https://www.kaggle.com',
    'Referer': 'https://www.kaggle.com/',
    'Host': 'www.kaggleusercontent.com',
    'Content-Type': 'application/json',
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) '
                  'Chrome/71.0.3578.98 Safari/537.36',
    'Accept': '*/*',
    'Cache-Control': 'no-cache',
    'Agent-Token': '592c2f33-3ff3-4c40-8c88-99677e01173f,4c3bcf86-cca6-4f8f-be02-a35ae370c5db',
    'Host': 'www.kaggleusercontent.com',
    'accept-encoding': 'gzip, deflate',
    'content-length': '1367',
    'Connection': 'keep-alive',
}
timeout = 2000 # 超时两秒
pageUrl = 'https://www.kaggleusercontent.com/services/datasets/kaggle.dataview.v1.DataViewer/GetDataView'
payloadData ='''
    {
    "jwe": {
        "encryptedToken": "eyJhbGciOiJkaXIiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2In0..UZF_MtdCj7iugGBTcFVDVw.52l2kfUGwQbEtkTCuNrtPEeNUotAyu5ilpt8FD3PVf679rpLAdj0W-FIyoO4rcvOgEfNVXkVIUvsyImHyJV7-mWpxtvdoBjaB1M8v-xSOUvQ4Evx19gRTqo_HHdoEtSMxKAlMoGobomo6eNKw_6Pol2gZ7GuV7cc83qDaxvrUh1herCnPAo2nXkU-lhg9h5hLiL0iYNfouWMZjWt8CsHfb6dB24e4UmkLg-rhu_Qx8WFvGd0O7OMfIxl0Q1lAMP4gnBIpvLGvO3WPfaqhQE2PFtCo8N5JBJT3dC5-weKEKjCatWwhuHeBDkNKgNkO6PNwzZWdhJFkCair8XdpeeOm7378JeXSRtDB3Iip4HsSzvJ29gwqsDd6hsksFDhXYTN2oJCioJs0896unZlvYKnwCURzEnuBf0zeuqGAeO2jpAoks6F3MtwgeO55txul2i8zHQuwJySFe0URZiv8b_Q8UeBSXJVzqWZLuJR0xiFqsemjPmEV55GJyTH9cJi9SfVai54gY-uZInub_f0Sz3gmMexsI2uomeS667aEoPupRY.CI6gaw_g-okeovcaz72c5A"
    },
    "source": {
        "type": 3,
        "dataset": {
            "url": "russellyates88/suicide-rates-overview-1985-to-2016",
            "tableType": 1,
            "csv": {
                "fileName": "master.csv",
                "delimiter": ",",
                "headerRows": 1
            }
        }
    },
    "operations": null,
    "select": [
        "country",
        "year",
        "sex",
        "age",
        "suicides_no",
        "population",
        "suicides/100k pop",
        "country-year",
        "HDI for year",
        " gdp_for_year ($) ",
        "gdp_per_capita ($)",
        "generation"
    ],
    "skip": 0,
    "take": 100
}
'''
# 
def getTextData(rawData):
    res = requests.post(pageUrl,data=rawData,headers=headers,timeout=timeout)
    return res.text

csvf = open('data.csv','w+',encoding='utf-8')
csvwriter = csv.writer(csvf)
csvwriter.writerow(('rowNumber','country',
'year','sex','age','suicides_no','population','suicides/100k pop',
'country-year','HDI for year'," gdp_for_year ($) ",
        "gdp_per_capita ($)",
        "generation"))
# nextskip = 0
# res = requests.post(pageUrl,data=payloadData,headers=headers,timeout=timeout)
# print(json.dumps(payloadData))
#print(res.text)
# print(json.dumps(payloadData))
def saveData(jsondata):
    for item in jsondata:
        rowNumber = item['rowNumber']
        text = item['text']
        country = text[0]
        year = text[1]
        sex = text[2]
        age = text[3]
        suicides_no = text[4]
        population = text[5]
        suicides_100k = text[6]
        country_year = text[7]
        HDI = text[8]
        gdp_year = text[9]
        gdp_caption = text[10]
        generation = text[11]
        csvwriter.writerow((rowNumber,country,year,sex,age,suicides_no,
        population,suicides_100k,country_year,HDI,gdp_year,gdp_caption,generation))

def start(payloadData):
    jsontext = getTextData(payloadData)
    i2 = 1500 # 爬去的数量
    nextskip = 0
    for i in range(1,i2):
        s = json.loads(jsontext) # 转换为json格式
        writejson = s['dataView']['rows']
        saveData(writejson) # 写入数据
        # 设置新的payload
        nextskip = i*100
        payloadobj = json.loads(payloadData)
        payloadobj['skip'] = nextskip ## 下一个请求
        payloadData2 = str(json.dumps(payloadobj))
        time.sleep(2+random.random())
        payloadData =payloadData2
        jsontext = getTextData(payloadData) # 开始下一个请求
    # jsonobj = json.loads(payloadData) # 转换成json对象
    # skip = jsonobj['skip']
    # print(skip)
    # jsonobj['skip'] = 100 # 赋值
    # print(jsonobj['skip'])
    #s = json.loads(jsondata)
    #jsondata=s['dataView']['rows']
    #saveData(jsondata)

# 开始爬取
start(payloadData)