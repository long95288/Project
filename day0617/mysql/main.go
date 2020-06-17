package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (error) {
	dsn := "user:password@(127.0.0.1:3306)/dbname?charset=utf8"
	var err error
	//db,err := sql.Open("mysql",dsn)会出现作用域覆盖导致空指针
	// 创建新的指针会把全局的db指针给覆盖掉
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	id   int64
	name string
	age  int
}

func InsertRowDemo() {
	sqlStr := "insert into user(name,age) value (?,?)"
	ret, err := db.Exec(sqlStr, "王五", 23)
	if err != nil {
		fmt.Printf("insert failed ,err :%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed ,err :%v \n", err)
		return
	}
	fmt.Printf("insert success,the id is %d.\n", theID)

}

func QueryRowDemo() {
	sqlStr := "select id,name,age from user where id = ?"
	var u User = User{}
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return
	}
	fmt.Printf("user:%v", u)
}
func QueryMultiDemo(){
	sqlStr := "select id,name,age from user where id > ?"
	rows,err := db.Query(sqlStr,1)
	if err != nil {
		fmt.Printf("query failed,err :%v\n",err)
		return
	}
	defer rows.Close()
	
	for rows.Next(){
		var u User
		err := rows.Scan(&u.id,&u.name,&u.age)
		if err != nil {
			fmt.Printf("Scan failed,err:\v\n",err)
			return
		}
		fmt.Printf("u:%v",u)
	}
}
func UpdateRowDemo(){
	sqlStr := "update user set age=? where id = ?"
	ret,err := db.Exec(sqlStr,33,2)
	if err != nil {
		fmt.Printf("update failed,err : %v\n",err)
		return
	}
	// 更新影响的行数
	n,err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowAffected failed,err:%v\n",err)
		return
	}
	fmt.Printf("update success,affected rows : %d\n",n)
}
func DeleteRowDemo(){
	sqlStr := "Delete FROM user where id = ?"
	ret,err := db.Exec(sqlStr,3)
	if err != nil {
		fmt.Printf("Delete failed ,err :%v\n",err)
		return
	}
	n,err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowAffected failed ,err :%v\n",err)
		return
	}
	fmt.Printf("delete success ,affect rows : %d\n",n)
}

func PrepareQueryDemo()  {
	sqlStr := "Select id,name,age from user where id > ?"
	stmt,err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("Prepare failed,err :%v \n",err)
		return
	}
	defer stmt.Close()
	// 放置参数
	rows,err := stmt.Query(0)
	if err != nil {
		fmt.Printf("prepare failed,err :%v\n",err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u User
		err := rows.Scan(&u.id,&u.name,&u.age)
		if err != nil {
			fmt.Printf("Scan failed ,err: %v\n",err)
			return
		}
		fmt.Printf("user: %v\n",u)
	}
	
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err %v \n", err)
		return
	}
	InsertRowDemo()
	QueryRowDemo()
	QueryMultiDemo()
	UpdateRowDemo()
	QueryMultiDemo()
	DeleteRowDemo()
	QueryMultiDemo()
	PrepareQueryDemo()
	
}
