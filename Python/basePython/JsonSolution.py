"""
json文件的添加和修改
"""
import json

if __name__ == '__main__':
    with open('test.json','r') as load_f:
        content = json.load(load_f)
        # print(content)
        print(content["list"])
