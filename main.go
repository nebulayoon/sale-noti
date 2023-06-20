package main

import (
	"fmt"
	"log"
	cusParser "sale-noti/cus-parser"
	"sale-noti/database"
	"sync"
)

const page = 5

func main() {
	// DB setting
	db, err := database.NewDBConnect()
	if err != nil {
		return
	}
	db.GenerateTable()

	var wait sync.WaitGroup
	wait.Add(page)
	
	for i := 1; i <= page; i++ {   
		go func(page int) {
			defer wait.Done()
			quasarzoneDatas := cusParser.Parsing(page)
			fmt.Println(quasarzoneDatas)
			for _, qd := range quasarzoneDatas {
				_, err := db.ProductInsert("products", "quasarzone", "", qd)
				if err != nil {
					log.Printf("insert failed %s", qd.Title)
				}
			}
		} (i)
	}

	wait.Wait()
	db.DBClose()

	// result, _ := db.FindOne("products", map[string]interface{}{ "title": "[G마켓] 5600g본체  335,120원 (우리카드한정)", "date": "06-15" }, database.Product{})
	// fmt.Println(result)
	// test := &database.Product{}
	// if err := mapstructure.Decode(result, &test); err != nil {
	// 		fmt.Println(err)
	// }
	// fmt.Println(test.Price)
}