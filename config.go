package wecqupt_health_card

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
)

type UserConfig struct {
	Name          string `toml:"name" json:"name"`
	StuNum        string `toml:"stu_num" json:"xh"`
	Sex           string `toml:"sex" json:"xb"`
	Openid        string `toml:"openid" json:"openid"`
	LocationBig   string `toml:"location_big" json:"locationBig"`
	LocationSmall string `toml:"location_small" json:"locationSmall"`
	Latitude      string `toml:"latitude" json:"latitude"`
	Longitude     string `toml:"longitude" json:"longitude"`
	Szdq          string `toml:"szdq" json:"szdq"`
	Xxdz          string `toml:"xxdz" json:"xxdz"`
	Ywjcqzbl      string `toml:"ywjcqzbl" json:"ywjcqzbl"`
	Ywjchblj      string `toml:"ywjchblj" json:"ywjchblj"`
	Xjzdywqzb     string `toml:"xjzdywqzb" json:"xjzdywqzb"`
	Twsfzc        string `toml:"twsfzc" json:"twsfzc"`
	Ywytdzz       string `toml:"ywytdzz" json:"ywytdzz"`
	Remarks       string `toml:"remarks" json:"beizhu"`
	Mrdkkey       string `toml:"-" json:"mrdkkey"`
	Timestamp     int64  `tomll:"-" json:"timestamp"`
}

type ClockConfig struct {
	Clocks []int `toml:"clocks"`
	Range  int   `toml:"range"`
}

type EmailConfig struct {
	Address  string `toml:"address"`
	Password string `toml:"Password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
}

type SettingsConfig struct {
	ImmediateWork        string `toml:"immediate_work"`
	TestMail             string `toml:"test_mail"`
	RetryWhenFailed      string `toml:"retry_when_failed"`
	RetryTimeGap         int    `toml:"retry_time_gap"`
	RetryCountLimit      int    `toml:"retry_count_limit"`
	ExitAfterRetryFailed string `toml:"exit_after_retry_failed"`
	LogPath              string `toml:"log_path"`
}

type Config struct {
	User     UserConfig     `toml:"user"`
	Clock    ClockConfig    `toml:"clock"`
	Email    EmailConfig    `toml:"email"`
	Settings SettingsConfig `toml:"settings"`
}

func ParseConfig() (c Config, err error) {
	fp, err := os.Open("config.toml")
	if err != nil {
		log.Println("open config file:", err)
		return
	}
	defer fp.Close()

	content, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Println("read file error:", err)
		return
	}

	err = toml.Unmarshal(content, &c)
	if err != nil {
		log.Println("parse config error", err)
		return
	}

	return
}
