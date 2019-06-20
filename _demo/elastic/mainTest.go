package main

import (
	"fmt"
	"github.com/lizhixin1992/test/commons"
	"github.com/olivere/elastic"
)

func main() {
	query := elastic.NewMatchQuery("passage", "elk rocks")
	result := commons.MatchQuery("book1", "english", query, 0, 10)

	for _, value := range result {
		fmt.Println(value)
	}
	ret := commons.SetResponseSuccessData(result)
	fmt.Println(ret)
	fmt.Println(commons.SetResponseSuccess())

}
