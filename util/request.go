package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func Post(url string, headers map[string]string, body interface{}) (code int, resBody []byte, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return
	}
	p := bytes.NewReader(payload)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, p)
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Length", strconv.Itoa(len(payload)))
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		log.Println(res)
		return
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		log.Println(b)
		return
	}

	return res.StatusCode, b, nil
}
