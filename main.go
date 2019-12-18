package main

import "fmt"

import "github.com/freecracy/code/http"

import "github.com/freecracy/code/sendmail"

func main() {
	fmt.Println("code")
	// sendmail
	http.Sendmail()
	sendmail.Sendmail()
}
