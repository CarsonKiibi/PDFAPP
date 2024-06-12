package main

import (
	"fmt"
	"time"

	"github.com/carsonkiibi/pdfapp/backend/process"
)

func main() {
	start := time.Now()
	str := "I like {{B,S14:apples}"
	out := process.ProcessInput(str)
	for _, token := range out {
		fmt.Println(token)
		fmt.Println("--------------")
	}
	duration := time.Since(start)
	fmt.Println(duration.Microseconds())
}
