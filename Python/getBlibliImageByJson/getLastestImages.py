"""
这个是获得最新的image数据的主类
"""
from ImageUtil import *
from MySqlUtil import *
import random
import time
import requests
import sched

headers = [
    {
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36'
    },
    {
        'user-agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36'
    }
]
"""
判断该动态是否已经存在了
"""
def dynamic_is_exist(uid,dynamic_id):
    result = queryRecordByUidAndDynamicId(uid, dynamic_id)
    if result:
        return True
    return False
"""
获得一个UID的最新的数据并且写入数据库中
"""
def getLastestImageByUid(uid):
    # 请求的网址
    root = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?host_uid={}&offset_dynamic_id={}"
    offset = 0
    # 动态的URL
    has_more = 1
    timeout = 1
    while has_more != 0:
        if has_more == 0:
            break
        # 需要插入的数据
        dynamic_list = []
        # 请求的URL
        url = root.format(uid, offset)
        # 请求获得动态的数据
        header = random.choice(headers)
        log("请求:{}".format(url))
        time.sleep(timeout + random.random())
        response = requests.get(url, header)
        if response.status_code == 200:
            # 加载数据为json
            response_data = json.loads(response.text)
            data = response_data['data']
            if data:
                # has_more
                has_more = data['has_more']
                # list
                if has_more != 0:
                    cards = data['cards']
                    if cards:
                        cards_num = len(cards)
                        for i in range(0, cards_num):
                            card = cards[i]
                            # 卡片的类型
                            card_type = card['desc']['type']
                            # 动态的id
                            dynamic_id = card['desc']['dynamic_id']
                            # 专栏图片
                            if card_type == 64:
                                # 判断该专栏是否已经纯在数据库中了:
                                if dynamic_is_exist(uid=uid, dynamic_id=dynamic_id):
                                    # 该专栏已经处在数据库中了，就不用再往下请求了
                                    log("动态:{} 已经存在".format(dynamic_id))
                                    has_more = 0
                                    break
                                else:
                                    # 该专栏还还未在数据库中，获得该专栏图片
                                    cv_id = card['desc']['rid']
                                    cv_image_list = getImageUrlListByCV(id=cv_id)
                                    # 将数据放进列表中
                                    dynamic_list.append((uid, dynamic_id, cv_image_list))
                                    log("获得专栏CV{} 的数据".format(cv_id))

                            elif card_type == 2:
                                if dynamic_is_exist(uid=uid, dynamic_id=dynamic_id):
                                    # 已经存在,退出
                                    log("动态:{} 已经存在".format(dynamic_id))
                                    has_more = 0
                                    break
                                else:
                                    # 非专栏card,获得图片数据
                                    card_str = card['card']
                                    # 解析为json
                                    card_json = json.loads(card_str)
                                    item = card_json['item']
                                    if item:
                                        pictures = item['pictures']
                                        # 获得pictures中的图片url加入列表中
                                        if pictures:
                                            img_list = []
                                            for picture in pictures:
                                                image_src = picture['img_src']
                                                img_list.append(image_src)
                                            # 添加动态数据
                                            dynamic_list.append((uid, dynamic_id, img_list))
                                            log("获得动态数据dynamic_id:{}".format(dynamic_id))
                            else:
                                log("未定义类型:{}".format(card_type), False)

                            # 判断退出
                            if has_more == 0:
                                break

                            # 获得最后一个card 的id
                            if i == cards_num - 1:
                                offset = card['desc']['dynamic_id']
            else:
                log("请求:{} 无数据\n".format(url))

        log("请求:{} 成功".format(url))
        dynamic_len = len(dynamic_list)
        if dynamic_len > 0:
            # 将请求到的数据写入数据库中
            log("写入{}条动态数据.................".format(len(dynamic_list)))
            insertDynamicList(dynamic_list)
            log("写入数据完成.............")
        else:
            log("动态数据为空.............")

"""
开始获得最新的数据
"""
def start_get_latest_images():
    # uid的列表
    uids = getUID()
    for uid in uids:
        log("正在获得uid:{}的最新图片".format(uid))
        getLastestImageByUid(uid)
        log("已经获取uid:{}的全部最新图片".format(uid))
    pass

"""
开始下载图片
"""
def start_download_images():
    # 下载
    log("开始下载......")
    timeout = 1
    log("获得数据库数据")
    un_download_dynamic_list = queryUndownloadDynamic(10)
    while len(un_download_dynamic_list) > 0:
        # 下载
        for dynamic in un_download_dynamic_list:
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
                time.sleep(timeout + random.random())
            # 更新数据库
            log("更新数据库：uid:{},dynamic_id:{}".format(uid, dynamic_id))
            updateDynamicToDownloaded(uid=uid, dynamic_id=dynamic_id)

        log("获得数据库数据......")
        un_download_dynamic_list = queryUndownloadDynamic(10)

"""
启动
"""
if __name__ == '__main__':
    start_get_latest_images()
    start_download_images()



