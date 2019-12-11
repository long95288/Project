

def exception_1():
    x = 10
    if x > 5:
        # 抛出异常
        raise Exception('x 不能大于5.x的值为:{}'.format(x))



if __name__ == '__main__':
    try:
        exception_1()
    except Exception as e:
        print("异常被接收，",e)

