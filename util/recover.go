package util

import "log"

func Recover() {
	if err := recover(); err != nil {
		log.Println(err)
	}
}
