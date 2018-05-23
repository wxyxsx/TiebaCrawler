package tieba

import (
	"encoding/json"
	"fmt"
	"log"
)

// TODO add error_code
type ForumInfo struct {
	Id          string `json:"forum_id"`
	Name        string `json:"forum_name"`
	Description string `json:"slogan"`
	Avatar      string `json:"avatar"`
	Membernum   string `json:"member_num"`
	Threadnum   string `json:"thread_num"`
}

type rawResult struct {
	ForumLst []ForumInfo `json:"forum_list"`
}

func (c *crawler) SearchForm(keyword string) []ForumInfo {
	srcurl := "http://c.tieba.baidu.com/c/f/forum/search"
	payload := c.getForm()
	payload["query"] = keyword
	signForm(payload)

	data := c.httpPost(srcurl, payload)

	var raw rawResult
	if err := json.Unmarshal(data, &raw); err != nil {
		log.Fatal(err)
	}
	return raw.ForumLst
}

func (c *crawler) RawSearch(keyword string) []byte {
	srcurl := "http://c.tieba.baidu.com/c/f/forum/search"
	payload := c.getForm()
	payload["query"] = keyword
	signForm(payload)

	data := c.httpPost(srcurl, payload)
	return data
}

func pForumLst(lst []ForumInfo) {
	for _, v := range lst {
		data, err := json.Marshal(v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
}
