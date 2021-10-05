package wecqupt_health_card

import (
	"log"
	"testing"
)

func TestSendMail(t *testing.T) {
	c, _ := TestParseConfig("./cmd/config.toml")
	err := SendMail("test", "Hello world.", c.Email)
	log.Println(err)
}
