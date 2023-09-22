package main

import (
	"fmt"
	"time"
)

func main() {
  presetTime := time.Now()

  fmt.Println("Present time is ", presetTime.Format("01-02-2006"))
}
