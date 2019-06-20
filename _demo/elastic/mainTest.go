package main

import (
	"fmt"
	"github.com/lizhixin1992/test/commons"
	"github.com/olivere/elastic"
)

type book1 struct {
	Passage string `json:"passage"`
}

type returnVO struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	data    []interface{} `json:"data"`
}

func main() {
	query := elastic.NewMatchQuery("passage", "elk rocks")
	result := commons.MatchQuery("book1", "english", query, 0, 10)

	for _, value := range result {
		fmt.Println(value)
	}
	ret := returnVO{Code: 0, Message: "success", data: result}
	fmt.Println(ret)

}
