package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
    db *gorm.DB
)
type UserTb struct {
    Account string `gorm:"column:account;primary_key"`
    Password string `gorm:"column:password"`
}
func (u *UserTb) TableName() string  {
    // 返回的值是数据库中数据库表名
    return "usertb"
}
func init() {
    var err error
    args := "root:root1234@(localhost:3306)/gamedb?charset=utf8mb4&parseTime=True&loc=Local"
    db,err = gorm.Open("mysql",args)
    db.LogMode(true)
    if err != nil{
        fmt.Println("err:",err)
    }
}
func main() {
    //users := []UserTb{}
    //db.Find(&users)
    //db.Delete(&UserTb{})
    //fmt.Println(users)
    db.Create(&UserTb{
        Account:  "12233",
        Password: "pswddd",
    })
    u := UserTb{}
    db.First(&u)
    fmt.Println(u)
    u2 := UserTb{}
    db.Where("account = ?","122").Find(&u2)
    fmt.Println(u2)
    u2.Password = "pswff"
    db.Model(&u2).Update("password", u2.Password)
    db.Where("account = ?","122").Find(&u2)
    fmt.Println(u2)
    
}
