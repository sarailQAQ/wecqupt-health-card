package wecqupt_health_card

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/sarailQAQ/wecqupt-health-card/util"

	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var headers = map[string]string{
	"Host":            "we.cqupt.edu.cn",
	"Connection":      "keep-alive",
	"User-Agent":      "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat",
	"content-type":    "application/json",
	"Referer":         "https://servicewechat.com/wx8227f55dc4490f45/76/page-frame.html",
	"Accept-Encoding": "json",
}

const (
	daka_list = "https://we.cqupt.edu.cn/api/mrdk/get_mrdk_list_test.php"
	daka      = "https://we.cqupt.edu.cn/api/mrdk/post_mrdk_info.php"
)

type ReqBody struct {
	Key string `json:"key"`
}

type DakaResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}
type ListResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		CreatedAt string `json:"created_at"`
		//Name string `json:"name"`
		//stsfjk 身体是否健康
		//szdq
		//xxdz
	} `json:"data"`
}

type Manager struct {
	C Config
}

func NewManager(c Config) *Manager {
	return &Manager{
		C: c,
	}
}
func IsDaka(u UserConfig) bool {
	nowtime := time.Now()
	b, err := json.Marshal(u.BasicInfo)
	if err != nil {
		log.Println(err)
		return false
	}
	decodeString := base64.StdEncoding.EncodeToString(b)
	reqBody := ReqBody{Key: decodeString}

	code, body, err := util.Post(daka_list, headers, reqBody)
	if err != nil {
		log.Println(err)
		return false
	}
	res := ListResponse{}
	_ = json.Unmarshal(body, &res)
	if code != 200 || res.Status != 200 {
		return false
	}
	for _, data := range res.Data {
		s := data.CreatedAt
		parse, err := time.Parse("2006-01-02 15:04:05", s)
		if err != nil {
			log.Println(err)
			continue
		}
		if nowtime.Day() == parse.Day() {
			return true
		}
	}
	return false
}

func (m *Manager) SendRequest(u UserConfig) error {

	u.Timestamp = time.Now().Unix()
	u.Mrdkkey = util.GetKey(time.Now().Day(), time.Now().Hour())

	if IsDaka(u) {
		return nil
	}
	b, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return err
	}
	decodeString := base64.StdEncoding.EncodeToString(b)
	reqBody := ReqBody{Key: decodeString}

	code, body, err := util.Post(daka, headers, reqBody)
	if err != nil {
		log.Println(err)
		return err
	}
	var res DakaResponse
	_ = json.Unmarshal(body, &res)
	if code != 200 || res.Status != 200 {
		err = errors.New("请求失败：" + string(body))
	}

	return err
}

func (m *Manager) selectRandTime() (t time.Time) {
	clock := &m.C.Clock

	i := rand.Intn(len(clock.Clocks))
	hour := clock.Clocks[i]
	if hour > 24 {
		hour = 10
	}
	if clock.Range > 60 {
		clock.Range = 60
	}
	minute := rand.Intn(clock.Range)
	sec := rand.Intn(60)
	now := time.Now()

	return time.Date(now.Year(), now.Month(), now.Day(), hour, minute, sec, 0, time.Local).Add(24 * time.Hour)
}

func (m *Manager) judgeBool(s string) bool {
	s = strings.TrimSpace(s)
	if strings.ToLower(s) == "true" || s == "1" {
		return true
	}
	return false
}

func (m *Manager) Work() {
	c := m.C
	if m.judgeBool(c.Settings.TestMail) {
		err := SendMail("邮箱现在可以用了哦", "^ ^", c.Email)
		if err != nil {
			log.Println(err)
		}
	}

	if m.judgeBool(c.Settings.ImmediateWork) || m.judgeBool(c.Settings.Once) {
		err := m.SendRequest(c.User)
		if err != nil {
			go SendMail("打卡失败", "error: "+err.Error(), c.Email)
			log.Println(err)
		}
	}

	if m.judgeBool(c.Settings.Once) {
		return
	}

	t := m.selectRandTime()
	timer := time.NewTimer(t.Sub(time.Now()))
	for {
		<-timer.C

		if m.judgeBool(c.Settings.RandomPos) {
			latitude, err := randPos(c.User.Latitude)
			if err != nil {
				log.Println(err)
				latitude = c.User.Latitude
			}
			longitude, err := randPos(c.User.Longitude)
			if err != nil {
				log.Println(err)
				longitude = c.User.Longitude
			}

			c.User.Latitude, c.User.Longitude = latitude, longitude
		}

		err := m.SendRequest(c.User)
		if err != nil {
			go SendMail("打卡失败！！！", "error: "+err.Error(), c.Email)
			log.Println(err)

			if m.judgeBool(c.Settings.RetryWhenFailed) {
				now := time.Now()
				end := time.Date(now.Year(), now.Month(), now.Day(), 23, 55, 30, 0, time.Local)
				limit := int(math.Floor(end.Sub(now).Minutes()))

				u := 0
				ticker := time.NewTicker(time.Duration(c.Settings.RetryCountLimit) * time.Minute)
				i := 0
				for ; i < c.Settings.RetryCountLimit; i++ {
					u += c.Settings.RetryTimeGap
					if u >= limit {
						break
					}
					<-ticker.C

					err := m.SendRequest(c.User)
					if err == nil {
						break
					}
					go SendMail("打卡失败！！！", "error: "+err.Error(), c.Email)
				}
				if m.judgeBool(c.Settings.ExitAfterRetryFailed) && (i == c.Settings.RetryCountLimit || u >= limit) {
					err := SendMail("打卡助手即将关闭", "我太没用了qaq", c.Email)
					if err != nil {
						log.Println(err)
					}
					break
				}
			}
		} else {
			go SendMail("打卡成功！", "芜湖起飞", c.Email)
		}

		c = m.C
		t = m.selectRandTime()
		timer.Reset(t.Sub(time.Now()))
	}
}

func randPos(s string) (res string, err error) {
	x, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return
	}
	x += float64(rand.Intn(6)-3) * 0.00001
	return strconv.FormatFloat(x, 'f', 5, 64), err
}
