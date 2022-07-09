package main

import (
	"fmt"
	"strings"
)

func PasswordDecode(pwd string) string {
	pwdTrim := strings.Split(pwd, "")
	leftPwd := pwdTrim[0:32]
	rightPwd := pwdTrim[32:64]
	tempPwd := rightPwd[0:16]
	temp := make([]string, 0)
	for _, item := range tempPwd {
		temp = append(temp, item)
	}
	for i := 0; i < 16; i++ {
		pos := 2 * i
		rightPwd[i] = leftPwd[pos]
		leftPwd[pos] = temp[i]
	}
	for _, item := range rightPwd {
		leftPwd = append(leftPwd, item)
	}
	pwdString := ""
	for _, item := range leftPwd {
		pwdString += item
	}
	return pwdString
}

func main() {
	a := "00112222333333334444444444444444aabbccccddddddddeeeeeeeeeeeeeeee"
	//a := "a0a1b2b2c3c3c3c3d4d4d4d4d4d4d4d40122333344444444eeeeeeeeeeeeeeee"
	b := PasswordDecode(a)
	fmt.Println("b:", b)
	c := PasswordDecode(b)
	fmt.Println("c:", c)
}
