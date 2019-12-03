"""
连接数据库的工具类
"""
import pymysql
"""
将数据写入数据库
"""
def insertListToDataBase(list):
    db = pymysql.connect("localhost","root","root2037","blibli")
    affect_count = -1
    cursor = db.cursor()
    sql = "INSERT INTO t_image (id,src,status) VALUES (%s,%s,%s)"
    insertList = []
    for item in list:
        id = str(item).split("/")[-1]
        status = 0
        insertData = (id,item,'0')
        insertList.append(insertData)

    try:
        affect_count = cursor.executemany(sql, insertList)
        db.commit()
        cursor.close()
    except EnvironmentError:
        db.rollback()
        print("插入失败")
    db.close()
    return affect_count

def queryUndownloadImage():
    db = pymysql.connect("localhost","root","root2037","blibli")
    cursor = db.cursor()
    sql = "SELECT * FROM t_image WHERE status = 0"
    list = []
    try:
        cursor.execute(sql)
        # 获得记录
        results = cursor.fetchall()
        for row in results:
            src = row[1]
            list.append(src)
    except EnvironmentError:
        print("无法获取数据")
    # 关闭数据库
    db.close()
    return list

