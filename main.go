package main

import (
	"fmt"
	// cusParser "sale-noti/cus-parser"
	"sale-noti/database"
)

func main(){
	fmt.Println("main")
	// cusParser.Test()
	db, err := database.NewDBConnect()
	if err != nil {
		return
	}
	db.GenerateTable()
	
	
}