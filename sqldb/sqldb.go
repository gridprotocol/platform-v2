package sqldb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // 导入SQLite驱动
)

func init() {
	fmt.Println("create grid db")

	// create db
	db, err := sql.Open("sqlite3", "./grid_db/grid.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create providers table
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

	log.Println("providers table created")

	// create nodes table
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

	log.Println("nodes table created")

	// create orders table
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

	log.Println("orders table created")
}
