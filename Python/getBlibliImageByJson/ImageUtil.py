"""
blibli 动态图片和专栏图片的工具
"""

# 获得动态的接口
# https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?host_uid=36802028&offset_dynamic_id=0

"""
host_uid 是用户的id
offset_dynamic=0 是动态的偏移量，如果后面有数据，has_more 字段为1，如果没有数据字段为0,偏移的id是上一次是最后的动态的id
desc.type = 8 视频动态
desc.type = 64 专栏图片 card.id 为cv的id
desc.type = 2 普通动态图片
"""
import requests
import json
import time
import random
import datetime
from getBlibliImageByJson.Log import log
from bs4 import BeautifulSoup

headers = [
    {
        'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36'
    },
    {
        'user-agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36'
    }
]
# 保存blibli的图片
day = datetime.datetime.now().strftime('%Y%m%d')
save_path = "D:\\bibliimage2\\" + day +"\\"
"""
获得动态图片的url列表
"""
def getImageUrlList(url):
    # 返回值
    list = []
    # 最后的id
    last_uid = 0
    # 是否还有新的
    has_more = 0
    # 根据URL请求数据
    header = random.choice(headers)
    log("请求:{}\n".format(url))
    response = requests.get(url, header)
    if response.status_code == 200:
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
                        card_type = card['desc']['type']
                        # 专栏图片
                        if card_type == 64:
                            cv_id = card['desc']['rid']
                            print("是CV Id = " + str(cv_id))
                            cv_image_list = getImageUrlListByCV(id=cv_id)
                            dy_id = card['desc']['dynamic_id']
                            # 添加进去
                            list.append((dy_id, cv_image_list))
                            # for img in cv_image_list:
                            #     list.append(img)
                        elif card_type == 2:
                            # 非专栏card,获得图片数据
                            card_str = card['card']
                            # 解析为json
                            card_json = json.loads(card_str)
                            item = card_json['item']
                            if item:
                                pictures = item['pictures']
                                # 获得pictures中的图片url加入列表中
                                if pictures:
                                    dy_id = card['desc']['dynamic_id']
                                    img_list = []
                                    for picture in pictures:
                                        image_src = picture['img_src']
                                        img_list.append(image_src)
                                        # list.append(image_src)
                                    list.append((dy_id, img_list))
                        # 获得最后一个card 的id
                        if i == cards_num - 1:
                            last_uid = card['desc']['dynamic_id']
                            print("最后一个动态ID = "+str(last_uid))
        else:
            log("请求:--{}--无数据\n".format(url))
    else:
        # 失败写入日志
        log_message = "请求--{}--失败\n".format(url)
        log(log_message)
    return list, has_more, last_uid


"""
根据CV的id,获得该cv内所有的图片连接
return:
    list: 图片列表
"""
def getImageUrlListByCV(id):
    url = "https://www.bilibili.com/read/cv{}".format(id)
    header = random.choice(headers)
    log("请求CV:{}".format(url))
    response = requests.get(url, header)
    list = []
    if response.status_code == 200:
        response_data = BeautifulSoup(response.content.decode('utf-8'), 'html.parser')
        image_box = response_data.select('.img-box > img')
        if image_box:
            for item in image_box:
                src = "https:" + item.attrs['data-src']
                list.append(src)
            #
            log("获得CV:{}图片成功".format(url))
        else:
            log("获得CV:{}图片失败".format(url))
    else:
        log("请求CV:{}失败".format(url))
    # 防止过快请求
    time.sleep(1)
    return list


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


"""
保存图片文件
"""
def save_image(image_name,image_conten):
    filename = save_path + image_name
    with open(filename, 'wb') as fp:
        fp.write(image_conten)
        log("保存:{}成功".format(image_name))
        fp.close()
        # print("保存:"+image_name+"成功")


def downloadImageByURL(url):
    image_name = str(url).split("/")[-1]
    image_content = get_image_content(url)
    if image_content is not None:
        save_image(image_name, image_content)
