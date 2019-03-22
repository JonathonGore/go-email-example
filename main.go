package main

import (
	"io/ioutil"
	"log"
)

func main() {
	contents, err := ioutil.ReadFile("sendgrid.txt")
	if err != nil {
		log.Fatalf("Unable to access API Key: %v", err)
	}

	apiKey := string(contents[:len(contents)-1])
	client, err := NewMailClient(apiKey)
	if err != nil {
		log.Fatalf("Unable to create mail client: %v", err)
	}

	dest := "jack@example.com" // Who we will send an email to.
	if err := client.SendWelcomeEmail("Jack", dest); err != nil {
		log.Printf("Error sending email: %v", err)
	}
}
