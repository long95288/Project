# encoding=utf-8
# coding=utf-8
import requests
import re
import time
from bs4 import BeautifulSoup
import random
import datetime
import os
import json

import sys

headers = [
    {
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36'
    },
    {
        'user-agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36'
    }
]

day = datetime.datetime.now().strftime('%Y%m%d')
log_file = day+"_log.txt"
rpc_url = "http://localhost:6800/jsonrpc"

# 网站根目录的url
root_url = 'https://m.lfkmhw.com/'
index_url = ''
list_prefix = ""
totalPageNumber = 0
images_pattern = ''
image_pattern = ''

script_tmp_path = os.getcwd()
script_tmp_name = "tmp.html"

# 保存图片的根路径
base_save_dir = 'D:\\manhua\\'

# 日志处理
def log(value,print_flag = True):
    logfile = open(log_file, 'a', encoding='utf-8')
    if logfile.writable():
        now_data = datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')
        log_message = "时间:{} : log : {}\n".format(now_data, value)
        if print_flag:
            print(log_message)
        logfile.write(log_message)
    try:
        logfile.close()
    except IOError:
        print("写入日志错误")
    else:
        return


def addDownloadTask(url,dir,out):
    postdata = {
        "jsonrpc": "2.0",
        "id": "QXJpYU5nXzE1NDgzODg5MzhfMC4xMTYyODI2OTExMzMxMzczOA=="
    }
    rpc_request = postdata
    rpc_request["method"] = "aria2.addUri"
    # rpc 的选项，去掉--就可以了
    options = {
        "dir":dir,
        "out":out,
        "allow-overwrite":"true"
    }
    rpc_request["params"] = [[url],options]
    response = requests.post(url=rpc_url, json=rpc_request)
    if response.status_code == 200:
        result = response.json().get("result", [])
        print("gid: {}".format(result))
        return result
    else:
        log("无法调用aria2")

def download_status(gid):
    postdata = {
        "jsonrpc": "2.0",
        "id": "QXJpYU5nXzE1NDgzODg5MzhfMC4xMTYyODI2OTExMzMxMzczOA=="
    }
    rpc_request = postdata
    rpc_request["method"] = "aria2.tellStatus"
    rpc_request["params"] = [gid]
    response = requests.post(url=rpc_url,json=rpc_request)
    if response.status_code == 200:
        result = response.json().get("result","")
        if result != "":
            status = result.get("status")
            if status != "":
                return status
    return None

"""
下载工具
url: 下载地址
dir: 保存路径
out: 保存名称
"""
def download(url,dir,out):
    log("开始下载:{}".format(url))
    gid = addDownloadTask(url,dir,out)
    status = download_status(gid)
    error = False
    error_num = 0
    while True and not error:
        if status == "active":
            time.sleep(3)
            print("下载中.....\n")
            status = download_status(gid)
        if status == "complete":
            break
        elif status == "waiting":
            log("下载队列已满")
            time.sleep(4)
            status = download_status(gid)
        elif status == "paused":
            log("暂停下载")
            break
        elif status == "error":
            log("下载错误")
            if error_num == 3:
                error = True
                break
            else:
                log("重新下载")
                gid = addDownloadTask(url,dir,out)
                status = download_status(gid)
                error_num = error_num + 1
            
        elif status == "removed":
            log("已经从下载队列中移除")
            break
    if error:
        log("下载:{}出错".format(url))
        return -1
    else:
        log("下载{}成功".format(url))
        return 0

# 获得首页内容
# 返回页面的list
def get_index_info(url):
  result = []
  if True:
    response_data = htmlContent(url)
    soup = BeautifulSoup(response_data,'html.parser')
    plist = soup.select('div.bd > ul >li')
    for item in plist:
      p_url = item.select('a.pic')[0].get('href')
      result.append(p_url)
  return result


# 下载图片
# url 图片的地址
# save_dir 保存的路径
# 保存的名称
def download_image(url,save_dir,filename):
    download(url=url,dir=save_dir,out=filename)

"""
获得html文件的内容
"""
def htmlContent(url):
    status = download(url=url,dir=script_tmp_path,out=script_tmp_name)
    data = None
    if status == 0:
        # 读取文件
        with open(script_tmp_path+"\\"+script_tmp_name,"r",encoding="utf-8") as f:
            data = f.read()
            f.close()
        return data

