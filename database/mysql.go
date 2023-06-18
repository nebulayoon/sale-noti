package database

import (
	"database/sql"
	"fmt"
	"log"
	cusParser "sale-noti/cus-parser"

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
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
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
	query := fmt.Sprintf("SELECT * FROM %s;", table)
	_, tableCheck := db.connection.Query(query)

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

func (db DBMysqlRepository) ProductInsert(table string, site string, link string, data cusParser.QuasarzoneData) (sql.Result, error){
	tableColumns := fmt.Sprintf("%s(site, state, title, category, price, date, link)", table)
	result, err := db.connection.Exec("INSERT INTO " + tableColumns + "VALUES (?, ?, ?, ?, ?, ?, ?)", site, data.State, data.Title, data.Category, data.Price, data.Date, link)
	fmt.Println(data.Date)
	if err != nil {
		fmt.Println(result, err)
		return nil, err
	}

	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	
	return result, err
}