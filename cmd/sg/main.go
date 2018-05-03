package main

import (
	"flag"
	"fmt"
	"github.com/jimmyjames85/eflag"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"
)

type sgFlags struct {
	// Tos []string `flag:"tos"`
	To      string `flag:"t,to"`
	From    string `flag:"f,from"`
	Body    string `flag:"b,body"`
	Subject string `flag:"s,subject"`
}

var V3Endpoint = "/v3/mail/send"

func main() {

	email := &sgFlags{}
	eflag.StructVar(email)
	flag.Usage = eflag.POSIXStyle
	flag.Parse()

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_APIKEY"), V3Endpoint, os.Getenv("SENDGRID_APIURL"))
	request.Method = "POST"
	request.Body = []byte(fmt.Sprintf(` {
	"personalizations": [
		{
			"to": [
				{
					"email": %q
				}
			],
			"subject": %q
		}
	],
	"from": {
		"email": %q
	},
	"content": [
		{
			"type": "text/plain",
			"value": %q
		}
	]
}`, email.To, email.Subject, email.From, email.Body))
	response, err := sendgrid.API(request)
	if err != nil {
		log.Printf("unable to send: %v", err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

}
