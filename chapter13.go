package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	checkConnMySql()
}

func checkConnMySql() {
	db, err := sql.Open("mysql", "user1:1234@tcp(localhost:3306)/mysqldb?charset=utf8")
	if err != nil {
		panic(err)
	}
	fmt.Println("DB 已建立")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("資料庫連線成功")

	// Establish Table Shema
	// DBCreate := `
	// CREATE TABLE employee
	// (
	// 	id INT NOT NULL UNIQUE,
	// 	name VARCHAR(20)
	// );`

	// _, err = db.Exec(DBCreate)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("表格 employee 已建立")

	// INSERT DATA
	// insertStmt, err := db.Prepare("INSERT INTO employee (id, name) VALUES (?, ?);")
	// if err != nil {
	// 	panic(err)
	// }
	// defer insertStmt.Close()
	// _, err = insertStmt.Exec(305, "Sue")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("成功插入資料 305, Sue")

	// QUERY DATA
	rows, err := db.Query("SELECT id,name FROM employee")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println("資料表查詢成功, 列出 employee 內容...")

	for rows.Next() {
		e := employee{}
		err := rows.Scan(&e.id, &e.name)
		if err != nil {
			panic(err)
		}
		fmt.Println(e.id, e.name)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	// UPDATE DATA
	updateStmt, err := db.Prepare("UPDATE employee SET name=? WHERE id=?;")
	if err != nil {
		panic(err)
	}
	defer updateStmt.Close()
	updatedResult, err := updateStmt.Exec("Clark", 305)
	if err != nil {
		panic(err)
	}
	updatedRecords, err := updatedResult.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("成功更新資料 , 筆數", updatedRecords)

	defer db.Close()
}

// DROP USER 'user1'@'localhost';
// CREATE USER 'user1'@'172.22.0.1' IDENTIFIED BY '1234';
// GRANT ALL PRIVILEGES ON *.* TO 'user1'@'172.22.0.1' WITH GRANT OPTION;
// FLUSH PRIVILEGES;

// 產生參數化查詢敘述(prepared statement)，將SQL指令編譯成位元組碼，然後再透過參數將值放進需要的地方，避免了SQL Injection
// 操作資料庫

type employee struct {
	id   int
	name string
}
