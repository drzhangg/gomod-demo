package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"reflect"
	"time"
)

type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_filed,omitempty"`
}

const mapping = `
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0
  },
  "mappings": {
    "tweet": {
      "properties": {
        "user": {
          "type": "keyword"
        },
        "message": {
          "type": "text",
          "store": true,
          "fielddata": true
        },
        "image": {
          "type": "keyword"
        },
        "created": {
          "type": "date"
        },
        "tags": {
          "type": "keyword"
        },
        "location": {
          "type": "geo_point"
        },
        "suggest_field": {
          "type": "completion"
        }
      }
    }
  }
}
`

func main() {

	ctx := context.TODO()

	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	//连接es测试
	info, code, err := client.Ping("http://47.103.9.218:9200").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("elasticsearch returned with code:%d and version:%s\n", code, info.Version.Number)

	exists, err := client.IndexExists("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(exists)

	if !exists {
		createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {

		}
	}

	tweet1 := Tweet{
		User:     "jerry",
		Message:  "want a job",
		Retweets: 0,
	}

	put1, err := client.Index().Index("twitter").Type("tweet").Id("1").BodyJson(tweet1).Do(ctx)
	if err != nil {
		log.Panicln("put1 err:", err)
	}
	fmt.Printf("indexed tweet:%s to index:%s,type:%s\n", put1.Id, put1.Index, put1.Type)

	tweet2 := `{"user":"tom","message":"want a job too"}`
	put2, err := client.Index().Index("twitter").Type("tweet").Id("2").BodyString(tweet2).Do(ctx)
	if err != nil {
		log.Panicln("put2 err:", err)
	}
	fmt.Printf("indexed tweet:%s to index:%s,type:%s\n", put2.Id, put2.Index, put2.Type)

	//Get
	get1, err := client.Get().Index("twitter").Type("tweet").Id("1").Do(ctx)
	if err != nil {
		log.Panicln("get1 err:", err)
	}
	fmt.Printf("get1 document:%s in version:%d from index:%s, type:%s\n", get1.Id, get1.Version, get1.Index, get1.Type)

	_, err = client.Flush().Index("twitter").Do(ctx)
	if err != nil {
		log.Panicln("get1 err:", err)
	}
	//fmt.Println(flush.Shards)

	termQuery := elastic.NewTermQuery("user", "zhang")
	searchResult, err := client.Search().Index("twitter").Query(termQuery).Sort("user.keyword", true).From(0).Size(10).Pretty(true).Do(ctx)
	if err != nil {
		log.Panicln("search err:", err)
	}
	fmt.Printf("query took:%d milliseconds\n", searchResult.TookInMillis)

	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("found a total of %d tweets\n", searchResult.Hits.TotalHits)

		for _, hit := range searchResult.Hits.Hits {
			var t Tweet
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	} else {
		// No hits
		fmt.Print("Found no tweets\n")
	}

	getResult, err := client.Get().Index("twitter").Type("tweet").Id("1").Do(ctx)
	if err != nil {
		log.Panicln("getResult err:", err)
	}

	var t Tweet
	json.Unmarshal(*getResult.Source, &t)
	fmt.Println(t)

	res, err := client.Search("twitter").Type("tweet").Do(ctx)
	fmt.Println(res.Hits.TotalHits)

	if res.Hits.TotalHits > 0 {
		var t Tweet
		for _, item := range res.Hits.Hits {
			//fmt.Println(item.Source)
			json.Unmarshal(*item.Source, &t)
			fmt.Println(t)
		}
	}

	fuzQ := elastic.NewFuzzyQuery("message", "want")
	sr, err := client.Search("twitter").Type("tweet").Query(fuzQ).Do(ctx)
	if err != nil {
		log.Panicln("fuzQ err:", err)
	}
	fmt.Println(sr.Hits.Hits)

	var t1 Tweet
	for _,item := range sr.Hits.Hits{
		json.Unmarshal(*item.Source, &t1)
		fmt.Println(t1.User,t1.Message)
	}
}
