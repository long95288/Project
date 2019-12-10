
def testGetImageUrlList():
    from getBlibliImageByJson.ImageUtil import getImageUrlList
    from getBlibliImageByJson.MySqlUtil import insertListToDataBase
    url = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?host_uid=372365045&offset_dynamic_id=0"
    list, has_more, last_uid = getImageUrlList(url)
    print("list:")
    print(list)
    print("has_more:")
    print(str(has_more))
    print("last_uid:"+str(last_uid))
    # affect_count = insertListToDataBase(list)
    # print("affect_count " + str(affect_count))


"""
解析字符串
"""
def testParseJsonData():
    import json
    with open('temp.json','r',encoding='utf-8') as f:
        response_data = json.load(f)
        # print(json_data)
        has_more = response_data['data']['has_more']
        print("has_more:"+str(has_more))
        cards = response_data['data']['cards']
        cards_num = len(cards)
        list = []
        for i in range(0, cards_num):
            card = cards[i]
            card_type = card['desc']['type']
            # 专栏图片
            if card_type == 64:
                cv_id = card['desc']['uid']
                print("是CV Id = "+str(cv_id))
            else:
                # 非专栏card,获得图片数据
                card_str = card['card']
                # 解析为json
                card_json = json.loads(card_str)
                pictures = card_json['item']['pictures']
                # 获得pictures中的图片url加入列表中
                for picture in pictures:
                    image_src = picture['img_src']
                    list.append(image_src)
            # 获得最后一个card 的id
            if i == cards_num -1:
                last_uid = card['desc']['dynamic_id']
                print("最后一个动态ID = "+str(last_uid))

        print("has_more = " + str(has_more))
        print("last_uid = "+str(last_uid))
        print("list: ")
        print(list)
        #
        # card0 = cards[0]
        # # print(card0['card'])
        # card_str = str(card0['card'])
        # jsonData = json.loads(card_str)
        # id = jsonData['item']['id']
        # pictures = jsonData['item']['pictures']

        # redata = re.sub(r'\/','',card_str,count=0)
        # print(redata)
        # id_matcher = re.search(r'"id":\d+', redata,re.M|re.I)
        # if id_matcher:
        #     group = id_matcher.group()
        #     id = group.split(":")[-1]
        #     print("id: "+str(id))
        #     # print("id"+id_matcher.group())
        # else:
        #     print("Nothing Found")
        # pattern = r'\S+(?:.jpg))'
        # imagelist = re.findall('^http.jpg',redata)
        # if imagelist:
        #     print(imagelist)
        # else:
        #     print("Nothing Found")
        # json2 = json.loads(redata)
        # print(json2)
        #print("id:"+str(id))
        #card_str_json = json.load(redata,encoding='utf-8')
        #print(card_str_json)
        # for card in cards:
        #     type = card['desc']['type']
        #     card_str = card['card']
        #     #pictures = card_str['pictures']
        #     #print(pictures)
        #     print(card_str)

def testInserListTODataBase():
    from getBlibliImageByJson.MySqlUtil import insertListToDataBase
    list = [
        "https://i0.hdslb.com/bfs/album/9c9b351c28e9193427912650f2e2f90c1b22da78.png",
        "https://i0.hdslb.com/bfs/album/d712a81dcc69c84b2810e255a2e8373ed3fd1e22.png"
    ]
    affect_count = insertListToDataBase(list)
    print("影响的行数:"+str(affect_count))

def testQueryUndownlownImage():
    from getBlibliImageByJson.MySqlUtil import queryUndownloadImage

    list = queryUndownloadImage()
    print(list)

def testURLParse():
    url = "https://i0.hdslb.com/bfs/album/9c9b351c28e9193427912650f2e2f90c1b22da78.png"
    print(url.split("/")[-1])

def testGetUID():
    import json
    with open('conf.json','r') as f:
        json_data = json.load(f)
        print(json_data['uids'])

"""
测试获得cv的图片列表
"""
def testGetImageUrlListByCV():
    id = 4091501
    from getBlibliImageByJson.ImageUtil import getImageUrlListByCV

    list = getImageUrlListByCV(id)
    print(list)

def testUID():
    from getBlibliImageByJson.MySqlUtil import getUID
    uid = getUID()
    print(uid)

def testQueryImagesByDidAndUid():
    from getBlibliImageByJson.MySqlUtil import queryImagesByDidAndUid
    uid = 299876758
    dynamic_id = 312545637717924441
    list = queryImagesByDidAndUid(uid, dynamic_id)
    print(list)

def testUpdateDynamicToDownloaded():
    from getBlibliImageByJson.MySqlUtil import updateDynamicToDownloaded
    dynamic_id = 328910025756276615
    uid = 27642052
    print("影响:{}行".format(updateDynamicToDownloaded(uid,dynamic_id)))

def testQueryUndownloadDynamic():
    from getBlibliImageByJson.MySqlUtil import queryUndownloadDynamic
    list = queryUndownloadDynamic(10)

    print(list)

def testaDownloadStatus():
    # 设置已经下载图片的数据库
    # import os
    # import pymysql
    #
    # root_path = "D:\\bibliimage\\"
    # # 获得全部图片的文件名
    # list = os.listdir(root_path)
    # list_num = len(list)
    # insert_num = 0
    # db = pymysql.connect("localhost", "root", "root2037", "blibli")
    # cursor = db.cursor()
    # sql = "UPDATE t_image SET status = 1 WHERE id = {}"
    #
    # for i in list:
    #     sql = sql.format(i)
    #     cursor.execute(sql)
    # db.commit()
    # while insert_num < list_num:
    #     insert_num += 100
    #     insert_list = []
    #     if insert_num > list_num:
    #         insert_list = list[insert_num - 100: list_num]
    #     else:
    #         insert_list = list[list_num-100:insert_num]
    #
    #     count = cursor.executemany(sql, insert_list)
    #     db.commit()
    #     print("影响行数:{}".format(count))
    #     break
    # print("文件长度:{}".format(len(list)))
    # sql = "UPDATE t_image SET status = 1 WHERE id = %s"
    # db = pymysql.connect("localhost", "root", "root2037", "blibli")
    # cursor = db.cursor()
    # count = cursor.executemany(sql,list)
    # db.commit()
    # print("更新数目:{}".format(count))
    pass

def testGetImageList():
    from getLastestImages import getLastestImageByUid

    uid = 388275388
    getLastestImageByUid(uid)