package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/xuri/excelize/v2"
)

type Word struct {
	En string `json:"en"`
	Ru string `json:"ru"`
}

var wordCollection = []Word{}
var cnt int

func main() {
	scrapPage("https://www.en365.ru/top1000.htm")
	scrapPage("https://www.en365.ru/top1000a.htm")
	scrapPage("https://www.en365.ru/top1000b.htm")
	fmt.Printf("cnt %v \n", cnt)

	writeResultXls()
}
func scrapPage(url string) {
	c := colly.NewCollector()

	c.OnHTML("table > tbody > tr ", func(e *colly.HTMLElement) {
		enWord := e.DOM.Find("td:nth-child(2)").Text()
		ruWord := e.DOM.Find("td:nth-child(3)").Text()

		if !strings.Contains(ruWord, "Перевод на русский") && ruWord != "" {
			wordCollection = append(wordCollection, Word{enWord, ruWord})
			cnt++

		}
	})
	c.Visit(url)

}
func writeResultXls() {
	xlsx := excelize.NewFile()

	xlsx.NewSheet("Sheet1")

	for i, word := range wordCollection {
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("A%v", i+1), word.En)
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("B%v", i+1), word.Ru)

	}
	err := xlsx.SaveAs(".RuEn.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
