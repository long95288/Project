
if __name__ == '__main__':
    dict = {
        'name': 'Runoob',
        'Age':7,
        'Class':'First'
    }

    print(dict)
    print(dict['name'])
    # 添加新的元组
    dict['school'] = '菜鸟教程'
    print(dict['school'])
    # 修改其中的元组
    dict['Age'] = 9
    print(dict['Age'])
    # 删除其中的一个元组
    del dict['name']
    print(dict)
    # 清空元组
    print("清空元组")
    dict.clear()
    print(dict)
    print("========")
    print("删除字典")
    del dict
    try:
        print(dict['name'])
    except Exception as e:
        print("捕获异常:",e)
    # 字典的键的特性相同键名只保留后面一个的内容
    dict  = {'Name': 'Runoob','Age':7,'Name':'后一个名字'}
    print(dict['Name'])


