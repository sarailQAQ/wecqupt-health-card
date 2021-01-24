package wecqupt_health_card

import "testing"

func TestSendRequest(t *testing.T) {
	c, _ := ParseConfig()
	_ = NewManager(c).SendRequest()


}
