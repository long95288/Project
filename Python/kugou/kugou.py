# encoding=utf-8
# 引入包
import requests
from bs4 import BeautifulSoup
import time
import random
header = {
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) '
                  'Chrome/71.0.3578.98 Safari/537.36'
}


# 获得头文件，随机生成
def get_header():

    return header


# 打开测试文件,模拟用的
def open_test_file(file_url):
    file = open(file_url, 'r', encoding='utf-8')
    return file.read()


# 打开网页，获得网页数据
def get_web_content(url):
    res = requests.get(url, headers=get_header())
    return res.text


# 分析测试文件,获得相应的数据
# 数据格式：排名 歌名 时长
def analyze(txt):
    # 变量区
    soup = BeautifulSoup(txt, 'html.parser')
    redata = []
    # 选择器 #rankWrap > div.pc_temp_songlist > ul >li
    datas = soup.select('#rankWrap > div.pc_temp_songlist > ul > li')

    # 解析list数据
    for data in datas:
        # 获取排名
        if data.find('strong') != None:
            num = data.find('strong').string  # 前三名加粗的处理
        else:
            num = data.find('span', 'pc_temp_num').string.strip()  # 排名
            # print("%r" % num)
        # 获取歌名
        songname = data.find('a', 'pc_temp_songname').string.strip()

        # 获取时间
        song_time = data.find('span', 'pc_temp_time').string.strip()

        # 生成数据redata
        temp_data = num + "----" + songname + "----" + song_time
        # print(temp_data)
        redata.append(temp_data)
    return redata


# 写回数据
def write_data(datas):
    f = open('song.txt', 'a', encoding='utf-8')  # 文件追加写入
    for data in datas:
        f.write(data + '\n')
    f.close()
    print("写入成功！")


# 单页爬取内容
def crawl_page(pageurl):
    txt = get_web_content(pageurl)
    datas = analyze(txt)
    write_data(datas)
    print("爬取成功")


# 函数入口,开始爬取
def start():
    # 生成爬取页面数组
    pages = ['https://www.kugou.com/yy/rank/home/{}-8888.html'.format(number)
             for number in range(1, 24)]
    # print(pages)
    for pageurl in pages:
        print(pageurl)
        crawl_page(pageurl)
        time.sleep(2+random.random())  # 暂停避免过快


# 调用开始函数开始爬取
start()

# 测试
# txt = open_test_file('test2.html')
# print(txt)
# datas = analyze(txt)
# write_data(datas)

# print(random.random())
