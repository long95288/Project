# encoding=utf-8

import xlrd
import xlwt

tiku = xlrd.open_workbook('tiku.xlsx')
sheet1 = tiku.sheet_by_index(0)

# print(sheet1.cell(0,0).value)
f = open('tiku.txt','a',encoding='utf-8')

# nrows = sheet1.nrows
for i in range(sheet1.nrows):
    # answers
    num = str(i)
    question = sheet1.row_values(i)[0] # 获得问题
    answer = sheet1.row_values(i)[1] # 获得答案
    option =num + "." + question+"答案:["+answer+"]"

    if("A" in answer):
        option =option + "A." +str(sheet1.row_values(i)[2])
    if("B" in answer):
        option =option + "B." + str(sheet1.row_values(i)[3])
    if("C" in answer):
        option =option +"C." + str(sheet1.row_values(i)[4])
    if("D" in answer):
        option =option +"D."+ str(sheet1.row_values(i)[5])
    if("E" in answer):
        option =option +"E."+ str(sheet1.row_values(i)[6])
    # if("对" in answer):
    #     print("对")
    # if("错" in answer):
    #     print("错")
    f.write(option+"\n")
    print(option)
    