# encoding=utf-8
from bs4 import BeautifulSoup
import requests
import time


headers = {
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 '
                  '(KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36'
}


def judgment_sex(class_name):  # 用户性别
    if class_name == ['member_icol']:
        return '女'
    else:
        return '男'


def get_links(url):
    wb_data = requests.get(url, headers=headers)
    soup = BeautifulSoup(wb_data, 'lxml')
    links = soup.select('#page_list > ul > li > a')  # link为url列表
    for link in links:
        href = link.get("href")
        get_info(href)


def get_info(url):
    wb_data = requests.get(url, headers=headers)
    soup = BeautifulSoup(wb_data.text, 'lxml')
    # 取得标题
    titles = soup.select('body > div.wrap.clearfix.con_bg > div.con_l > div.pho_info > h4 > em')
    print(titles)
    # 取得地址
    address = soup.select('span.pr5')
    print(address)
    # 取得价钱
    prices = soup.select('#pricePart > div.day_l > span')
    print(prices)
    # 取得房东头像
    imgs = soup.select('#floatRightBox > div.js_box.clearfix > div.member_pic > div')
    # 房东的名字
    names =soup.select('#floatRightBox > div.js_box.clearfix > div.w_240 > h6 > a')
    # 房东的性别
    sexs = soup.select('#floatRightBox > div.js_box.clearfix > div.member_pic > div')

    # 测试得到的数据

'''   
if _name_ == '_main_':
    urls = ['http://bj.xiaozhu..com/search-duanzufang-p{ }-0'.format
            (number) for number in range(1, 14)]
    for single_url in urls:
        get_links(single_url)
        time.sleep(2)
'''

myurl = 'http://bj.xiaozhu.com/fangzi/30308825803.html'
get_info(myurl)

