package commons

import (
	"context"
	"encoding/json"
	"github.com/lizhixin1992/test/conf"
	"github.com/olivere/elastic"
	"github.com/pelletier/go-toml"
)

var es = NewElastic()

var ctx = context.Background()

type SearchBuild struct {
	Index              string                      `json:"index"`
	Typ                string                      `json:"typ"`
	Query              elastic.Query               `json:"query"`
	From               int                         `json:"from"`
	Size               int                         `json:"size"`
	FetchSourceContext *elastic.FetchSourceContext `json:"fetch_source_context"`
	SortInfo           *elastic.SortInfo           `json:"sort_info"`
	AggregationName    string                      `json:"aggregation_name"`
	Aggregation        elastic.Aggregation         `json:"aggregation"`
	Highlight          *elastic.Highlight          `json:"highlight"`
}

func NewElastic() *elastic.Client {
	tree := conf.GlobalConf.Get("elastic").(*toml.Tree)
	es, err := elastic.NewClient(elastic.SetURL(tree.Get("Urls").(string)))
	if err != nil {
		panic(err)
	}
	return es
}

//设置查询相关
func setSearchService(build SearchBuild) *elastic.SearchService {
	searchService := es.Search()
	if build.Index != "" {
		searchService.Index(build.Index)
	}
	if build.Typ != "" {
		searchService.Type(build.Typ)
	}
	if build.Query != nil {
		searchService.Query(build.Query)
	}
	if build.From >= 0 {
		searchService.From(build.From)
	}
	if build.Size >= 0 {
		searchService.Size(build.Size)
	}
	if build.FetchSourceContext != nil {
		searchService.FetchSourceContext(build.FetchSourceContext)
	}
	if build.SortInfo != nil {
		searchService.SortWithInfo(*build.SortInfo)
	}
	if build.AggregationName != "" && build.Aggregation != nil {
		searchService.Aggregation(build.AggregationName, build.Aggregation)
	}
	if build.Highlight != nil {
		searchService.Highlight(build.Highlight)
	}
	searchService.Pretty(true)

	return searchService
}

//转换返回数据-list
func setReturnValue(result *elastic.SearchHits) (list []interface{}) {
	if result != nil && result.TotalHits > 0 {
		list = make([]interface{}, len(result.Hits))
		len := 0
		for _, hit := range result.Hits {
			var data interface{}
			err := json.Unmarshal(*hit.Source, &data)
			if err == nil {
				list[len] = data
				len++
			}
		}
	}
	return list
}

//match查询
func MatchQuery(build SearchBuild) (list []interface{}) {
	searchResult, err := setSearchService(build).Do(ctx)
	if err != nil {
		panic(err)
	}
	return setReturnValue(searchResult.Hits)
}
