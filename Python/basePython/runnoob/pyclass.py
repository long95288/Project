# encode=utf-8

"""
定义父类people类
"""
class People:
    # 基本属性
    name = ''
    age = 0
    # 私有变量
    __weight = 0

    def __init__(self,n,a,w):
        self.name = n
        self.age = a
        self.__weight = w

    def speak(self):
        print("%s 说: 我 %d 岁。"%(self.name,self.age))
        print("我有%d斤"%self.__weight)


# 单继承
class Student(People):
    grade = ''
    def __init__(self,name,age,weight,grade):
        # 调用父类的构造函数
        People.__init__(self,name,age,weight)
        self.grade = grade
    # @Override
    def speak(self):
        print("%s 说：我 %d 岁了，我在读 %d 年级"%(self.name,self.age,self.grade))


class Speaker:
    topic = ''
    name = ''
    def __init__(self,name,topic):
        self.name = name
        self.topic = topic

    def speak(self):
        print("%s 的演讲主题是 %s" %(self.name, self.topic))


# 多重继承
class Sample(Student,Speaker):
    a = ''
    def __init__(self,name,age,weight,grade,topic):
        # 初始化
        Student.__init__(self,name,age,weight,grade)
        Speaker.__init__(self,name,topic)
        #

class Parent:
    def method(self):
        print("this is parent's method()")

class Child(Parent):
    def method(self):
        print("this is child's method()")

if __name__ == '__main__':
    p = People('world',10,70)
    p.speak()
    s = Student('hello',10,60,3)
    s.speak()
    sample = Sample("Tim",12,44,4,"hello bad world")
    sample.speak()
    # 当前继承的顺序是Speaker类、然后是Student类。同名方法调用Speaker中的speak方法
    # result: Tim 的演讲主题是 hello bad world
    # 更改继承顺序之后会相应的改变

    c = Child()
    c.method()
    # 子类调用父类同名方法 使用super(类，实例).同名方法
    super(Child, c).method()
