
try:
    fh = open("testfile","w")
    fh.write("这是一个测试文件")
except IOError:
    print("文件错误")
else:
    print("写入文件成功")
    fh.close()
