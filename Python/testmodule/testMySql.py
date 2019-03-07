# encoding=utf-8
import pymysql

# 连接数据库
conn = pymysql.connect(host='localhost', user='root',
                       passwd='123456', db='resource')
cursor = conn.cursor()  # 光标对象
# 执行查询语句
cursor.execute("select * from tbuploadfile")

# 获取全部数据
data = cursor.fetchall()
print(data)
# 插入新数据
insertdata="newdata.jpg"
cursor.execute("insert into tbuploadfile(filename) values ('newdata')")
conn.commit()  # 提交数据


cursor.execute("select * from tbuploadfile")
data = cursor.fetchall()
# 打印对
print(data)
