package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := []int{303, 404, 506, 7}

	for _, item := range a {
		b := strconv.Itoa(item)
		c := strings.Split(b, "")
		fmt.Println(c[0])
	}

}
