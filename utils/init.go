package utils

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func DBInit() {
	var err error
	DB, err = sql.Open("sqlite3", "./anouncement.db")
	if err != nil {
		panic("数据库连接失败, 请检查")
	}
	// 创建数据库
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS anouncement_info (
			id INTEGER PRIMARY KEY,
			platform TEXT,
			total Integer,
			lastTime INTEGER
		)
	`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckErr(err error) bool {
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return true
	}
	return false
}
