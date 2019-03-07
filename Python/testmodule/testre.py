# encoding=utf-8
import re

# 正则表达式的使用
# search 使用
a = 'one1two2three3'
inf0 = re.search('\d', a)  # 获得数字
print(inf0)
print(inf0.group())

# sub()函数使用
# 用于替换掉字符串的匹配项
phone = '123-4567-789'
# 替换
new_phone = re.sub('\D', '', phone)
print(new_phone)

# findall()函数，用于获得匹配之后的列表
b = 'one1two2three3'
new_b = re.findall('\d+', b)
print(new_b)

# 模块修饰符
# 使用r.S来跨行匹配
rS = '''<div>指数
</div>'''
word = re.findall('<div>(.* ?)</div>', rS, re.S)
print(word[0].strip())  # 使用strip()来消除换行符

test_data = '''
&nbsp;&nbsp;&nbsp;&nbsp;“斗之力，三段！”<br />
<br />
'''
print("测试文件：")
# 去掉&bpsp
new_data = re.sub('&nbsp;','',test_data,re.S)
print(new_data)

# 去掉<br>
new_data2 = re.sub('<br .*?>','',new_data,re.S)
print(new_data2)

# con_data = re.findall('&nbsp;&nbsp;&nbsp;&nbsp;(.*?)',test_data,re.S)
# print(con_data)
