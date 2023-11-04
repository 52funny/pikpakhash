package main

import (
	"fmt"
	"os"
	"time"

	"github.com/52funny/pikpakhash"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "pikpakhash [file]\n")
		return
	}
	path := os.Args[1]
	ph := pikpakhash.Default()
	t := time.Now()
	hash, err := ph.HashFromPath(path)
	t2 := time.Since(t)
	fmt.Println(t2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "calc hash err: %s\n", err.Error())
		return
	}
	fmt.Println(hash)
}
