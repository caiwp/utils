package utils

import "testing"

import "fmt"

func TestStrToTime(t *testing.T) {
	s := "2017-08-21 14:51:22"
	i, err := StrToTime(s, DateTimeFormat)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(i)

	s = "2017-08-21"
	i, err = StrToTime(s, DateFormat)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(i)
}

func TestDate(t *testing.T) {
	d := 1503244800
	s := Date(DateFormat, int64(d))
	fmt.Println(s)

	d = 1503298289
	s = Date(DateTimeFormat, int64(d))
	fmt.Println(s)
}
