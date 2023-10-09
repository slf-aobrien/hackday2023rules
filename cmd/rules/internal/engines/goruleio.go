package engines

import (
	rules "github.com/slf-aobrien/hackday2023rules/internal"
)

func ValidateWithGoRule(user rules.Users) rules.Message {
	//insert real rule set here
	message := rules.Message{}
	message.Message = "Success"
	message.Code = "OK"
	message.Extra = "No Go Rule Validation Performed"

	return message
}
