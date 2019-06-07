# encoding=utf-8

class Student():
    
    def __init__(self,name,score):
        self.name = name
        self.score = score
    
    def printScore(self):
        print('%s:%s',self.name,self.score)
    
    pass
