package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/browser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: stcko <QUESTION> ")
		return
	}

	question := os.Args[1:]
	questionStr := strings.Join(question, "+")

	url := `https://www.google.com/search?q=site%3Astackoverflow.com+` + questionStr

	err := browser.OpenURL(url)
	if err != nil {
		fmt.Printf("Failed to open URL %s: %v\n", url, err)
		return
	}

	fmt.Printf("Opened URL: %s\n", url)
}
