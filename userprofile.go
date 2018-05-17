package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	//turl := "http://tieba.baidu.com/home/get/panel?un=wxyxsx"
	turl := "http://ip.gs"
	proxyUrl, err := url.Parse("socks5://127.0.0.1:1080")
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	// dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080",
	// 	&proxy.Auth{}, &net.Dialer{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// transport := &http.Transport{Dial: dialer.Dial}
	//client := &http.Client{}
	req, err := http.NewRequest("GET", turl, nil)
	req.Header.Add("User-Agent", "curl/7.47.0")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}

}
