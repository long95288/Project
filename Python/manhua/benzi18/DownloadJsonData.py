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
root_url = 'https://18comic.one'
index_url = 'https://18comic.one/albums/doujin?o=tf'
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
    plist = soup.select('#wrapper > div.container > div > div.col-xs-12.col-md-9.col-sm-8 > div.row > div > div > a')
    for item in plist:
      p_url = root_url + item.get("href")
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

# 获得单个本子的标题和对应的url
def get_single_pic_url(url):
    url = url.replace("album", "photo")
    if True:
        response_data = htmlContent(url)
        if response_data is None:
            log("!!!获取url:{}失败".format(url))
            return
        waiting_download_list = re.findall("https://cdn-msp.18comic.one/media/photos/.*?/.*?.jpg", response_data, re.S)
        soup = BeautifulSoup(response_data, 'html.parser')
        titles = soup.select('#wrapper > div.top-nav.visible-xs > div > ul > span')
        if len(titles) == 0:
            log("!!!获取url:{}失败".format(url))
            return str(random.random() * 100 * random.random() * 100),waiting_download_list
        title = titles[0].string
        if None is title:
            return str(random.random() * 100 * random.random() * 100),waiting_download_list
        title = title.replace(":", "")
        title = title.replace(";", "")
        title = title.replace("\\", "").replace("/", "").replace("*", "").replace("?", "").replace("\"", "").replace(
            "<", "").replace(">", "").replace("|", "")
        return title, waiting_download_list

if __name__ == '__main__':
    # init()
    # 两个核心功能
    # 已经下载的列表
    downloaded_list = []
    with open('list.json','r') as f:
      json_data = f.read()
      f.close()
      downloaded_list = json.loads(json_data)
    # 1、根据主页获得漫画的列表
    totalPageNumber = 65535
    for i in range(6, totalPageNumber):
      new_list_url = index_url + '&page={}'.format(i)
      if new_list_url in downloaded_list:
          print("已经下载{}".format(new_list_url))
          continue
      print('列表url:{}'.format(new_list_url))
      url_list = get_index_info(new_list_url)
      if len(url_list) == 0:
        break
      # json数据
      json_list = []
      for item in url_list:
          pic_url = item
          # 未下载,下载
          title,url_list = get_single_pic_url(pic_url)
          if url_list is None:
              continue
          if len(url_list) <= 0:
              continue
          one_pic = {
              "title":title,
              "url":url_list
          }
          json_list.append(one_pic)
          # 下载完成,写入文件
          time.sleep(1+random.random())
      if len(json_list) > 0:
        with open("data\\image_{}_data_{}.json".format(day,i),"a",encoding="utf-8") as f:
          f.write(json.dumps(json_list, ensure_ascii=False))
          f.close()
        downloaded_list.append(new_list_url)

      with open('list.json', 'w') as f:
          f.write(json.dumps(downloaded_list))
          f.close()
