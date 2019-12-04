"""
连接数据库的工具类
"""
import pymysql
from getBlibliImageByJson.Log import log
import json
"""
将数据写入数据库
"""
def insertListToDataBase(list,uid,dynamic_id):
    db = pymysql.connect("localhost","root","root2037","blibli")
    affect_count = -1
    cursor = db.cursor()
    sql = "INSERT INTO t_image (id,src,status,uid,dynamic_id) VALUES (%s,%s,%s,%s,%s)"
    insertList = []
    for item in list:
        id = str(item).split("/")[-1]
        status = 0
        insertData = (id,item,'0',uid,dynamic_id)
        insertList.append(insertData)

    try:
        affect_count = cursor.executemany(sql, insertList)
        db.commit()
        cursor.close()
    except EnvironmentError:
        log("执行sql:{}失败".format(sql))
        db.rollback()
    db.close()
    return affect_count
"""
废弃了
"""
def insertDynamic(uid,dynamic_id,image_list):
    db = pymysql.connect("localhost", "root", "root2037", "blibli")
    cursor = db.cursor()
    # 写入数据库分成两次写入
    # 1、写入动态数据
    insert_dynamic_sql = "INSERT INTO t_dynamic (dynamic_id,uid,download_status) VALUES ({},{},{})".format(dynamic_id, uid, '0')
    insert_image_sql = "INSERT INTO t_image (id,src,status,uid,dynamic_id) VALUES (%s,%s,%s,%s,%s)"
    try:
        log("执行sql:{}".format(insert_dynamic_sql))
        cursor.execute(insert_dynamic_sql)
        insertList = []
        for item in image_list:
            id = str(item).split("/")[-1]
            insertData = (id, item, '0', uid, dynamic_id)
            insertList.append(insertData)
        log("插入图片数据")
        cursor.executemany(insert_image_sql, insertList)
        # 提交数据
        db.commit()
        cursor.close()
    except EnvironmentError:
        log("执行sql:{}失败".format(insert_dynamic_sql))
        db.rollback()
    db.close()
    # 2、写入图片
    insertListToDataBase(list=image_list, uid=uid, dynamic_id=dynamic_id)

"""
获得一定量的未下载的动态数据
"""
def queryUndownloadDynamic(size):
    db = pymysql.connect("localhost","root","root2037","blibli")
    cursor = db.cursor()
    sql = "SELECT DISTINCT dynamic_id,uid FROM t_dynamic WHERE download_status = 0 LIMIT 0,{} ".format(size)
    results = []
    try:
        cursor.execute(sql)
        # 获得记录
        results = cursor.fetchall()
        cursor.close()
    except EnvironmentError:
        print("无法获取数据")
    # 关闭数据库
    db.close()

    # 返回结果
    return results

"""
根据uid和动态id获得相应的图片url列表
"""
def queryImagesByDidAndUid(uid,dynamic_id):

    db = pymysql.connect("localhost", "root", "root2037", "blibli")
    cursor = db.cursor()
    sql = "SELECT DISTINCT src FROM t_image WHERE uid ={} AND dynamic_id = {}".format(uid, dynamic_id)
    list =[]
    try:
        cursor.execute(sql)
        results = cursor.fetchall()
        for item in results:
            src = item[0]
            list.append(src)

        cursor.close()
    except EnvironmentError:
        log("执行sql:{}失败".format(sql))

    db.close()
    return list

"""
根据uid 和dyid 查询数据
"""
def queryRecordByUidAndDynamicId(uid,dynamic_id):
    db = pymysql.connect("localhost", "root", "root2037", "blibli")
    cursor = db.cursor()
    sql = "SELECT * FROM t_dynamic WHERE uid={} AND dynamic_id={}".format(uid, dynamic_id)
    result = None
    try:
        cursor.execute(sql)
        result = cursor.fetchone()
        cursor.close()
    except EnvironmentError:
        log("执行sql:{}异常".format(sql))

    db.close()
    return result

def getUID():
    db = pymysql.connect("localhost", "root", "root2037", "blibli")
    cursor = db.cursor()
    sql = "SELECT * FROM t_user"
    uids = []
    try:
        cursor.execute(sql)
        result = cursor.fetchall()
        cursor.close()
        for item in result:
            uid = item[0]
            uids.append(uid)
    except EnvironmentError:
        log("执行sql:{}异常".format(sql))
    db.close()

    return uids

"""
根据动态列表统一插入数据库
dynamic_list的结构
[
(
uid,
dynamic_id,
[
 'http:',,,,
]
)
]
"""
def insertDynamicList(dynamic_list):
    # 生成动态的数据
    insert_dynamic_list = []
    # 生成图片的数据
    insert_image_list =[]
    for dynamic in dynamic_list:
        uid = dynamic[0]
        dynamic_id = dynamic[1]
        image_list = dynamic[2]
        # 生成动态表的数据
        insert_dynamic_list.append((dynamic_id, uid, '0'))
        # 生成图片表的数据
        for image in image_list:
            image_id = str(image).split("/")[-1]
            # 图片数据
            insert_image_data = (image_id,image,'0',uid,dynamic_id)
            insert_image_list.append(insert_image_data)
    # 获得全部数据后分两次批量插入
    insert_dynamic_sql = "INSERT INTO t_dynamic (dynamic_id,uid,download_status) VALUES (%s,%s,%s)"
    insert_image_sql = "INSERT INTO t_image (id,src,status,uid,dynamic_id) VALUES (%s,%s,%s,%s,%s)"
    db = pymysql.connect("localhost", "root", "root2037", "blibli")
    cursor = db.cursor()
    try:
        # 批量插入动态
        if len(dynamic_list) > 0:
            log("插入动态数据......")
            cursor.executemany(insert_dynamic_sql,insert_dynamic_list)
            # 提交变更
            db.commit()
        # 批量插入图片
        if len(insert_image_list) > 0:
            log("插入动态图片数据.....")
            cursor.executemany(insert_image_sql, insert_image_list)
            # 提交变更
            db.commit()
    except EnvironmentError:
        db.rollback()
        print("插入错误")
    db.close()

"""
更新数据为已经下载的状态
"""
def updateDynamicToDownloaded(uid,dynamic_id):
    db = pymysql.connect("localhost", "root", "root2037", "blibli")
    cursor = db.cursor()
    sql = "UPDATE t_dynamic SET download_status = 1 WHERE uid = {} AND dynamic_id = {}".format(uid, dynamic_id)
    affect_count = 0
    try:
        affect_count = cursor.execute(sql)
        db.commit()
        cursor.close()
    except EnvironmentError:
        log("执行:sql:{} 失败".format(sql))
        db.rollback()
    db.close()

    return affect_count

