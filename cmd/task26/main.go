package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	firstSleep(3)
	fmt.Println("end")
}

func firstSleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}
