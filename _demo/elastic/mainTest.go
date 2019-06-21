package main

import (
	"fmt"
	"github.com/lizhixin1992/test/commons"
	"github.com/olivere/elastic"
)

func main() {
	query := elastic.NewMatchQuery("passage", "elk rocks")
	searchBuild := commons.SearchBuild{
		From:               0,
		Size:               10,
		Query:              query,
		Index:              "book1",
		Typ:                "english",
		FetchSourceContext: elastic.NewFetchSourceContext(true).Include("passage"),
	}

	result := commons.MatchQuery(searchBuild)

	for _, value := range result {
		fmt.Println(value)
	}
	ret := commons.SetResponseSuccessData(result)
	fmt.Println(ret)
	fmt.Println(commons.SetResponseSuccess())

}
