package wecqupt_health_card

import (
	"log"
	"testing"
)

func TestSendMail(t *testing.T) {
	c, _ := ParseConfig()
	err := SendMail("test", "Hello world.", c.Email)
	log.Println(err)
}
