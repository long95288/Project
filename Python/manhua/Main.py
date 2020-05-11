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
root_url = 'https://m.lfkmhw.com/'
index_url = ''
list_prefix = ""
totalPageNumber = 0
images_pattern = ''
image_pattern = ''

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
  imgurl = None
  try:
    response = requests.get(url,header)
    if response.status_code == 200:
      response_data = response.content.decode('utf-8')
      # print(response_data)
      # http://lf.veestyle.cn/uploads/2020/05sx/0511/tu2.jpg
      # images_pattern = '<a .*?><img src="http://lf.veestyle.cn/uploads/.*?" .*?></a>'
      imgs = re.findall(images_pattern,response_data,re.S)
      # print(imgs)
      if len(imgs) > 0:
        # image_pattern = 'src="http://lf.veestyle.cn/uploads/.*?"'
        urls = re.findall(image_pattern,imgs[0])
        if len(urls) > 0:
          imgurl = urls[0].split("src=")[-1].replace('\"',"")
          print('图片url: {}'.format(imgurl))
  except RuntimeError:
    log('请求:{} 异常'.format(url))
  else:
    return imgurl
  return imgurl

# 获得一本漫画的名称和url
# 第一个名称
def get_single_pic(url):
  header = random.choice(headers)
  response = requests.get(url,header)
  if response.status_code == 200:
    response_data = response.content
    soup = BeautifulSoup(response_data,'html.parser')
    
    title = soup.select('#content > div.primary > center > h1')[0].string

    number = int(re.findall('共(.*)页:',response_data.decode('utf-8'))[0]) + 1

    # 获得封面,第一张图片
    # pattern = '<p style="text-align: center;">.*?<a .*?><img src="http://3qwd.lzsysj.com/qhysfe/uploads/allimg/.*?" .*?></a>.*?</p>'
    # imgs = re.findall(pattern,response_data.decode('utf-8'),re.S)
    # if len(imgs) > 0:
    #   urls = re.findall('src="http://3qwd.lzsysj.com/qhysfe/uploads/allimg/.*?"',imgs[0])
    #   if len(urls) > 0:
    #     cover_url = urls[0].split("src=")[-1].replace('\"',"")

    cover_url = get_image_url(url)
    print("封面:{}".format(cover_url))
    
    title = title.replace(":","")
    title = title.replace(";","")
    title = title.replace("\\","").replace("/","").replace("*","").replace("?","").replace("\"","").replace("<","").replace(">","").replace("|","")

    save_dir = base_save_dir + title + "\\"
    # 创建该漫画的保存的文件夹
    if not os.path.exists(save_dir):
      os.mkdir(save_dir)
    
    # 文件扩展名
    file_extend = cover_url.split('.')[-1]
    image_name = "1."+file_extend
    print('save_dir '+save_dir)
    print('封面:{}'.format(image_name))
    download_image(cover_url,save_dir,image_name)
    # 获得全部的url
    pic_base_url = url.split('.html')[0]
    for i in range(2,number):
      new_url = pic_base_url + "_" + str(i)+".html"
      # 获得新的数据
      print("下载:{} 中的图片".format(new_url))

      image_url = get_image_url(new_url)
      if image_url is None:
        continue

      extend = image_url.split('.')[-1]
      new_image_name = str(i)+ "." + extend
      print("下载图片:{}".format(image_url))
      download_image(image_url,save_dir,new_image_name)
      time.sleep(3 + random.random())
    
  
def atestRe():
  str = """
  <p style="text-align: center;">
	<a href="381_3.html"><img src="http://3qwd.lzsysj.com/qhysfe/uploads/allimg/180713/liv3xjat0of687.jpg" title="里番库口工漫画:[秋葉魔王]h本子優等生の吉田さんは先生に監禁されて肉便器になりました。 " original="http://3qwd.lzsysj.com/qhysfe/uploads/allimg/180713/liv3xjat0of687.jpg"></a></p>
  """
  pattern = '<p style="text-align: center;">.*?<a .*?><img src="http://3qwd.lzsysj.com/qhysfe/uploads/allimg/.*?" .*?></a>.*?</p>'
  #imgs = re.findall('<img src="http://3qwd.lzsysj.com/qhysfe/uploads/allimg/.*?" title=.*? original="http://3qwd.lzsysj.com/qhysfe/uploads/allimg/.*?">',str)
  imgs = re.findall(pattern,str,re.S)
  print(imgs)
  url = re.findall('src="http://3qwd.lzsysj.com/qhysfe/uploads/allimg/.*?"',imgs[0])
  print(url[0].split("src=")[-1])
  pass

def testInit():
  init()

def init():
  with open('conf.json','r') as f:
    conf = json.loads(f.read())
    f.close()
    global root_url
    root_url = conf.get("root_url")
    global index_url
    index_url = conf.get("index_url")
    global list_prefix
    list_prefix = conf.get("list_prefix")
    global totalPageNumber
    totalPageNumber = conf.get("total_page_number")
    global images_pattern
    images_pattern = conf.get("images_pattern")
    global image_pattern
    image_pattern = conf.get("image_pattern")
  print("================conf===================")
  print("root_url:{}".format(root_url))
  print("index_url{}".format(index_url))
  print("list_prefix{}".format(list_prefix))
  print("totalPageNumber{}".format(totalPageNumber))
  print("images_pattern{}".format(images_pattern))
  print("image_pattern{}".format(image_pattern))
  print("=======================================")


if __name__ == '__main__':
    init()
    # 两个核心功能
    # 已经下载的列表
    downloaded_list = []
    with open('list.json','r') as f:
      json_data = f.read()
      f.close()
      downloaded_list = json.loads(json_data)

    # 1、根据主页获得漫画的列表
    # index_url = 'https://m.lfkmhw.com/shaonvmanhua/'
    # totalPageNumber = 245
    for i in range(1, totalPageNumber):
      new_list_url = index_url + str(list_prefix).format(i)
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