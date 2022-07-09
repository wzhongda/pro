package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	SetupLogger()
	//simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.sina.com")

}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("status code of %s : %s", url, resp.Status)
		resp.Body.Close()
	}
	return
}
func SetupLogger() {
	logFileLocation, _ := os.OpenFile("./pro/test.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	log.SetOutput(logFileLocation)
}
