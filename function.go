package main

import (
	"log"
	"time"
)

// defer函数
func BigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	BigSlowOperation()
}
