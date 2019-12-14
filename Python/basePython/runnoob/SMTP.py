"""
邮件相关
"""
import smtplib
from email.mime.text import MIMEText
from email.header import Header
if __name__ == '__main__':
    sender = 'long@long.com'
    receivers = ['1186481819@qq.com']
    message = MIMEText('测试python发送邮件','plain','utf-8')
    message['From'] = Header('LONG','utf-8')
    message['To'] = Header('测试接收者', 'utf-8')

    subject = 'Python SMTP 测试'
    message['Subject'] = Header(subject,'utf-8')
    try:
        smtpObj = smtplib.SMTP('localhost')
        smtpObj.sendmail(sender,receivers,message.as_string())
        print("发送成功")
    except smtplib.SMTPException:
        print("发送失败")
    pass