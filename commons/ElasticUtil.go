package commons

import (
	"context"
	"encoding/json"
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

//
func setReturnValue(result *elastic.SearchHits) (list []interface{}) {
	if result != nil {
		list = make([]interface{}, len(result.Hits))
		len := 0
		for _, hit := range result.Hits {
			var b interface{}
			err := json.Unmarshal(*hit.Source, &b)
			if err != nil {
				// Deserialization failed
			} else {
				list[len] = b
				len++
			}
		}
	}
	return list
}

//match查询
func MatchQuery(index, typ string, query elastic.Query, from, size int) (list []interface{}) {
	searchResult, err := es.Search().Index(index).Type(typ).Query(query).From(from).Size(size).Pretty(true).Do(ctx)
	if err != nil {
		panic(err)
	}

	if searchResult.Hits.TotalHits > 0 {
		return setReturnValue(searchResult.Hits)
	} else {
		fmt.Print("Found no total\n")
		return nil
	}
}
