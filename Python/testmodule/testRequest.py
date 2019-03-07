# encoding=utf-8

import requests
from bs4 import BeautifulSoup

# 请求头部信息
header = {
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko)'
                  ' Chrome/71.0.3578.98 Safari/537.36'
}
url = 'http://www.022003.com/2_2447/973170.html'
res = requests.get(url, headers=header)
if res.status_code == 200:
    soup = BeautifulSoup(res.text, 'html.parser')
    print(soup.prettify())
    # item = soup.find_all('div', 'content')
    #content = soup.select('#content')
    # #content
    # print(content)
else:
    pass
#try:
    # 正常执行的代码
    # soup = BeautifulSoup(res.text, 'html.parser')
    # print("格式化后的数据")
    # print(soup.prettify('gbk'))
    # 查找
    # item = soup.find_all('img', 'lazy')
    # print(item)
    # hrefs = soup.select('img > src')
    # print(hrefs)
    # print(res)
    # selector 使用
    # print(res.text)
#except ConnectionError:  # 连接出错时可以用这个
#    print("连接出错了")
