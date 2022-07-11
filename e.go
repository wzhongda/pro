package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	str := "149109"
	h := md5.New()
	h.Write([]byte(str))
	c := hex.EncodeToString(h.Sum(nil))
	fmt.Println(c)
	d := "7f57b94559f7e4e03f91cf10a6d3f5e6844fdd2dc106fec7a3e6f8061a6cfd14"
	m := strings.Replace(d, c, "", -1)
	fmt.Println(m, len(m))
	dd := "844fdd2dc106fec7a3e6f8061a6cfd14"
	fmt.Println(len(dd))
}
