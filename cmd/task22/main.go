package main

import "fmt"

type Target interface {
	Request() string
}

type Adaptee struct {
}

func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

type Adapter struct {
	*Adaptee // embedded
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee}

	fmt.Println(adapter.SpecificRequest())
	fmt.Println(adapter.Request())
}
