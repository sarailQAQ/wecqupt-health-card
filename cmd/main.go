package main

import (
	_ "embed"
	"github.com/sarailQAQ/wecqupt-health-card"
	"log"
	"math/rand"
	"time"
)

//go:embed config.toml
var config []byte

func main() {
	rand.Seed(time.Now().Unix())

	c, err := wecqupt_health_card.ParseConfig(config)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("clock-in assistant start work")
	log.SetPrefix("[Error]")                            //日志前缀
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime) //日志输出样式
	wecqupt_health_card.NewManager(c).Work()
}
