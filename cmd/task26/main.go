package main

import (
	"fmt"
	"time"
)

// todo: поискать еще другие способы.
func main() {
	fmt.Println("start")
	firstSleep(3)
	fmt.Println("end")
}

func firstSleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}
