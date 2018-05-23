package tieba

import (
	"fmt"
	"testing"
)

func _TestTimeStamp(t *testing.T) {
	curtime := newTimestamp()
	fmt.Println(curtime)
	if len(curtime) == 13 {
		t.Log("pass")
	} else {
		t.Log("failed")
	}
}

func TestInitForm(t *testing.T) {
	c := NewCrawler()
	f := c.GetForm()
	for k, v := range f {
		fmt.Println(k, v)
	}
	t.Log("pass")
}
