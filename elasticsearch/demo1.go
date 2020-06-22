package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"reflect"
)

var client *elastic.Client

var host = "http://47.103.9.218:9200"

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func init() {
	errlog := log.New(os.Stdout, "APP", log.LstdFlags)

	client, err := elastic.NewClient(elastic.SetErrorLog(errlog), elastic.SetURL(host), elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		panic(err)
	}

	info, code, err := client.Ping(host).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Printf("elasticsearch returned with code： %d， version： %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("elasticsearch version： %v\n", esversion)
}

func create() {
	//使用结构体
	e1 := Employee{
		FirstName: "jerry",
		LastName:  "tom",
		Age:       24,
		About:     "want to find a job",
		Interests: []string{"music", "code", "eat"},
	}

	put1, err := client.Index().
		Index("employees").
		Type("employee").
		Id("1").
		BodyJson(e1).
		Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed employee:%s to index:%s, type:%s\n", put1.Id, put1.Index, put1.Type)

	//使用字符串
	e2 := `{"first_name":"jerry","last_name":"bob","age":23,"about":"want 13k","interesting":["eatting","movie"]"}`
	put2, err := client.Index().Index("employees").Type("employee").Id("2").BodyString(e2).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed employee:%s to index:%s, type:%s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"zhang","last_name":"jh","age":24,"about":"want a job and 13k","interesting":["eatting","movies"]"}`
	put3, err := client.Index().Index("employees").Type("employee").Id("2").BodyString(e3).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed employee:%s to index:%s, type:%s\n", put3.Id, put3.Index, put3.Type)

}

//修改
func update() {
	res, err := client.Update().Index("employees").Type("employee").Id("2").Doc(map[string]interface{}{"age": 30}).Do(context.TODO())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)
}

func delete() {
	res, err := client.Delete().Index("employees").Type("employee").Id("1").Do(context.TODO())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result:%s\n", res.Result)
}

func gets() {
	//通过id查找
	get1, err := client.Get().Index("employees").Type("employee").Id("2").Do(context.TODO())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Get document:%s in version:%d from index:%s, type:%s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}

func query() {
	var res *elastic.SearchResult
	var err error

	//查找所有
	res, err = client.Search("employees").Type("employee").Do(context.TODO())
	if err != nil {
		print(err.Error())
		return
	}

	var emp Employee
	for _, item := range res.Each(reflect.TypeOf(emp)) {
		t := item.(Employee)
		fmt.Printf("emp:%#v\n", t)
	}

	//指定查询
	q := elastic.NewQueryStringQuery("last_name:jh")
	res1, err := client.Search("employees").Type("employee").Query(q).Do(context.TODO())
	if err != nil {
		print(err.Error())
		return
	}
	var emp1 Employee
	for _, item := range res1.Each(reflect.TypeOf(emp1)) {
		t := item.(Employee)
		fmt.Printf("emp1:%#v\n", t)
	}

	if res1.Hits.TotalHits.Value > 0 {
		fmt.Printf("found a total of %d Employee \n", res.Hits.TotalHits)

		for _, hit := range res.Hits.Hits {

			var e Employee
			err := json.Unmarshal(hit.Source, &e)
			if err != nil {
				fmt.Println("Deserialization failed")
			}

			fmt.Printf("Employee name %s : %s\n", e.FirstName, e.LastName)
		}
	} else {
		fmt.Printf("Found no Employee \n")
	}

	//查询年龄大于25的
	boolq := elastic.NewBoolQuery()
	//boolq.Must(elastic.NewMatchQuery("last_name","jh"))
	boolq.Filter(elastic.NewRangeQuery("age").Gt(25))
	res2, err := client.Search("employees").Type("employee").Query(boolq).Do(context.TODO())
	if err != nil {
		print(err.Error())
		return
	}
	var emp2 Employee
	for _, item := range res2.Each(reflect.TypeOf(emp2)) {
		t := item.(Employee)
		fmt.Printf("emp2:%#v\n", t)
	}

	//短语搜索	搜索about子段中有job 字段的
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "job")
	res3, err := client.Search("employees").Type("employee").Query(matchPhraseQuery).Do(context.TODO())
	if err != nil {
		print(err.Error())
		return
	}
	var emp3 Employee
	for _, item := range res3.Each(reflect.TypeOf(emp3)) {
		t := item.(Employee)
		fmt.Printf("emp3:%#v\n", t)
	}

	//分析interests
	aggs := elastic.NewTermsAggregation().Field("interests")
	res4, err := client.Search("employees").Type("employee").Aggregation("all_interests", aggs).Do(context.TODO())
	if err != nil {
		print(err.Error())
		return
	}
	var emp4 Employee
	for _, item := range res4.Each(reflect.TypeOf(emp4)) {
		t := item.(Employee)
		fmt.Printf("emp4:%#v\n", t)
	}

}

func main() {
	create()
	delete()
	update()
	gets()
	query()
}
