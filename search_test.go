package tieba

import (
	"fmt"
	"testing"
)

func TestNetwork(t *testing.T) {
	c := NewCrawler()
	result := c.RawSearch("wp")
	fmt.Println(string(result[:100]))
	t.Log("pass")
}

func TestProxy(t *testing.T) {
	c := NewCrawler()
	c.SetProxy(DefaultProxy)
	result := c.RawSearch("android")
	fmt.Println(string(result[:100]))
	t.Log("pass")
}

func TestJson(t *testing.T) {
	c := NewCrawler()
	result := c.SearchForm("八中")
	pForumLst(result)
	t.Log("pass")
}
