package tieba

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits  = "123456789"
)

type crawler struct {
	reqForm  map[string]string
	modProxy bool
	proxyUrl string
}

func copyForm(f map[string]string) map[string]string {
	m := make(map[string]string)
	for k, v := range f {
		m[k] = v
	}
	return m
}

func (c *crawler) getForm() map[string]string {
	return copyForm(c.reqForm)
}

func (c *crawler) initForm() {
	randstr := func(n int, bytelst string) string {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		str := make([]byte, 0, n)
		for i := 0; i < n; i++ {
			str = append(str, bytelst[r.Intn(len(bytelst))])
			// Maybe there are new ways to use packagae rand
		}
		return string(str)
	}
	rd := func(n int) string {
		return randstr(n, digits)
	}
	rld := func(n int) string {
		return randstr(n, letters+digits)
	}

	c.reqForm = map[string]string{
		"_client_id":      "wappc_" + rd(13) + "_" + rd(3),
		"_client_type":    "2",
		"_client_version": "4.5.5",
		"_phone_imei":     rld(15),
		"app_id":          "version_campus",
		"cuid":            rld(32) + "|" + rld(15),
		"from":            "campus",
		"model":           "ANDROID",
		"net_type":        "3",
		"stErrorNums":     "0",
		"stMethod":        "1", // In some requests, this value is set to 2
		"stMode":          "1",
		"stSize":          rd(5),
		"stTime":          rd(3),
		"stTimesNum":      "0",
	}
}

func newTimestamp() string {
	t := time.Now().UnixNano()
	str := fmt.Sprintln(t)
	return str[:13]
}

func signForm(f map[string]string) {
	f["timestamp"] = newTimestamp()
	keylst := make([]string, 0, len(f))
	for k, _ := range f {
		keylst = append(keylst, k)
	}
	sort.Strings(keylst)

	var s bytes.Buffer
	for _, v := range keylst {
		s.WriteString(v + "=" + f[v])
	}
	s.WriteString("tiebaclient!!!")
	//log.Println(s.String())
	sign := fmt.Sprintf("%X", md5.Sum(s.Bytes()))
	f["sign"] = sign
}
