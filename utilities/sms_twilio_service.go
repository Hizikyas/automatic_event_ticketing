package utilities

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(to string, messageBody string) error {
	client := twilio.NewRestClient()

	from := os.Getenv("TWILIO_PHONE_NUMBER") 
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(messageBody)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("twilio error: %v", err)
	}
	return nil
}