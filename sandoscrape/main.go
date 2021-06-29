package main

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("https://brandonsanderson.com")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	progress := doc.Find("div", "class", "vc_progress_bar").FindAll("small")
	for _, book := range progress {
		fmt.Println(book.Text() + book.Find("span").Text())
	}
}
