package wecqupt_health_card

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestTestParseConfig(t *testing.T) {
	c, _ := TestParseConfig("./cmd/config.toml")
	fmt.Println(c)
}

func TestLog(t *testing.T) {
	c, err := TestParseConfig("./cmd/config.toml")
	if err != nil {
		log.Println(err)
	}

	file, err := os.OpenFile(c.Settings.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //打开日志文件，不存在则创建
	defer file.Close()

	if err == nil {
		log.SetOutput(file)                                 //设置输出流
		log.SetPrefix("[Error]")                            //日志前缀
		log.SetFlags(log.Llongfile | log.Ldate | log.Ltime) //日志输出样式
		log.Println("hello world")
	} else {
		log.Println(err)
	}
}
