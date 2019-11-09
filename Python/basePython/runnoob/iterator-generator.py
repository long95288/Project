"""
迭代器的使用
"""
import sys
"""
数组的迭代器
"""
def sample1():
    list =[1,2,3,4]
    it = iter(list)
    # 使用for来迭代
    for x in it:
        print(x,end=" ")
    print("")

"""
next()迭代
"""
def sample2():
    list = [1,2,3,4]
    it = iter(list)
    while True:
        try:
            print(next(it))
        except StopIteration:
            sys.exit()
"""
创建迭代器
"""
class MyNumbers:
    def __iter__(self):
        self.a = 1
        return self

    def __next__(self):
        if self.a <= 20:
            x = self.a
            self.a += 1
            return x
        else:
            raise StopIteration

if __name__ == '__main__':
    # 使用迭代器
    sample1()
    print("===============")
    # sample2()

    myClass = MyNumbers()
    myClassIterator = iter(myClass)
    # print(next(myClassIterator))
    for i in range(10):
        print(next(myClassIterator))


