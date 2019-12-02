
def testGetImageUrlList():
    from getBlibliImageByJson.ImageUtil import getImageUrlList
    url = "https://api.vc.bilibili.com/dynamic_svr/v1/dynamic_svr/space_history?host_uid=326544280&offset_dynamic_id=0"
    list, has_more, last_uid = getImageUrlList(url)
    print("list:")
    print(list)
    print("has_more:")
    print(has_more)
    print("last_uid:"+str(last_uid))

"""
解析字符串
"""
def testParseJsonData():
    import json
    import re
    with open('temp.json','r') as f:
        json_data = json.load(f)
        # print(json_data)
        has_more = json_data['data']['has_more']
        print("has_more:"+str(has_more))
        cards = json_data['data']['cards']
        card0 = cards[0]
        # print(card0['card'])
        card_str = str(card0['card'])
        redata = re.sub(r'\/','',card_str,count=0)
        print(redata)
        id_matcher = re.search(r'"id":\d+', redata,re.M|re.I)
        if id_matcher:
            group = id_matcher.group()
            id = group.split(":")[-1]
            print("id: "+str(id))
            # print("id"+id_matcher.group())
        else:
            print("Nothing Found")
        pattern = r'\S+(?:.jpg))'
        imagelist = re.findall('^http.jpg',redata)
        if imagelist:
            print(imagelist)
        else:
            print("Nothing Found")
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

