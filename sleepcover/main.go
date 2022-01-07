package main

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("https://us.kobobooks.com/collections/accessories/products/kobo-libra-h2o-sleepcover?variant=29581366427718")
	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)
	blackCover := doc.Find("span", "class", "atc-button--text")
	result := strings.TrimSpace(blackCover.Text())

	if result == "Add To Cart" {
		auth := smtp.PlainAuth("", "automatedpickle@gmail.com", "Pickle42", "smtp.gmail.com")
		host := "smtp.gmail.com:587"
		from := "automatedpickle@gmail.com"
		to := []string{"8016948594@tmomail.net", "cmpickle@gmail.com"}
		message := `Subject: Kobo Libra Black Sleepcover available

		Kobo Libra Black Sleepcover available. Buy it here https://us.kobobooks.com/collections/accessories/products/kobo-libra-h2o-sleepcover?variant=29581366427718`
		err := smtp.SendMail(host, auth, from, to, []byte(message))
		if err != nil {
			fmt.Println(err)
		}
	}
}
