package main

import (
	"fmt"
	"log"
	cusParser "sale-noti/cus-parser"
	"sale-noti/database"
	"sync"
)

const page = 5

func main(){
	// DB setting
	db, err := database.NewDBConnect()
	if err != nil {
		return
	}
	db.GenerateTable()

	var wait sync.WaitGroup
	wait.Add(page)
	
	for i := 1; i <= page; i++ {
		go func(page int){
			defer wait.Done()
			quasarzoneDatas := cusParser.Parsing(page)
			fmt.Println(quasarzoneDatas)
			for _, qd := range quasarzoneDatas {
				_, err := db.ProductInsert("products", "quasarzone", "", qd)
				if err != nil{
					log.Printf("insert failed %s", qd.Title)
				}
			}
		}(i)
	}

	wait.Wait()
}