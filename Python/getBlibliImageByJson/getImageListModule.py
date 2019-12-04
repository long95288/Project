"""
获得数据的主类
"""
import time
import json
import random
from getBlibliImageByJson.ImageUtil import *
from getBlibliImageByJson.MySqlUtil import *

"""
判断该动态是否已经存在了
"""
def dynamic_is_exist(uid,dynamic_id):
    result = queryRecordByUidAndDynamicId(uid, dynamic_id)
    if result:
        return True
    return False

if __name__ == '__main__':
    root = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?host_uid={}&offset_dynamic_id={}"
    uids = getUID()
    for uid in uids:
        offset = 0
        # 动态的url
        dy_url = root.format(uid, offset)
        has_more = 1
        # 需要插入的动态数据
        dynamic_list = []
        while has_more != 0:
            print("has_more = {}".format(has_more))
            if has_more == 0:
                break
            list, has_more, hast_uid = getImageUrlList(url=dy_url)
            for dy_item in list:
                dynamic_id = dy_item[0]
                dynamic_image_list = dy_item[1]
                # 判断该动态是否已经存在了？
                if dynamic_is_exist(uid, dynamic_id):
                    # 当前的动态已经存在,说明下面的动态也已经存在，不需要再判断了
                    log("uid:{} dynamic_id:{}已经存在".format(uid,dynamic_id))
                    print("uid:{} dynamic_id:{}已经存在".format(uid,dynamic_id))
                    has_more = 0
                    break
                else:
                    # 数据库中没有,添加到列表中等待插入数据
                    dynamic_list.append((uid, dynamic_id, dynamic_image_list))
                    # insertDynamic(uid=uid,dynamic_id=dynamic_id,image_list=dynamic_image_list)
            # 将数据插入数据库中
            insertDynamicList(dynamic_list)
            # 休眠两个秒后添加数据
            time.sleep(2 + random.random())
            # 更新数据获得
            dy_url = root.format(uid, hast_uid)
        message = "获得:uid:{}全部动态图片数据".format(uid)
        print(message)
        # 将数据写入数据库中
        log(message)
