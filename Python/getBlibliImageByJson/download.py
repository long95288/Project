
import requests
import json
import time
import random
from getBlibliImageByJson.MySqlUtil import queryUndownloadDynamic
from getBlibliImageByJson.MySqlUtil import queryImagesByDidAndUid
from getBlibliImageByJson.MySqlUtil import updateDynamicToDownloaded
from getBlibliImageByJson.Log import log

from getBlibliImageByJson.ImageUtil import downloadImageByURL

"""
下载的主类
"""
if __name__ == '__main__':
    # 获得未下载的动态
    timeout = 1
    log("获得10个未下载的动态")
    undownload_dynamic_list = queryUndownloadDynamic(10)
    while len(undownload_dynamic_list) > 0:
        # 下载
        for dynamic in undownload_dynamic_list:
            dynamic_id = dynamic[0]
            uid = dynamic[1]
            # 获得动态的图片列表的url
            log("下载:dynamic_id = {}".format(dynamic_id))
            image_url_list = queryImagesByDidAndUid(uid, dynamic_id)
            for image in image_url_list:
                url = str(image).replace("http:", "")
                url = url.replace("https:", "")
                url = "https:{}".format(url)
                url = url.split("@")[0]
                downloadImageByURL(url)
                # 睡眠
                time.sleep(timeout+random.random())
            # 更新数据库
            log("更新数据库：uid:{},dynamic_id:{}".format(uid,dynamic_id))
            updateDynamicToDownloaded(uid=uid, dynamic_id=dynamic_id)

        log("获得10个未下载的动态")
        undownload_dynamic_list = queryUndownloadDynamic(10)
