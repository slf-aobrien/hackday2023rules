package engines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	rules "github.com/slf-aobrien/hackday2023rules/internal"
)

type GoRuleResponse struct {
	Performance string `json:"performance"`
	Result      struct {
		Fee float64 `json:"fee"`
	} `json:"result"`
}

type GoRuleRequest struct {
	Context GoRuleContext `json:"context"`
}

type GoRuleContext struct {
	Member GoRuleMember `json:"member"`
}

type GoRuleMember struct {
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	OcCode string `json:"ocCode"`
}

func ValidateWithGoRule(user rules.Users) rules.Message {
	defer duration(track("ValidateWithCode"))
	defer elapsed("GoRules")()
	overallStart := time.Now()

	aMessage := ""

	// Hard coding example gorules demo URL and token here for now (better to move into env file)
	posturl := "https://eu.engine.gorules.io/documents/d1284028-cdf0-43d0-aa92-dcb781a61156/demo/policies/life"
	token := "EXun1yh0xf7upSQOPIIMxsVy"

	// @todo validate user - copy from basic engine
	// @todo validate coverages - copy from basic engine

	// Build payload

	// if user.HasDental {

	// }

	// if user.HasLife {

	// }

	// if user.HasLtd {

	// }

	payload := &GoRuleRequest{
		Context: GoRuleContext{
			Member: GoRuleMember{
				Age:    25,
				Gender: "male",
				OcCode: user.OcCode,
			},
		},
	}

	body, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(body))

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

	response := &GoRuleResponse{}
	derr := json.NewDecoder(res.Body).Decode(response)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}

	// @todo replace with Message from response
	strFee := fmt.Sprintf("%.2f", response.Result.Fee)
	aMessage = aMessage + "Life Cost: " + strFee + ", "

	message := rules.Message{}
	message.Message = strings.TrimRight(aMessage, ", ")
	message.Code = "OK"
	message.Extra = "Total Cost: " + fmt.Sprintf("%.2f", response.Result.Fee)

	duration := time.Since(overallStart)
	fmt.Println("Start:   ", overallStart.UnixNano())
	fmt.Printf("elapsed: %v\n", time.Since(overallStart))

	message.ElapsedTime = fmt.Sprintf("local %v, remote %v", duration.String(), response.Performance)

	return message
}
