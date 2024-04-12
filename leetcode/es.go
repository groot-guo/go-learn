package main

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

type IndexSettings struct {
	Settings struct {
		Number_of_shards   int `json:"number_of_shards"`
		Number_of_replicas int `json:"number_of_replicas"`
	} `json:"settings"`
	Mappings struct {
		Properties map[string]interface{} `json:"properties"`
	} `json:"mappings"`
}

func main() {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 定义索引设置和映射
	settings := IndexSettings{
		Settings: struct {
			Number_of_shards   int "json:\"number_of_shards\""
			Number_of_replicas int "json:\"number_of_replicas\""
		}{5, 1},
		Mappings: struct {
			Properties map[string]interface{} "json:\"properties\""
		}{map[string]interface{}{
			"field1": map[string]interface{}{"type": "text"},
			"field2": map[string]interface{}{"type": "date"},
		}},
	}

	// 创建索引
	data, err := json.Marshal(settings)
	if err != nil {
		log.Fatalf("Error marshaling settings: %s", err)
	}

	res, err := es.Indices.Create(
		"my_index",
		es.Indices.Create.WithBody(bytes.NewReader(data)),
	)
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)

	//doc1 := map[string]interface{}{
	//	"title":   "Test Document",
	//	"content": "This is a test document for Elasticsearch.",
	//	"date":    "2023-04-01",
	//}

	res2, err := es.Create(
		"my_index",
		//"my_type",
		"my_id",
		strings.NewReader(`{"title": "Test Document","content": "This is a test document for Elasticsearch.","date": "2023-04-01"}`),
		//nil, // 没有查询参数
	)
	if err != nil {
		log.Fatalf("Error creating document: %s", err)
	}
	defer res2.Body.Close()
	log.Println(res2)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"field1": "some text",
			},
		},
	}
	body, _ := json.Marshal(query)

	// 执行查询
	rep, err := es.Search(
		es.Search.WithIndex("my_index"),
		es.Search.WithBody(bytes.NewReader(body)),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer rep.Body.Close()
	log.Println(rep)

}
