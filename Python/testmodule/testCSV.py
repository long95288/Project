# encoding=utf-8

import csv
fp = open('test.csv', 'w+')
writer = csv.writer(fp)
writer.writerow(('id', 'name'))
writer.writerow(('1', 'xiaoming'))
writer.writerow(('2', '张三'))
writer.writerow(('3', '李四'))


