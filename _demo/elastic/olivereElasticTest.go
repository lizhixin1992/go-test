package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
)

func main() {
	ctx := context.Background()
	es, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}

	info, code, err := es.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	//get 通过_id查询
	resp, err := es.Get().Index("book1").Type("english").Id("2").Do(ctx)
	if err != nil {
		panic(err)
	}
	if resp.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", resp.Id, resp.Version, resp.Index, resp.Type)
	}

	//match查询
	query := elastic.NewMatchQuery("passage", "elk rocks")
	searchResult, err := es.Search().Index("book1").Type("english").Query(query).From(0).Size(10).Pretty(true).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	fmt.Printf("Found a total of %d book\n", searchResult.TotalHits())

	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d book\n", searchResult.Hits.TotalHits)

		for _, hit := range searchResult.Hits.Hits {
			var b book
			err := json.Unmarshal(*hit.Source, &b)
			if err != nil {
				// Deserialization failed
			}

			fmt.Printf("book is %s\n", b.Passage)
		}
	} else {
		fmt.Print("Found no book\n")
	}
}

type book struct {
	Passage string `json:"passage"`
}
