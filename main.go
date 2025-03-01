package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库连接信息
	dsn := "root:pi@tcp(192.168.31.159:3306)/"

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// 创建数据库的 SQL 语句
	createDbSQL := `CREATE DATABASE IF NOT EXISTS english;`

	// 执行创建数据库的操作
	_, err = db.Exec(createDbSQL)
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	log.Println("Database created successfully!")
}
