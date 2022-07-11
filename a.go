package main

import (
	"fmt"
	"time"
)

func main() {

	/*array := [...]int{9, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Println(*pointer, " ")
	MarrayAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	fmt.Println(MarrayAddress)*/

	bb := "2022-06-21T08:45:04.161Z"
	//test, _ := time.Parse("2006-01-02 15:04:05", bb)
	cc, _ := time.ParseInLocation(time.RFC3339, bb, time.Local)
	dd := cc.Add(8 * time.Hour).Format("2006-01-02 15:04:05")
	fmt.Println(dd)

}
