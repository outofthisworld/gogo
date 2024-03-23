package main

import (
	"fmt"
	"practice/time"
)

func main() {
	fmt.Printf("time is ::: %s", time.Millseconds(2000).ToSeconds().ToTime().AsString())
}
