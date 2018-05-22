package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

//get the profile of tieba users
//id name name_show portrait? sex tb_age post_num followed_count is_private
//
type Profile struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Nameshow  string `json:"name_show"`
	Image     string `json:"portrait_h"`
	Sex       string `json:"sex"`
	Tbage     string `json:"tb_age"`
	Postnum   int    `json:"post_num"`
	Isprivate int    `json:"is_private"`
	Followed  int    `json:"followed_count"`
}

type tprofile struct {
	Data Profile `json:"data"`
}

func main() {
	turl := "http://tieba.baidu.com/home/get/panel?un=wxyxsx"

	proxyUrl, err := url.Parse("socks5://127.0.0.1:1080")
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	req, err := http.NewRequest("GET", turl, nil)

	//req.Header.Add("User-Agent", "curl/7.47.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	var cprofile tprofile
	jsonerr := json.Unmarshal(b, &cprofile)
	if jsonerr != nil {
		log.Fatal(err)
	}
	// 	for k, v := range cprofile.Data {
	// 		fmt.Println("map:", k, v)
	// 	}
	ndata, err := json.Marshal(cprofile.Data)
	fmt.Println(string(ndata))

	f, err := os.OpenFile("tempdata", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(ndata); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
