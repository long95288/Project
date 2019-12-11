
def this_fails():
    x = 1/0

if __name__ == '__main__':
    try:
        this_fails()
    except ZeroDivisionError as err:
        # 接收并打印异常信息
        print('handling run-time error',err)
