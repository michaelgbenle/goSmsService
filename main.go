package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var (
	accountSid string
	authToken string
	fromPhone string
	toPhone string
	client *twilio.RestClient
)

func SendMessage(msg string) {
	params := openapi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(fromPhone)
	params.SetBody(msg)

	response, err := client.ApiV2010.CreateMessage(&params)
}