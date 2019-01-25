package main

import (
	"fmt"
	"yunyinyue/spider"
)

func main() {
	musicSpider := spider.NewMusic164Spider()
	comments, _ := musicSpider.GetComments("376635")
	fmt.Printf("get count: %d\n", len(comments))
}
