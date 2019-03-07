# encoding=utf-8

# 写文件打开
f = open('testfilesys.txt', 'w')
f.write("写进去的内容")
f.close()

# 读文件打开

f = open('testfilesys.txt','r')
content = f.read()
print("读取文件的内容:" + content)
f.close()

# 追加文件的写入
f = open('testfilesys.txt','a')
f.write("\n追加的内容")
f.close()

f = open('testfilesys.txt', 'w',encoding='utf-8')

f.write("写了新的东西，读出来了吗？")
f.close()

f = open('testfilesys.txt', 'r',encoding='utf-8')
print(f.read())

