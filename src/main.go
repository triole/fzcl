package main

import (
	"fmt"
	"time"
)

func main() {
	parseArgs()
	fmt.Println(fuzzy(time.Now()))
}
