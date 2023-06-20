package database

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	cusParser "sale-noti/cus-parser"
	"strings"
	"time"

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

type Product struct {
	Id int
	Site string
	State string
	Title string
	Category string
	Price string
	Date string
	Link string
	CreatedAt string
}

type DBMysqlRepository struct {
	connection *sql.DB
}

func scanStructFields(v interface{}) ([]interface{}) {
	value := reflect.ValueOf(v).Elem()
	numFields := value.NumField()
	
	fields := make([]interface{}, numFields)
	for i := 0; i < numFields; i++ {
		fields[i] = value.Field(i).Addr().Interface()
	}
	
	return fields
}

func connect() (*DBMysqlRepository, error) {
	db := DBMysqlRepository{}
	connection, err := sql.Open("mysql", "noti:1234@tcp(127.0.0.1:3306)/sale_noti")
	if err != nil{
		log.Fatal(err)
		return nil, err
	}

	db.connection = connection
	return &db, nil
}

func NewDBConnect() (*DBMysqlRepository, error) {
	db, err := connect()

	err = db.connection.Ping()
	if err != nil{
		for retries := 0; retries < 3; retries++ {
			db, err = connect()

			if err != nil {
				fmt.Println("재연결 시도 중 오류:", err)
			} else {
				err = db.connection.Ping()
				if err == nil {
					return db, nil
				}
			}

			time.Sleep(time.Second * 3)
		}

		return nil, err
	}
	
	return db, nil
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

func (db DBMysqlRepository) createTable() (sql.Result, error) {
	result, err := db.connection.Exec(productTableQuery)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db DBMysqlRepository) GenerateTable() (sql.Result, error) {
	hasTable := db.checkTable("products")
	if hasTable == true {
		return nil, nil
	}
	return db.createTable()
}

func (db DBMysqlRepository) DBClose() (error) {
	return db.connection.Close()
}

func (db DBMysqlRepository) FindOne(table string, fields map[string]interface{}, resultType interface{}) (interface{}, error) {
	var whereClauses []string
	var values []interface{}

	for field, value := range fields {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", field))
		values = append(values, value)
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", table, strings.Join(whereClauses, " AND "))
	result := db.connection.QueryRow(query, values...)
	
	convertedResult := reflect.New(reflect.TypeOf(resultType)).Interface()
	err := result.Scan(scanStructFields(convertedResult)...)
	if err != nil {
		return nil, err
	}

	return convertedResult, nil
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