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

// This is a connector for a gorules.io decision engine
// The admin UI and API are hosted in a free tier account, but in theory the standalone self hosted solution should work as well using the components
// https://github.com/gorules/zen
// https://github.com/gorules/jdm-editor

type GoRuleResponse struct {
	Performance string `json:"performance"`
	Result      struct {
		Total  float64 `json:"total"`
		Life   float64 `json:"life"`
		Dental float64 `json:"dental"`
		Ltd    float64 `json:"ltd"`
	} `json:"result"`
}

type GoRuleRequest struct {
	Context GoRuleContext `json:"context"`
}

type GoRuleContext struct {
	Member GoRuleMember `json:"member"`
}

// @todo extend User struct?
type GoRuleMember struct {
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	OcCode    string `json:"ocCode"`
	HasDental bool   `json:"hasDental"`
	HasLife   bool   `json:"hasLife"`
	HasLtd    bool   `json:"hasLtd"`
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
	// @todo - calculate age
	payload := &GoRuleRequest{
		Context: GoRuleContext{
			Member: GoRuleMember{
				Age:       25,
				Gender:    user.Gender,
				OcCode:    user.OcCode,
				HasLife:   user.HasLife,
				HasDental: user.HasDental,
				HasLtd:    user.HasLtd,
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

	// @todo - for some reason not all of the response values are being parsed correctly, only "life"

	if user.HasDental {
		strFee1 := fmt.Sprintf("%.2f", response.Result.Dental)
		aMessage = aMessage + "Dental Cost: " + strFee1 + ", "
	}

	if user.HasLife {
		strFee2 := fmt.Sprintf("%.2f", response.Result.Life)
		aMessage = aMessage + "Life Cost: " + strFee2 + ", "
	}

	if user.HasLtd {
		strFee3 := fmt.Sprintf("%.2f", response.Result.Ltd)
		aMessage = aMessage + "Ltd Cost: " + strFee3 + ", "
	}

	message := rules.Message{}
	message.Message = strings.TrimRight(aMessage, ", ")
	message.Code = "OK"
	message.Extra = "Total Cost: " + fmt.Sprintf("%.2f", response.Result.Total)

	duration := time.Since(overallStart)
	fmt.Println("Start:   ", overallStart.UnixNano())
	fmt.Printf("elapsed: %v\n", time.Since(overallStart))

	message.ElapsedTime = fmt.Sprintf("local %v, remote %v", duration.String(), response.Performance)

	return message
}
