package tieba

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (c *crawler) httpPost(srcurl string, payload map[string]string) []byte {
	val := url.Values{}
	for k, v := range payload {
		val.Set(k, v)
	}

	var client *http.Client
	if c.modProxy {
		proxyUrl, _ := url.Parse("socks5://127.0.0.1:1080")
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	} else {
		client = &http.Client{}
	}

	req, err := http.NewRequest("POST", srcurl, strings.NewReader(val.Encode()))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
