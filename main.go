package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const Letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Digits = "123456789"

var reqform map[string]string
var viaproxy bool = false

type Sresult struct {
	Image       string `json:"avatar"`
	Name        string `json:"forum_name"`
	Description string `json:"slogan"`
	Id          string `json:"forum_id"`
	Member      string `json:"member_num"`
	Thread      string `json:"thread_num"`
}

type tsearch struct {
	Flist []Sresult `json:"forum_list"`
}

func initform() {
	randstr := func(i int, lst string) string {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rs := make([]byte, 0, i)
		for j := 0; j < i; j++ {
			rs = append(rs, lst[r.Intn(len(lst))])
		}
		return string(rs)
	}
	rs1 := func(i int) string {
		return randstr(i, Digits)
	}
	rs2 := func(i int) string {
		return randstr(i, Letters+Digits)
	}
	reqform = map[string]string{
		"_client_id":      "wappc_" + rs1(13) + "_" + rs1(3),
		"_client_type":    "2",
		"_client_version": "4.5.5",
		"_phone_imei":     rs2(15),
		"app_id":          "version_campus",
		"cuid":            rs2(32) + "|" + rs2(15),
		"from":            "campus",
		"model":           "ANDROID",
		"net_type":        "3",
		"stErrorNums":     "0",
		"stMethod":        "1", //2
		"stMode":          "1",
		"stSize":          rs1(5),
		"stTime":          rs1(3),
		"stTimesNum":      "0",
	}
}

func copyform(f map[string]string) map[string]string {
	m := make(map[string]string)
	for k, v := range f {
		m[k] = v
	}
	return m
}

func timestamp() string {
	t := time.Now().UnixNano()
	str := strconv.Itoa(int(t))
	return string(str[:13])
}

func signform(f map[string]string) {
	f["timestamp"] = timestamp()
	lst := make([]string, 0, len(f))
	for k, _ := range f {
		lst = append(lst, k)
	}
	sort.Strings(lst)
	var s bytes.Buffer
	for _, v := range lst {
		s.WriteString(v + "=" + f[v])
	}
	s.WriteString("tiebaclient!!!")
	log.Println(s.String())
	sign := fmt.Sprintf("%X", md5.Sum(s.Bytes()))
	f["sign"] = sign
}

func httpPost(curl string, payload map[string]string) []byte {
	val := url.Values{}
	for k, v := range payload {
		val.Set(k, v)
	}

	var client *http.Client
	if viaproxy {
		proxyUrl, _ := url.Parse("socks5://127.0.0.1:1080")
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	} else {
		client = &http.Client{}
	}

	req, err := http.NewRequest("POST", curl, strings.NewReader(val.Encode()))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func getLikeForm(id string) {
	url := "http://c.tieba.baidu.com/c/f/forum/like"
	payload := copyform(reqform)
	payload["uid"] = id
	signform(payload)
	str := httpPost(url, payload)
	fmt.Println(str)
}

func searchForm(keyword string) {
	url := "http://c.tieba.baidu.com/c/f/forum/search"
	payload := copyform(reqform)
	payload["query"] = keyword
	signform(payload)
	data := httpPost(url, payload)

	var sr tsearch
	jsonerr := json.Unmarshal(data, &sr)
	if jsonerr != nil {
		log.Fatal(jsonerr)
	}
	for _, v := range sr.Flist {
		ndata, _ := json.Marshal(v)
		fmt.Println(string(ndata))
	}

}

func main() {
	initform()
	viaproxy = true
	searchForm("巨人")
}