# 获得html中的图片的url
def get_image_url(url):
  imgurl = None
  try:
    if True:
      response_data = htmlContent(url)
      imgs = re.findall(images_pattern,response_data,re.S)
      # 启用备用解析
      if len(imgs) ==0 :
        bpattern = '<a .*?><img src="http://lf.mz0731.com/uploads/.*?" .*?></a>'
        imgs = re.findall(bpattern,response_data,re.S)
      # print(imgs)
      if len(imgs) > 0:
        # image_pattern = 'src="http://lf.veestyle.cn/uploads/.*?"'
        urls = re.findall(image_pattern,imgs[0])
        if len(urls) > 0:
          imgurl = urls[0].split("src=")[-1].replace('\"',"")
          print('图片url: {}'.format(imgurl))
        elif len(urls) == 0:
          # 启用备用解析
          urls = re.findall('src="http://lf.mz0731.com/uploads/.*?"',imgs[0])
          imgurl = urls[0].split("src=")[-1].replace('\"',"")
          # 替换成当前的地址
          imgurl = imgurl.replace("http://lf.mz0731.com/",root_url)
          print("备用解析图片url:{}".format(imgurl))

  except RuntimeError:
    log('请求:{} 异常'.format(url))
  else:
    return imgurl
  return imgurl

# 获得一本漫画的名称和url
# 第一个名称
def get_single_pic(url):
  if True:
    response_data = htmlContent(url)
    # print(response_data)
    if response_data is None:
        log("!!!获取url:{}失败".format(url))
        return
    soup = BeautifulSoup(response_data,'html.parser')
    titles = soup.select('#content > div.primary > center > h1')
    if len(titles) == 0:
        log("!!!获取url:{}失败".format(url))
        return
    title = titles[0].string
    cover_url = get_image_url(url)
    print("封面:{}".format(cover_url))
    title = title.replace(":","")
    title = title.replace(";","")
    title = title.replace("\\","").replace("/","").replace("*","").replace("?","").replace("\"","").replace("<","").replace(">","").replace("|","")

    suffix = url.split(".html")[0].split("/")[-1];
    save_dir = base_save_dir + title + suffix + "\\"
    # 创建该漫画的保存的文件夹
    if not os.path.exists(save_dir):
      os.mkdir(save_dir)
    
    # 文件扩展名
    file_extend = cover_url.split('.')[-1]
    image_name = "1."+file_extend
    print('save_dir '+save_dir)
    print('封面:{}'.format(image_name))
    download_image(cover_url,save_dir,image_name)
    numbers = re.findall('共(.*)页:',response_data)
    if len(numbers) == 0:
        return;
    
    number = int(numbers[0]) + 1
    # 获得全部的url
    pic_base_url = url.split('.html')[0]
    download_list = []
    with open("imagelist.json","r") as f:
        download_list = json.loads(f.read())
        f.close()
    
    for i in range(2,number):
      new_url = pic_base_url + "_" + str(i)+".html"
      
      # 获得新的数据
      print("下载:{} 中的图片".format(new_url))
      image_url = get_image_url(new_url)
      if image_url is None:
        continue
      if image_url in download_list:
        print("已经下载:{}".format(image_url))
        continue
      extend = image_url.split('.')[-1]
      new_image_name = str(i)+ "." + extend
      print("下载图片:{}".format(image_url))
      download_image(image_url,save_dir,new_image_name)
      download_list.append(image_url)
      time.sleep(3 + random.random())
      # 完成一个之后
      with open("imagelist.json","w") as f:
          f.write(json.dumps(downloaded_list))
          f.close()

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
    global base_save_dir
    base_save_dir = conf.get("base_save_dir",base_save_dir)
    if not os.path.exists(base_save_dir):
      os.makedirs(base_save_dir)
    
  print("================conf===================")
  print("root_url:{}".format(root_url))
  print("index_url{}".format(index_url))
  print("list_prefix{}".format(list_prefix))
  print("totalPageNumber{}".format(totalPageNumber))
  print("images_pattern{}".format(images_pattern))
  print("image_pattern{}".format(image_pattern))
  print("base_save_dir{}".format(base_save_dir))
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
    # totalPageNumber = 2
    for i in range(1, totalPageNumber):
      new_list_url = index_url + str(list_prefix).format(i)
      print('列表url:{}'.format(new_list_url))
      url_list = get_index_info(new_list_url)
      if len(url_list) == 0:
        print("列表数据为空")
        break
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