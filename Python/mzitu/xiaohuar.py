# encoding=utf-8

import requests
import time
from bs4 import BeautifulSoup
# 地址

url = 'http://www.mzitu.com/'
# url = 'http://www.xiaohuar.com/'
# 头部
header = {
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) '
                  'Chrome/71.0.3578.98 Safari/537.36'
}


# 通过img_url来保存图片
def save_img(img_url):

    img_response = requests.get(img_url, headers=header, timeout=30)
    print(img_response)
    # 图片的内容
    img_data = img_response.content
    # print(img_data)
    # 图片的名称
    img_name = img_url.split('/')[-1]
    print(img_name)
    # urllib.request.urlretrieve(img_url, img_name)
    # 创建图片文件
    with open(img_name, 'wb') as f:
        # 写入数据
        f.write(img_data)
        f.close()


# 1.解析网页，获得图片的url地址
res = requests.get(url, headers=header)
soup = BeautifulSoup(res.text, 'html.parser')

# 获得图片的url
img_urls = soup.find_all('img', 'lazy')
#print(img_urls)
save_img('http://111.231.133.229:8888/long/document/77ed8bfaecbf87df0f2093708dcf487c.jpg')
save_img('https://i.meizitu.net/2019/01/06c01.jpg')


"""
# 遍历地址，把图片下载下来
for img_url in img_urls:
    # print(img_url.get('src'))
    tempurl = img_url.get('data-original')
    print("图片地址:"+tempurl)
    # 调用保存图片函数
    save_img(tempurl)
    time.sleep(2)

"""
