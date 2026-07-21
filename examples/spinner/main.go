package main

import (
	"time"

	"github.com/zynaps-365/nicecli"
)

func veryLongFunction() string {
	time.Sleep(5 * time.Second)
	return "Admin"
}

func main() {
	data := nicecli.Spinner("Loading...", veryLongFunction)

	nicecli.Success("Data received! Hi, %s!", data)
}
