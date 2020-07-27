package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)
type UserTb struct {
    Account string `gorm:"column:account;primary_key"`
    Password string `gorm:"column:password"`
}
var (
    db *gorm.DB
)
func (u *UserTb) TableName() string  {
    return "usertb"
    
}

func InsertDemo() {
    insertUser := UserTb{
        Account:  "testUserAccount",
        Password: "testPassword",
    }
    // 判断主键是否为空
    re := db.NewRecord(insertUser)
    // 插入数据
    db.Create(&insertUser)
    fmt.Println("insert operation :",re)
}

func SelectDemo(){
    user := UserTb{}
    db.Where("account=?","testUserAccount").Find(&user)
    fmt.Println(user)
}
func UpdateDemo(){
    user := UserTb{Account: "testUserAccount"}
    db.Where("account=?","testUserAccount").First(&user)
    fmt.Println("find : ",user)
    user.Password = "testUserPasswordUpdate"
    // db.Save(&user)
    db.Model(&user).Update("password","testUserPasswordUpdate2")
    //
    fmt.Println("update")
    fmt.Println(user)
    //
    fmt.Println("find : ")
    db.Where("account","testUserAccount").First(&user)
    fmt.Println(user)
}
func DeleteDemo() {
    user := UserTb{Account: "testUserAccount"}
    db.Delete(&user)
    fmt.Println("Delete success")
}
func main() {
    var err error
    db,err = gorm.Open("mysql","root:root1234@(localhost:3306)/gamedb?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    usr := UserTb{}
    db.First(&usr)
    fmt.Println(usr)
    users := []UserTb{}
    db.Find(&users)
    fmt.Println(users)
    InsertDemo()
    SelectDemo()
    UpdateDemo()
    DeleteDemo()
}
