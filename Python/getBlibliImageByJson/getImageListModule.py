"""
获得数据的主类
"""
import time
import json
import random
from getBlibliImageByJson.ImageUtil import *
from getBlibliImageByJson.MySqlUtil import *

if __name__ == '__main__':
    root = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?host_uid={}&offset_dynamic_id={}"
    with open('conf.json','r') as f:
        conf = json.load(f)
        uids = conf['uids']
    for uid in uids:
        offset = 0
        # 动态的url
        dy_url = root.format(uid, offset)
        has_more = 1
        while has_more != 0:
            print("has_more = {}".format(has_more))
            if has_more == 0:
                break
            list, has_more, hast_uid = getImageUrlList(url=dy_url)
            affect_count = insertListToDataBase(list)
            time.sleep(2 + random.random())
            print("插入数据:"+str(affect_count))
            dy_url = root.format(uid, hast_uid)
        print("获取:uid "+str(uid)+"全部数据")
