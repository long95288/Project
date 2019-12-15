
import sys
import socket

server_socket = socket.socket(
    socket.AF_INET,
    socket.SOCK_STREAM
)

host = socket.gethostname()

port = 9999

server_socket.bind((host,port))
server_socket.listen(5)

while True:
    print("监听阻塞中。。。。")
    clientsocket,addr = server_socket.accept()
    print("检测到连接。。。。。")
    print("连接地址:{}".format(addr))
    msg = "服务端返回数据\r\n"
    clientsocket.send(msg.encode('utf-8'))
    clientsocket.close()
