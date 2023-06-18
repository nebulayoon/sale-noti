package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const productTableQuery = `
	CREATE TABLE products (
		id int NOT NULL AUTO_INCREMENT,
		site VARCHAR(50),
		state VARCHAR(20),
		title VARCHAR(200),
		category VARCHAR(100),
		price VARCHAR(100),
		date VARCHAR(100),
		link VARCHAR(1024),
		PRIMARY KEY(id)
	);
`

type DBMysqlRepository struct {
	connection *sql.DB
}

func NewDBConnect() (*DBMysqlRepository, error) {
	db := DBMysqlRepository{}
	connection, err := sql.Open("mysql", "noti:1234@tcp(127.0.0.1:3306)/sale_noti")
	if err != nil{
		log.Fatal(err)
		return nil, err
	}

	err = connection.Ping()
	if err != nil{
		return nil, err
	}
	
	db.connection = connection
	return &db, nil;
}

func (db DBMysqlRepository) checkTable(table string) (bool) {
	_, tableCheck := db.connection.Query("SELECT * FROM ?", table)
	if tableCheck == nil {
		return true
	} else {
		return false
	}
}

func (db DBMysqlRepository) createTable() () {
	// "CREATE TABLE test1 (id INT, name VARCHAR(50));"
	result, err := db.connection.Exec(productTableQuery)
	if err != nil {
		fmt.Println("쿼리 실행 에러:", err)
		return
	}

	fmt.Println(result)
}

func (db DBMysqlRepository) GenerateTable() {
	hasTable := db.checkTable("products")
	if hasTable == false {
		db.createTable()
	}
}

func (db DBMysqlRepository) DBClose() (error) {
	return db.connection.Close()
}

func (db DBMysqlRepository) Insert() (sql.Result, error){
	result, err := db.connection.Exec("INSERT INTO test1 VALUES (?, ?)", 11, "Jack")
	if err != nil {
		return nil, err
	}

	n, err := result.RowsAffected()
    if n == 1 {
      fmt.Println("1 row inserted.")
    }
	
	return result, err
}

