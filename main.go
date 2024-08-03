package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	flag.Parse()
	fmt.Println("Took:", time.Since(start))
}
