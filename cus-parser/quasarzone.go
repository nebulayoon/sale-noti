package cusParser

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type QuasarzoneData struct {
	State string
	Title string
	Category string
	Price string
	Date string
}

func noSpaces(str string) (string){
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(str, "")
}

func Parsing(page int) ([]QuasarzoneData){
	// fmt.Println("Test function start")
	url := fmt.Sprintf("https://quasarzone.com/bbs/qb_saleinfo?page=%s", strconv.Itoa(page))

	res, err := http.Get(url)
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

	// fmt.Println("search test")
	var rows []QuasarzoneData

	doc.Find("div.market-info-list-cont").Each(func(idx int, s *goquery.Selection) {
		state := s.Find("span.label").Text()
		title := s.Find("span.ellipsis-with-reply-cnt").Text()
		category := s.Find("span.category").Text()
		price := s.Find("span.text-orange").Text()
		date := noSpaces(s.Find("span.date").Text())

		rows = append(rows, QuasarzoneData{State: state, Title: title, Category: category, Price: price, Date: date})
	})

	return rows
}