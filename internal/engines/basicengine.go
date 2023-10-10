package engines

import (
	"fmt"
	"log"
	"strings"
	"time"

	rules "github.com/slf-aobrien/hackday2023rules/internal"
)

type Coverage struct {
	AgeLow  int `json:"ageLow"`
	AgeMed  int `json:"ageMed"`
	AgeHigh int `json:"ageHigh"`

	AgeLowCost  float64 `json:"ageLowCost"`
	AgeMedCost  float64 `json:"ageMedCost"`
	AgeHighCost float64 `json:"ageHighCost"`

	Male        float64 `json:"male"`
	Female      float64 `json:"female"`
	NotProvided float64 `json:"np"`

	LowRisk  float64 `json:"lowRisk"`
	MedRisk  float64 `json:"medRisk"`
	HighRisk float64 `json:"highRisk"`
}

func ValidateWithCode(user rules.Users) rules.Message {
	defer duration(track("ValidateWithCode"))
	defer elapsed("My Function")()
	overallStart := time.Now()
	message := rules.Message{}
	validUser := validateUserData(user)
	if !validUser {
		message.Message = "Fail"
		message.Code = "Invalid Member"
		message.Extra = "No Code Validation was Performed"
		return message
	}

	validCoverages := validateUserHasCoverages(user)
	if !validCoverages {
		message.Message = "Fail"
		message.Code = "Invalid Coverages For Member"
		message.Extra = "No Code Validation was Performed"
		return message
	}
	dentalCoverage := buildDentalCoverage()
	lifeCoverage := buildLifeCoverage()
	ltdCoverage := buildLtdCoverage()

	dentalFee := 0.0
	lifeFee := 0.0
	ltdFee := 0.0
	totalFee := 0.0
	aMessage := ""

	if user.HasDental {
		dentalFee = calculateFee(user, dentalCoverage)
		strFee := fmt.Sprintf("%.2f", dentalFee)
		aMessage = aMessage + "Dental Cost: " + strFee + ", "
		totalFee = totalFee + dentalFee
	}
	if user.HasLife {
		lifeFee = calculateFee(user, lifeCoverage)
		strFee := fmt.Sprintf("%.2f", lifeFee)
		aMessage = aMessage + "Life Cost: " + strFee + ", "
		totalFee = totalFee + lifeFee

	}
	if user.HasLtd {
		ltdFee = calculateFee(user, ltdCoverage)
		strFee := fmt.Sprintf("%.2f", ltdFee)
		aMessage = aMessage + "Ltd Cost: " + strFee + ", "
		totalFee = totalFee + ltdFee
	}
	message.Code = "Success"
	message.Message = strings.TrimRight(aMessage, ", ")
	message.Extra = "Total Cost: " + fmt.Sprintf("%.2f", totalFee)
	duration := time.Since(overallStart)
	fmt.Println("Start:   ", overallStart.UnixNano())
	fmt.Printf("elapsed: %v\n", time.Since(overallStart))

	message.ElapsedTime = fmt.Sprintf("%v", duration.String())

	return message
}

func buildDentalCoverage() Coverage {
	dentalCoverage := Coverage{}
	dentalCoverage.AgeHigh = 50
	dentalCoverage.AgeLow = 20

	dentalCoverage.AgeLowCost = 6
	dentalCoverage.AgeMedCost = 10
	dentalCoverage.AgeHighCost = 20

	dentalCoverage.Female = -.05
	dentalCoverage.Male = .1
	dentalCoverage.NotProvided = .1

	dentalCoverage.LowRisk = -.05
	dentalCoverage.MedRisk = 0.0
	dentalCoverage.HighRisk = .1

	return dentalCoverage
}

func buildLifeCoverage() Coverage {
	lifeCoverage := Coverage{}

	lifeCoverage.AgeHigh = 50
	lifeCoverage.AgeLow = 20

	lifeCoverage.AgeLowCost = 10
	lifeCoverage.AgeMedCost = 17
	lifeCoverage.AgeHighCost = 30

	lifeCoverage.Female = -.05
	lifeCoverage.Male = .1
	lifeCoverage.NotProvided = .1

	lifeCoverage.LowRisk = -.05
	lifeCoverage.MedRisk = 0.0
	lifeCoverage.HighRisk = .1

	return lifeCoverage
}

func buildLtdCoverage() Coverage {
	ltdCoverage := Coverage{}

	ltdCoverage.AgeHigh = 50
	ltdCoverage.AgeLow = 20

	ltdCoverage.AgeLowCost = 15
	ltdCoverage.AgeMedCost = 23
	ltdCoverage.AgeHighCost = 35

	ltdCoverage.Female = -.05
	ltdCoverage.Male = .1
	ltdCoverage.NotProvided = .1

	ltdCoverage.LowRisk = -.05
	ltdCoverage.MedRisk = 0.0
	ltdCoverage.HighRisk = .1

	return ltdCoverage
}

func elapsed(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func calculateFee(user rules.Users, coverage Coverage) float64 {
	userAge := calculateAge(user.Birthday)
	baseFee := 1.0
	gMod := .1 //default for male or not provided
	occMod := 0.0
	if userAge < coverage.AgeLow {
		baseFee = coverage.AgeLowCost
	} else if userAge < coverage.AgeHigh {
		baseFee = coverage.AgeMedCost
	} else if userAge >= coverage.AgeHigh {
		baseFee = coverage.AgeHighCost
	}

	if user.Gender == "F" {
		gMod = coverage.Female
	} else if user.Gender == "M" {
		gMod = coverage.Male
	} else if user.Gender == "NP" {
		gMod = coverage.NotProvided
	}

	if user.OcCode == "high risk" {
		occMod = coverage.HighRisk
	} else if user.OcCode == "low risk" {
		occMod = coverage.LowRisk
	} else if user.OcCode == "med risk" {
		occMod = coverage.MedRisk
	}

	calc := baseFee + (baseFee * gMod) + (baseFee * occMod)

	return calc
}

func calculateAge(birthday string) int {
	birthdate, error := time.Parse("01/02/2006", birthday)
	if error != nil {
		fmt.Println(error)
		return 18
	}

	today := time.Now()
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age

}
func validateUserHasCoverages(user rules.Users) bool {
	if user.HasDental {
		return true
	}
	if user.HasLife {
		return true
	}
	if user.HasLtd {
		return true
	}
	return false
}

func validateUserData(user rules.Users) bool {

	if len(user.FirstName) == 0 {
		return false
	}
	if len(user.LastName) == 0 {
		return false
	}
	if len(user.Birthday) == 0 {
		return false
	}
	if len(user.OcCode) == 0 {
		return false
	}
	return true
}
func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
