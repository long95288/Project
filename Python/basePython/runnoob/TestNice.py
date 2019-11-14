
from basePython.runnoob.Nice import Nice

def callbackFunction(value1,value2):
    print("this is callBack function ")
    print("value1= "+value1)
    print("value2= "+value2)
    print("callbackFunction return ....")

if __name__ == '__main__':
    n = Nice(value1="dd",value2="33",callback1_123456_abcdefg=callbackFunction)
    n.print()


