# encoding=utf-8
from bs4 import BeautifulSoup

# 需要处理的文档
html_doc = """
<html><head><title>The Dormouse's story</title></head>
<body>
<p class="title"><b>The Dormouse's story</b></p>

<p class="story">Once upon a time there were three little sisters; and their names were
<a href="http://example.com/elsie" class="sister" id="link1">Elsie</a>,
<a href="http://example.com/lacie" class="sister" id="link2">Lacie</a> and
<a href="http://example.com/tillie" class="sister" id="link3">Tillie</a>;
and they lived at the bottom of a well.</p>

<p class="story">...</p>
"""

soup = BeautifulSoup(html_doc, 'html.parser')
# print(soup.prettify())

# 获得数据的几种方法
# 1.获得title标签字
print(soup.title)

# 2.
print("title'name:"+soup.title.name)

# 3.获得标签里面的内容
print(soup.title.string)

# 4.获得父标签的名字
print(soup.title.parent.name)

#
print(soup.p)
print(soup.p['class'])
print(soup.find_all('p'))

print(soup.find_all('a'))
print("单例超链接："+soup.a.get('href'))

print(soup.find(id="link3"))

print("分解出相应的数据")
for link in soup.find_all('a'):
    # 获得标签为a 的超链接的内容
    print(link.get('href'))

# 获得所有的文字
print(soup.get_text())
