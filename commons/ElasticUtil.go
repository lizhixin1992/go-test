package commons

import (
	"context"
	"fmt"
	"github.com/lizhixin1992/test/conf"
	"github.com/olivere/elastic"
	"github.com/pelletier/go-toml"
)

var es = NewElastic()

var ctx = context.Background()

func NewElastic() *elastic.Client {
	tree := conf.GlobalConf.Get("elastic").(*toml.Tree)
	es, err := elastic.NewClient(elastic.SetURL(tree.Get("Urls").(string)))
	if err != nil {
		panic(err)
	}
	return es
}

//match查询
func MatchQuery(index, typ string, query elastic.Query, from, size int) *elastic.SearchHits {
	searchResult, err := es.Search().Index(index).Type(typ).Query(query).From(from).Size(size).Pretty(true).Do(ctx)
	if err != nil {
		panic(err)
	}

	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d \n", searchResult.Hits.TotalHits)
		return searchResult.Hits
	} else {
		fmt.Print("Found no total\n")
		return nil
	}
}
