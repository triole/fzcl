package main

import (
	"fmt"
	"time"
)

var (
	BUILDTAGS      string
	appName        = "fzcl"
	appDescription = "fuzzy clock"
	appMainversion = "0.1"
)

func main() {
	fmt.Println(fuzzy(time.Now()))
}
