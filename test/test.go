package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	// "log"
	"time"
)

const (
	DB_USER     = "shen"
	DB_PASSWORD = "gaotang"
	DB_NAME     = "demo"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s  dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	var LastUid int
	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid;", "astaxie", "研发部门", "2012-12-09").Scan(&LastUid)
	checkErr(err)

	fmt.Println("last inserted id =", LastUid)
	rows, err := db.Query("SELECT * FROM userinfo ")
	for rows.Next() {
		fmt.Println(122)
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid | username | department | created ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	}

	// for rows.Next() {
	// 	var uid int
	// 	var username string
	// 	var department string
	// 	var created time.Time
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Println("uid | username | department | created ")
	// 	fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	// }
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
