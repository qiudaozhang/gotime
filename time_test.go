package gotime

import (
	"fmt"
	"testing"
)

func TestParseRFC3339TimeYYMMDDHHMMSS(t *testing.T) {
	time := "2015-04-28T18:16:19+08:00"
	result := ParseRFC3339TimeYYMMDDHHMMSS(time)
	fmt.Println(result)
}

func TestParseRFC3339TimeYYMMDD(t *testing.T) {
	time := "2015-04-28T18:16:19+08:00"
	result := ParseRFC3339TimeYYMMDD(time)
	fmt.Println(result)
}
