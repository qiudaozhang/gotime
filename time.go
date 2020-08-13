package gotime

import (
	"log"
	"time"
)

const (
	LAYOUT20060102150405 = "2006-01-02 15:04:05"
	LAYOUT20060102       = "2006-01-02"
	LAYOUT150405         = "15:04:05"
)

func ParseRFC3339TimeYYMMDDHHMMSS(t string) string {
	//2006-01-02T15:04:05+08:00--> 2006-01-02 15:04:05
	newTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Println(err.Error())
	}
	formatTime := newTime.Format(LAYOUT20060102150405)
	return formatTime
}

func ParseRFC3339TimeYYMMDD(t string) string {
	//2006-01-02T15:04:05+08:00--> 2006-01-02
	newTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Println(err.Error())
	}
	formatTime := newTime.Format(LAYOUT20060102)
	return formatTime
}

func ParseRFC3339TimeHHMMSS(t string) string {
	//2006-01-02T15:04:05+08:00--> 15:04:05
	newTime, err := time.Parse(time.RFC3339, t)
	if err != nil {
		log.Println(err.Error())
	}
	formatTime := newTime.Format(LAYOUT150405)
	return formatTime
}
