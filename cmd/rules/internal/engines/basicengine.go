package engines


rules "github.com/slf-aobrien/hackday2023rules"

func Validate(user rules.Users) Message {
	//insert real rule set here
	message := Message{}
	message.Message = "Success"
	message.Code = "OK"
	message.Extra = "No Validation Performed"

	return message
}
