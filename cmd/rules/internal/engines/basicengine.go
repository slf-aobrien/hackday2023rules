package basicengine

func validate(user Users) Message {
	//insert real rule set here
	message := Message{}
	message.Message = "Success"
	message.Code = "OK"
	message.Extra = "No Validation Performed"

	return message
}
