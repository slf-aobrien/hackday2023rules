package rules

type Message struct {
	Message     string      `json:"message"`
	Code        string      `json:"status"`
	Extra       interface{} `json:"extra"`
	ElapsedTime string      `json:"elapsedTime"`
}
