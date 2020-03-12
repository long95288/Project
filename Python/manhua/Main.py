# encoding=utf-8
# coding=utf-8
import requests
import re
import time
from bs4 import BeautifulSoup
import random
import os
import json

headers = [
    {
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36'
    },
    {
        'user-agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36'
    }
]
# 网站根目录的url
root_url = 'https://m.qhysfe.com'

# 保存图片的根路径
base_save_dir = 'D:\\manhua\\'

# 日志处理
def log(massage):
  print(massage)

# 获得首页内容
# 返回页面的list
def get_index_info(url):
  header = random.choice(headers)
  result = []
  response = requests.get(url,header)
  if response.status_code == 200:
    response_data = response.content
    soup = BeautifulSoup(response_data,'html.parser')
    plist = soup.select('body > div.mBox > div.bd > ul > li')
    for item in plist:
      p_url = item.select('a.pic')[0].get('href')
      result.append(p_url)
  # print(result)
  return result

# 获得图片的内容
def get_image_content(url):
    header = random.choice(headers)
    response_data = None
    try:
        response = requests.get(url, header)
        if response.status_code == 200:
            response_data = response.content
        else:
            log("获取:{}失败".format(url))
    except RuntimeError:
        log("请求{}异常".format(url))

    return response_data

# 下载图片
# url 图片的地址
# save_dir 保存的路径
# 保存的名称
def download_image(url,save_dir,filename):
  # 创建文件夹
  image_content = get_image_content(url)
  if image_content is not None:
    filedir = save_dir + filename
    with open(filedir,'wb') as f:
      f.write(image_content)
      log("下载: {} 成功".format(url))
  else:
    log("下载: {} 失败".format(url))

# 获得html中的图片的url
def get_image_url(url):
  header = random.choice(headers)
  try:
    response = requests.get(url,header)
    if response.status_code == 200:
      response_data = response.content
      soup = BeautifulSoup(response_data,'html.parser')
      img_list = soup.select('#content > div.primary > p > a > img')
      if img_list is not None:
        url = img_list[0].get('src')
      
  except RuntimeError:
    log('请求:{} 异常'.format(url))
  else:
    return url
  return url

# 获得一本漫画的名称和url
# 第一个名称
def get_single_pic(url):
  header = random.choice(headers)
  response = requests.get(url,header)
  if response.status_code == 200:
    response_data = response.content
    soup = BeautifulSoup(response_data,'html.parser')
    title = soup.select('#content > div.primary > center > h1')[0].string
    total_number = soup.select('#content > div.showpage > a:nth-child(1)')[0].string
    number = int(re.findall('共(.*)页:',total_number)[0]) + 1
    # 获得封面,第一张图片
    cover_url = soup.select('#content > div.primary > p > a > img')[0].get('src')
    print(cover_url)
    title = title.replace(":","")
    title = title.replace(";","")
    save_dir = base_save_dir + title + "\\"

    # 创建该漫画的保存的文件夹
    if not os.path.exists(save_dir):
      os.mkdir(save_dir)
    
    # 文件扩展名
    file_extend = cover_url.split('.')[-1]
    image_name = "1."+file_extend
    print('save_dir '+save_dir)
    print(image_name)
    download_image(cover_url,save_dir,image_name)
    # 获得全部的url
    pic_base_url = url.split('.html')[0]
    for i in range(2,number):
      new_url = pic_base_url + "_" + str(i)+".html"
      # 获得新的数据

      image_url = get_image_url(new_url)
      
      if image_url is None:
        continue

      print("下载:{}".format(new_url))
      extend = image_url.split('.')[-1]
      new_image_name = str(i)+ "." + extend
      download_image(image_url,save_dir,new_image_name)
      time.sleep(3 + random.random())
    
    # print(url_list)
    # print(number)
    # print(title)
  # print(response_data)
  


if __name__ == '__main__':
    # 两个核心功能
    # 已经下载的列表
    downloaded_list = []
    with open('list.json','r') as f:
      json_data = f.read()
      f.close()
      downloaded_list = json.loads(json_data)

    # 1、根据主页获得漫画的列表
    index_url = 'https://m.qhysfe.com/xieemanhua/'

    for i in range(1,36):
      new_list_url = index_url + "list_1_"+str(i)+".html"
      print('列表url:{}'.format(new_list_url))
      url_list = get_index_info(new_list_url)
      for item in url_list:
        pic_url = root_url + item
        if pic_url in downloaded_list:
          print('{} 已经下载'.format(pic_url))
          continue
        else:
          # 未下载,下载
          print("未下载:{}".format(pic_url))
          get_single_pic(pic_url)
          # 下载完成,写入文件
          downloaded_list.append(pic_url)
          #
          print('写入文件')
          with open('list.json','w') as f:
            f.write(json.dumps(downloaded_list))
            f.close()
          
          time.sleep(1+random.random())