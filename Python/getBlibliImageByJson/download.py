
import requests
import json
import time
import random

header = {
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) '
                  'Chrome/71.0.3578.98 Safari/537.36'
}
save_path = "D:\\biblivideo\\"

def save_image(image_name,image_conten):
    filename = save_path + image_name
    with open(filename, 'wb') as fp:
        fp.write(image_conten)
        print("保存:"+image_name+"成功")

def get_image_content(url):
    response = requests.get(url, header)
    response_data = None
    if response.status_code == 200:
        response_data = response.content
    return response_data

if __name__ == '__main__':
    url = "https://i0.hdslb.com/bfs/album/106b690e589026e4ef6ed4815754a2f7e395f966.png"
    with open("list.json",'r') as f:
        list = json.load(f)
    print(list)
    # list = [
    #     "https://i0.hdslb.com/bfs/album/00aae0928c341499f5d5b69a0fdca601d6d8f2d6.jpg@104w_104h_1e_1c.jpg",
    #     "https://i0.hdslb.com/bfs/album/6eddc9843505efc0135b02e8b8975e5f4303196f.jpg@214w_214h_1e_1c.jpg"
    # ]
    timeout = 2
    for item in list:
        origin_url = str(item).strip()
        high_pix_url = str(origin_url).split("@")[0]
        # 图片文件名
        image_name = high_pix_url.split("/")[-1]
        url = high_pix_url
        print(url)
        content = get_image_content(url)
        if content is None:
            url = origin_url
            content = get_image_content(url)
            if content is None:
                print("获取图片失败")
            else:
                save_image(image_name=image_name, image_conten=content)
        else:
            save_image(image_name=image_name, image_conten=content)
        time.sleep(timeout + random.random())
