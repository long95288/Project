# -*- coding: utf-8 -*-

# Form implementation generated from reading ui file 'MainWind.ui'
#
# Created by: PyQt5 UI code generator 5.13.0
#
# WARNING! All changes made in this file will be lost!


from PyQt5 import QtCore, QtGui, QtWidgets


class Ui_Form(object):
    def setupUi(self, Form):
        Form.setObjectName("Form")
        Form.resize(379, 250)
        self.label = QtWidgets.QLabel(Form)
        self.label.setGeometry(QtCore.QRect(10, 20, 51, 31))
        font = QtGui.QFont()
        font.setFamily("Adobe Arabic")
        font.setPointSize(14)
        self.label.setFont(font)
        self.label.setObjectName("label")
        self.url_text_line_edit = QtWidgets.QLineEdit(Form)
        self.url_text_line_edit.setGeometry(QtCore.QRect(50, 20, 231, 31))
        self.url_text_line_edit.setObjectName("url_text_line_edit")
        self.analyze_btn = QtWidgets.QPushButton(Form)
        self.analyze_btn.setGeometry(QtCore.QRect(290, 20, 81, 31))
        self.analyze_btn.setObjectName("analyze_btn")
        self.progressBar = QtWidgets.QProgressBar(Form)
        self.progressBar.setGeometry(QtCore.QRect(50, 220, 321, 23))
        self.progressBar.setProperty("value", 0)
        self.progressBar.setObjectName("progressBar")
        self.label_2 = QtWidgets.QLabel(Form)
        self.label_2.setGeometry(QtCore.QRect(10, 220, 41, 21))
        font = QtGui.QFont()
        font.setFamily("Adobe 楷体 Std R")
        font.setPointSize(12)
        self.label_2.setFont(font)
        self.label_2.setObjectName("label_2")
        self.download_btn = QtWidgets.QPushButton(Form)
        self.download_btn.setGeometry(QtCore.QRect(10, 180, 361, 31))
        self.download_btn.setObjectName("download_btn")
        self.label_3 = QtWidgets.QLabel(Form)
        self.label_3.setGeometry(QtCore.QRect(13, 80, 41, 21))
        self.label_3.setObjectName("label_3")
        self.novel_name_edit = QtWidgets.QLineEdit(Form)
        self.novel_name_edit.setGeometry(QtCore.QRect(70, 79, 221, 21))
        self.novel_name_edit.setObjectName("novel_name_edit")
        self.label_4 = QtWidgets.QLabel(Form)
        self.label_4.setGeometry(QtCore.QRect(10, 116, 51, 20))
        self.label_4.setObjectName("label_4")
        self.chapter_count = QtWidgets.QLabel(Form)
        self.chapter_count.setGeometry(QtCore.QRect(70, 117, 201, 20))
        self.chapter_count.setObjectName("chapter_count")
        self.label_5 = QtWidgets.QLabel(Form)
        self.label_5.setGeometry(QtCore.QRect(20, 140, 31, 21))
        self.label_5.setObjectName("label_5")
        self.status_label = QtWidgets.QLabel(Form)
        self.status_label.setGeometry(QtCore.QRect(70, 140, 201, 21))
        self.status_label.setObjectName("status_label")

        self.retranslateUi(Form)
        QtCore.QMetaObject.connectSlotsByName(Form)

    def retranslateUi(self, Form):
        _translate = QtCore.QCoreApplication.translate
        Form.setWindowTitle(_translate("Form", "Form"))
        self.label.setText(_translate("Form", "URL："))
        self.analyze_btn.setText(_translate("Form", "分析"))
        self.label_2.setText(_translate("Form", "进度:"))
        self.download_btn.setText(_translate("Form", "开始下载"))
        self.download_btn.setEnabled(False)
        self.label_3.setText(_translate("Form", "小说名:"))
        self.label_4.setText(_translate("Form", "章节数:"))
        self.chapter_count.setText(_translate("Form", "未知"))
        self.label_5.setText(_translate("Form", "状态:"))
        self.status_label.setText(_translate("Form", "未知"))
