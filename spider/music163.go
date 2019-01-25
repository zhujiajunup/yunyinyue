/*
Package spider provides Music163Spider to crawl song/user comments from https://music.163.com
 */
package spider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	url2 "net/url"
	"strconv"
	"strings"
	"yunyinyue/spider/constants"
	"yunyinyue/spider/encrypt"
	"yunyinyue/spider/entity"
)

type Music163Spider struct {
	client  *http.Client
	headers map[string]string
}

func NewMusic164Spider() (spider Music163Spider) {
	headers := make(map[string]string)
	headers["Accept"] = "ext/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	// empty here
	headers["Accept-Encoding"] = ""
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	headers["Host"] = constants.Music163Host
	headers["Cache-Control"] = "no-cache"
	headers["Connection"] = "keep-alive"
	headers["Pragma"] = "no-cache"
	headers["Origin"] = fmt.Sprintf("%s%s", constants.HttpsPrefix, constants.Music163Host)
	headers["Accept"] = "ext/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	return Music163Spider{
		client:  &http.Client{},
		headers: headers,
	}
}

func (spider Music163Spider) GetUserInfo(userId string){

}

func (spider Music163Spider) GetTopListen(userId string){

}

func (spider Music163Spider) GetComments(songId string) (comments []entity.Comment, err error) {

	commentReqBody := entity.CommentRequestBody{
		Rid:       "R_SO_4_" + songId,
		Offset:    "0",
		Totail:    "true",
		Limit:     "20",
		CsrfToken: "",
	}
	commentUrl := fmt.Sprintf("%s%s%s/R_SO_4_%s?csrf_token=", constants.HttpsPrefix, constants.Music163Host, constants.CommentApi, songId)
	content, err := spider.httpPost(commentUrl, spider.headers, commentReqBody)

	comments = make([]entity.Comment, 0)
	for i := 0; i < len(content.Comments); i ++ {
		comments = append(comments, content.Comments[i])
	}
	if err != nil {
		return
	}
	totalPage := int(math.Ceil(float64(content.Total / 20)))
	for curr := 1; curr < totalPage; curr ++ {
		commentReqBody.Offset = strconv.Itoa(curr)
		commentReqBody.Totail = "false"
		content, err := spider.httpPost(commentUrl, spider.headers, commentReqBody)
		if err != nil {
			return comments, err
		}
		for i := 0; i < len(content.Comments); i ++ {
			comments = append(comments, content.Comments[i])
		}
		fmt.Printf("commentCount: %d\t%s\n", len(content.Comments), content.Comments)
		if curr == 10 {
			break
		}
	}

	return comments, err
}

func (spider Music163Spider) dataEncrypt(dataBytes []byte) (content map[string]string) {
	content = make(map[string]string)
	randomBytes := encrypt.Random(16)
	params, err := encrypt.AesEncrypt(string(dataBytes), constants.SrcretKey, constants.AseKey)
	if err != nil {
		fmt.Println(err)
	}
	params, err = encrypt.AesEncrypt(params, string(randomBytes), constants.AseKey)
	if err != nil {
		fmt.Println(err)
	}
	encSecKey := encrypt.RsaEncrypt(string(randomBytes), constants.PubKey, constants.Modulus)
	if err != nil {
		fmt.Println(err)
	}
	content["params"] = string(params)
	content["encSecKey"] = string(encSecKey)
	return content
}

func (spider Music163Spider) httpPost(url string, headers map[string]string, params interface{}) (commentResp entity.CommentResp, err error) {
	body := make(url2.Values)
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return entity.CommentResp{}, err
	}
	encryptResultMap := spider.dataEncrypt(jsonParams)
	body["params"] = []string{encryptResultMap["params"]}
	body["encSecKey"] = []string{encryptResultMap["encSecKey"]}
	request, err := http.NewRequest("POST", url, strings.NewReader(body.Encode()))
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	if err != nil {
		return
	}
	resp, err := spider.client.Do(request)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(data), &commentResp)
	return
}
