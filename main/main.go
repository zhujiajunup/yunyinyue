package main

import (
	"fmt"
	"os"
	"yunyinyue/spider"
	"yunyinyue/spider/sink"
)

func main() {
	musicSpider := spider.NewMusic164Spider()
	comments, _ := musicSpider.GetComments("376635")
	fmt.Printf("get count: %d\n", len(comments))
	mysqlSink, err := sink.NewMysqlSink()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer mysqlSink.Close()
	for _, comment := range comments {
		err := mysqlSink.InsertComment(comment)
		if err != nil {
			fmt.Println(err)
			// break
		}
	}
	//record, _ := musicSpider.GetPlayRecord("62947535")
	//jsonData, _ := json.MarshalIndent(record, "", "\t")
	//fmt.Println(string(jsonData))
}
