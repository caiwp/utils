package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

func Filter(s string) string {
	chars := [...]string{"\n", "\t", "\"", "\\", " "}
	for _, v := range chars {
		s = strings.Replace(s, v, "", -1)
	}
	return s
}
