package tieba

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	DefaultProxy = "socks5://127.0.0.1:1080"
)

type Crawler interface {
	SetProxy(proxyUrl string)
	SearchForm(keyword string) []ForumInfo
	RawSearch(keyword string) []byte
}

func NewCrawler() Crawler {
	var c crawler
	c.initForm()
	return &c
}

func (c *crawler) SetProxy(proxyUrl string) {
	// Todo add test for proxy
	if len(proxyUrl) == 0 {
		c.modProxy = false
	} else {
		c.modProxy = true
		c.proxyUrl = proxyUrl
	}
}

func (c *crawler) newClient() *http.Client {
	var client *http.Client
	if c.modProxy {
		proxyUrl, _ := url.Parse(c.proxyUrl)
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	} else {
		client = &http.Client{}
	}
	return client
}

func (c *crawler) httpPost(srcurl string, payload map[string]string) []byte {
	client := c.newClient()

	val := url.Values{}
	for k, v := range payload {
		val.Set(k, v)
	}

	req, err := http.NewRequest("POST", srcurl, strings.NewReader(val.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
