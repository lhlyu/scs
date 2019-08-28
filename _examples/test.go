package main

//func main() {
//
//	url := "http://192.168.1.207"
//	port := "9200"
//	client, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("%s:%s", url, port)))
//	fmt.Println(fmt.Sprintf("%s:%s", url, port))
//	if err != nil {
//		panic("failed to connect redis:" + err.Error())
//	}
//	client.DeleteByQuery().Index("1").Type("_doc").Query(elastic.NewMatchAllQuery()).Do(context.Background())
//	elastic.NewQueryStringQuery("")
//}

// http://192.168.1.207:9200/_cat/indices/key  搜索索引
// http://192.168.1.207:9200/community_active_info/_search?pretty