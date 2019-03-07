# encoding=utf-8

import xlwt

book = xlwt.Workbook(encoding='utf-8')  # 创建工作簿
sheet = book.add_sheet('Sheet1')  # 创建工作表
sheet.write(0, 0, 'python')
sheet.write(1, 1, 'love')
book.save('testxlwt.xls')
