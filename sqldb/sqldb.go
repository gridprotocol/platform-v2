package sqldb

import (
	"database/sql"
	"log"
)

func init() {
	// 创建或打开数据库
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建provider表
	createTableSQL := `CREATE TABLE IF NOT EXISTS providers (
        "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, 
        "address" TEXT,
        "name" TEXT,
		"ip" TEXT,
		"port" TEXT,
		"domain" TEXT,
        "created_at" DATETIME
    );`

	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatal(err)
	}

	log.Println("provider表创建成功")

	// 创建nodes表
	createTableSQL = `CREATE TABLE IF NOT EXISTS nodes (
        "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, 
        "provider" TEXT,
        "cpu_price" TEXT,
		"cpu_core" TEXT,
		"gpu_price" TEXT,
		"gpu_model" TEXT,
        "mem_price" TEXT,
		"mem_num" TEXT,
		"disk_price" TEXT,
		"disk_num" TEXT
    );`

	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatal(err)
	}

	log.Println("nodes表创建成功")

	// 创建orders表
	createTableSQL = `CREATE TABLE IF NOT EXISTS orders (
        "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT, 
        "user" TEXT,
        "provider" TEXT,
		"node_id" TEXT,
		"app_name" TEXT,
		"value" TEXT,
        "create_time" DATETIME,
		"duration" TEXT,
		"status" TEXT
    );`

	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatal(err)
	}

	log.Println("orders表创建成功")
}
