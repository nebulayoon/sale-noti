package cusParser

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type originData struct {
	state string
	title string
	category string
	price string
	date string
}

type QuasarzoneParser struct {
	originData
}

func Test(){
	fmt.Println("Test function start")
	res, err := http.Get("https://quasarzone.com/bbs/qb_saleinfo?page=1")
	if err != nil {
		panic(err)
	}	

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 		panic(err)
	// }
  // fmt.Printf("%s\n", string(data))

	// f, err := os.Create("test.html")
	// defer f.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }

	fmt.Println("search test")
	doc.Find("div.market-info-list-cont").Each(func(idx int, s *goquery.Selection) {
		state := s.Find("span.label").Text()
		title := s.Find("span.ellipsis-with-reply-cnt").Text()
		category := s.Find("span.category").Text()
		price := s.Find("span.text-orange").Text()
		date := s.Find("span.date").Text()

		fmt.Printf("%s %s %s %s %s\n", state, title, category, price, date)
	})
	
}