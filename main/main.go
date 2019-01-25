package main

import (
	"encoding/json"
	"fmt"
	"yunyinyue/spider"
)

func main() {
	musicSpider := spider.NewMusic164Spider()
	// comments, _ := musicSpider.GetComments("376635")
	//fmt.Printf("get count: %d\n", len(comments))

	record, _ := musicSpider.GetPlayRecord("62947535")
	jsonData, _ := json.MarshalIndent(record, "", "\t")
	fmt.Println(string(jsonData))
}
