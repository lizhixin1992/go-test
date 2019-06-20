package main

import (
	"encoding/json"
	"fmt"
	"github.com/lizhixin1992/test/commons"
	"github.com/olivere/elastic"
)

type book1 struct {
	Passage string `json:"passage"`
}

func main() {
	query := elastic.NewMatchQuery("passage", "elk rocks")
	result := commons.MatchQuery("book1", "english", query, 0, 10)

	for _, hit := range result.Hits {
		var b book1
		err := json.Unmarshal(*hit.Source, &b)
		if err != nil {
			// Deserialization failed
		}

		fmt.Printf("book is %s\n", b.Passage)
	}
}
