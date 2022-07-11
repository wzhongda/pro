package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Println(i, " ")
	}

}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Println(i, " ")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Println(n, " ")
		}(i)
	}
}

func one(aLog *log.Logger) {
	aLog.Println("function one ---")

	defer aLog.Println("---function one ----")
	for i := 0; i < 10; i++ {
		aLog.Println(i)
	}
}

var LOGFILE = "mGo.log"

func pLog() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "logDefer ", log.LstdFlags)
	iLog.Println("hello world")
	iLog.Println("method log entry!")
	one(iLog)
}

func ap() {
	fmt.Println("a...")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println(c)
			fmt.Println("recovered inside a")
		}
	}()
	bp()
	fmt.Println("b ....")
}
func bp() {
	fmt.Println("inside b ")
	panic("panic in b")
	fmt.Println("exiting b")
}

func pp() {
	if len(os.Args) == 1 {
		panic("not enough arguments")
	}
	fmt.Println("thanks for the argument")
}
func goEv() {
	fmt.Println(runtime.Compiler, runtime.GOARCH, runtime.Version(), runtime.NumCPU(), runtime.NumGoroutine())
}

func mapDel() {
	a := make(map[string]int, 0)
	delete(a, "a")
	aMap := map[string]int{}
	aMap = nil
	fmt.Println(aMap)
	delete(aMap, "b")

}
func tTime() {
	a := time.Now().Unix()
	fmt.Println(a)
	b := time.Now().Format(time.RFC3339)
	fmt.Println(b)
	c := time.Now()
	fmt.Println(c.Weekday().String())
	fmt.Println(c.Day())
	fmt.Println(c.Month(), c.Year(), c.YearDay())
	time.Sleep(time.Second * 10)
	t1 := time.Now()
	fmt.Println(t1.Sub(c))
	fmt.Println(time.Since(c))

}
func main() {

	/*d1()
	d2()
	d3()
	pLog()*/
	/*	ap()
		fmt.Println("main ended")*/
	/*pp()*/
	//goEv()
	//mapDel()
	tTime()
}
