# encoding=utf-8
# coding=utf-8
import requests
import re
import time
from bs4 import BeautifulSoup
import random

header = {
		'User-Agent':'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3724.8 Safari/537.36'
}

f = open('ye.txt', 'a', encoding='utf-8')
# 先读后写
list = open('chapterlist.txt', 'r', encoding='utf-8')

# 这是剃掉杂毛&nbsp;<br />等
# 默认count=0为替换全部
def clear(str):
    temp1_data = re.sub(r'.*?&nbsp;', '', str, count=0)
    temp2_data = re.sub(r'<br .*?>', '', temp1_data, count=0)
    return temp2_data


test = open('test.html', 'r', encoding='utf-8')
test_data = test.read()

def get_info(url):
    res = requests.get(url, headers=header)
    if res.status_code == 200:
        next_chapter_url = None
        soup = BeautifulSoup(res.content.decode('utf-8'), 'html.parser')
        # 章节名称
        chapter_name = soup.select('#wrapper > div.content_read > div.box_con > div.bookname > h1')
        print("章节名称"+chapter_name[0].string)
        # 写入章节名称
        in_name = chapter_name[0].string.strip()
        # print(in_name)
        # print("章节名称:"+in_name)
        f.write(in_name + '\n')
        # 下一个章节的地址
        next_chapter_url = soup.select('#wrapper > div.content_read > div.box_con > div.bookname > '
                                       'div.bottem1 > a:nth-child(4)')
        print("下一章节地址："+next_chapter_url[0].get('href'))
        # 章节的内容
        tempcontents = re.findall('<div id="content">(.*?)</div>', res.content.decode('utf-8'),
                                  re.S)
        pure_content = clear(tempcontents[0])
        # 写入内容
        f.write(pure_content)
        print("写入成功")
        return str(next_chapter_url[0].get('href').strip())
    else:
        print("状态码:" + res.status_code)
        pass


# start_chapter_url = '/2_2447/973170.html'
# 读取文件获得最新需要下载章节地址

start_chapter_url = list.read().strip()
# start_chapter_url = '/2_2447/973248.html'
# 更改文件模式
list = open('chapterlist.txt', 'w', encoding='utf-8')

rooturl = 'http://www.022003.com'
chapter_num = 1
while True:
    if start_chapter_url == '/54_54136/':
        break
    # elif chapter_num == 3:
        #break
    else:
        page_url = rooturl + start_chapter_url
        print("开始爬取："+page_url)

        # start_chapter_url = '/2_2447/'
        start_chapter_url = get_info(page_url)
        # 日志跟踪
        list.write(start_chapter_url + '\n')
        if start_chapter_url == '/54_54136/':
            break
        page_url = ""
        # chapter_num += 1
        time.sleep(1+random.random())

'''
soup = BeautifulSoup(test_data,'html.parser')
chapter_name = soup.select('#wrapper > div.content_read > div.box_con > div.bookname > h1')
print(chapter_name[0].string)
# 下一个章节的地址
next_chapter_url = soup.select('#wrapper > div.content_read > div.box_con > div.bookname > div.bottem1 > a')
print(next_chapter_url[3].get('href'))
n = soup.select('#wrapper > div.content_read > div.box_con > div.bookname > div.bottem1 > a:nth-child(4)')
print(n[0].get('href'))
##wrapper > div.content_read > div.box_con > div.bookname > div.bottem1 > a:nth-child(4)
'''
