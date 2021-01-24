package wecqupt_health_card

import (
	"fmt"
	"testing"
)

func TestSendRequest(t *testing.T) {
	c, _ := ParseConfig()
	_ = NewManager(c).SendRequest(c.User)
}

func TestRandPos(t *testing.T) {
	x, err := randPos("30.67373")
	fmt.Println(err, x)
}
