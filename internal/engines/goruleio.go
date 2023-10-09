package engines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	rules "github.com/slf-aobrien/hackday2023rules/internal"
)

type Response struct {
	Performance string `json:"performance"`
	Result      struct {
		Fee int `json:"fee"`
	} `json:"result"`
}

func ValidateWithGoRule(user rules.Users) rules.Message {
	// Hard coding example gorules demo URL and token here for now (better to move into env file)
	posturl := "https://eu.engine.gorules.io/documents/d1284028-cdf0-43d0-aa92-dcb781a61156/demo/policies/life"
	token := "EXun1yh0xf7upSQOPIIMxsVy"

	body := []byte(`{
		"context": {
			"member": {
				"age": 25,
				"gender": "male",
				"risk": "high"
			}
		}
	}`)

	// Create a HTTP post request
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("X-Access-Token", token)
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	response := &Response{}
	derr := json.NewDecoder(res.Body).Decode(response)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}

	fmt.Println(response.Result.Fee)

	message := rules.Message{}
	message.Message = "Success"
	message.Code = "OK"
	message.Extra = "Total Cost: " + fmt.Sprintf("%d", response.Result.Fee)

	return message
}
