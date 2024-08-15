package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(192.168.1.21:3306)/glkj-hb")
	if err != nil {
		fmt.Println("Error connect to database! cause: ", err)
		return
	}
	defer db.Close()
	db.SetMaxOpenConns(5)
	// 测试连接
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connect to database! cause: ", err)
		return
	}
	fmt.Println("Successfully connect to MySQL!")
	rows, errQuery := db.Query("select id, account, employee_name name from fbp_employee t where t.enable_flag = 1 limit 5")
	if errQuery != nil {
		fmt.Println("one error happened when query! cause: ", errQuery)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id, account, name string
		err = rows.Scan(&id, &account, &name)
		if err != nil {
			fmt.Println("error Scanning row: ", err)
			continue
		}
		fmt.Printf("id: %s, name: %s, account: %s\n", id, name, account)
	}
	// 检查遍历过程
	if err = rows.Err(); err != nil {
		fmt.Println("error iterating through rows: ", err)
	}

	fmt.Println("end.")

	// 开启服务端
	server := NewServer("127.0.0.1", 4004)
	server.Start()

}
