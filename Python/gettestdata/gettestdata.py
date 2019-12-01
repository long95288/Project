# encoding=utf-8
import requests

url = 'https://space.bilibili.com/36802028/dynamic'
header = {
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko)'
                  ' Chrome/71.0.3578.98 Safari/537.36'
}
res = requests.get(url, headers=header, timeout=30)
if res.status_code == 200:
    f = open('testDynamic.html', 'w', encoding='utf-8')
    txt = res.content.decode('utf-8')

    # print("测试："+txt)
    f.write(txt)
    f.close()
    print("获得测试文件成功")
else:
    print("出错了")

