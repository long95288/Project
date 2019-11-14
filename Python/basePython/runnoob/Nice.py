
class Nice():
    def __init__(self,value1,value2,callback1_123456_abcdefg):
        self.value1 = value1
        self.value2 = value2
        self.callback1 = callback1_123456_abcdefg

    def print(self):
        print("value1:"+self.value1)
        print("value2:"+self.value2)
        print("callback ===")
        self.callback1(self.value1,self.value2)
        print("print return ====")



