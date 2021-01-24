package main

import (
	"github.com/sarailQAQ/wecqupt-health-card"
	"log"
	"os"
)

func main() {
	c, err := wecqupt_health_card.ParseConfig()
	if err != nil {
		log.Println(err)
		return
	}

	file, err := os.OpenFile(c.Settings.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)    //打开日志文件，不存在则创建
	defer file.Close()

	if err == nil{
		log.SetOutput(file)                                 //设置输出流
	}

	log.Println("clock-in assistant start work")
	log.SetPrefix("[Error]")  //日志前缀
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime) //日志输出样式
	wecqupt_health_card.NewManager(c).Work()
}
