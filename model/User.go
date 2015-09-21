package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	// "log"
	"time"
)

var db *sql.DB

const (
	DB_USER     = "shen"
	DB_PASSWORD = "gaotang"
	DB_NAME     = "demo"
)

func conn() {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s  dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err = sql.Open("postgres", dbinfo)
	checkErr(err)
	// defer db.Close()
}
func GetUserAll() []map[string]interface{} {
	// log
	conn()
	defer db.Close()
	// var LastUid int
	var err error
	var m []map[string]interface{}
	// err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid;", "astaxie", "研发部门", "2012-12-09").Scan(&LastUid)
	// checkErr(err)

	// fmt.Println("last inserted id =", LastUid)
	rows, err := db.Query("SELECT * FROM userinfo ")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		rowInfo := map[string]interface{}{
			"uid":        uid,
			"username":   username,
			"department": department,
			"created":    created,
		}
		m = append(m, rowInfo)
		// fmt.Println("uid | username | department | created ")
		// fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	}
	return m

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
