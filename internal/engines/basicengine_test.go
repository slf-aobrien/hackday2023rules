package engines

import (
	"fmt"
	"log"
	"testing"

	rules "github.com/slf-aobrien/hackday2023rules/internal"
)

func TestValidateWithEmptyMember(t *testing.T) {
	emptyUser := rules.Users{}

	fmt.Println("Running Test Validate With Empty Member ", emptyUser.FirstName)
	msg := ValidateWithCode(emptyUser)
	if msg.Message != "Fail" {
		t.Fatalf("User not validated wanting %q received %q", "validated", msg.Message)
	}
}

func TestValidateWithFullMemberNoCoverages(t *testing.T) {
	testUser := rules.Users{}
	testUser.FirstName = "Aaron"
	testUser.LastName = "O'Brien"
	testUser.Birthday = "A Day"
	testUser.OcCode = "Risky"
	fmt.Println("Running Test Validate With Empty Member ", testUser.FirstName)
	msg := ValidateWithCode(testUser)
	if msg.Message != "Fail" {
		t.Fatalf("Test Faield wanting %q received %q", "validated", msg.Message)
	}
}

func TestValidateWithFullMemberWithDentalHighRisk(t *testing.T) {
	testUser := rules.Users{}
	testUser.FirstName = "Aaron"
	testUser.LastName = "O'Brien"
	testUser.Birthday = "A Day"
	testUser.OcCode = "high risk"
	testUser.HasDental = true
	testUser.HasLife = false
	testUser.HasLtd = false
	t.Log("Running Test Dental Coverage ", testUser.FirstName)
	msg := ValidateWithCode(testUser)
	log.Print(msg.Message)
	if msg.Message != "Dental Cost: 7.20" {
		t.Fatalf("Test Failed wanting %q received %q", "Dental Cost: 7.20", msg.Message)
	}
}

func TestValidateWithFullMemberWithDentalLowRisk(t *testing.T) {
	testUser := rules.Users{}
	testUser.FirstName = "Aaron"
	testUser.LastName = "O'Brien"
	testUser.Birthday = "A Day"
	testUser.OcCode = "low risk"
	testUser.HasDental = true
	testUser.HasLife = false
	testUser.HasLtd = false
	t.Log("Running Test Dental Coverage ", testUser.FirstName)
	msg := ValidateWithCode(testUser)
	log.Print(msg.Message)
	if msg.Message != "Dental Cost: 6.30" {
		t.Fatalf("Test Failed wanting %q received %q", "Dental Cost: 6.30", msg.Message)
	}
}

func TestValidateWithFullMemberWithLifeHighRisk(t *testing.T) {
	testUser := rules.Users{}
	testUser.FirstName = "Aaron"
	testUser.LastName = "O'Brien"
	testUser.Birthday = "A Day"
	testUser.OcCode = "high risk"
	testUser.HasDental = false
	testUser.HasLife = true
	testUser.HasLtd = false
	t.Log("Running Test Life Coverage ", testUser.FirstName)
	msg := ValidateWithCode(testUser)
	log.Print(msg.Message)
	if msg.Message != "Life Cost: 12.00" {
		t.Fatalf("Test Failed wanting %q received %q", "Dental Cost: 7.20", msg.Message)
	}
}

func TestValidateWithFullMemberWithLifeLowRisk(t *testing.T) {
	testUser := rules.Users{}
	testUser.FirstName = "Aaron"
	testUser.LastName = "O'Brien"
	testUser.Birthday = "A Day"
	testUser.OcCode = "low risk"
	testUser.HasDental = false
	testUser.HasLife = true
	testUser.HasLtd = false
	t.Log("Running Test Life Coverage ", testUser.FirstName)
	msg := ValidateWithCode(testUser)
	log.Print(msg.Message)
	if msg.Message != "Life Cost: 10.50" {
		t.Fatalf("Test Failed wanting %q received %q", "Life Cost: 10.50", msg.Message)
	}
}

func TestValidateWithFullMemberWithLtdHighRisk(t *testing.T) {
	testUser := rules.Users{}
	testUser.FirstName = "Aaron"
	testUser.LastName = "O'Brien"
	testUser.Birthday = "A Day"
	testUser.OcCode = "high risk"
	testUser.HasDental = false
	testUser.HasLife = false
	testUser.HasLtd = true
	t.Log("Running Test LTD Coverage ", testUser.FirstName)
	msg := ValidateWithCode(testUser)
	log.Print(msg.Message)
	if msg.Message != "Ltd Cost: 18.00" {
		t.Fatalf("Test Failed wanting %q received %q", "Ltd Cost: 18.00", msg.Message)
	}
}

func TestValidateWithFullMemberWithLtdLowRisk(t *testing.T) {
	testUser := rules.Users{}
	testUser.FirstName = "Aaron"
	testUser.LastName = "O'Brien"
	testUser.Birthday = "A Day"
	testUser.OcCode = "low risk"
	testUser.HasDental = false
	testUser.HasLife = false
	testUser.HasLtd = true
	t.Log("Running Test LTD Coverage ", testUser.FirstName)
	msg := ValidateWithCode(testUser)
	log.Print(msg.Message)
	if msg.Message != "Ltd Cost: 15.75" {
		t.Fatalf("Test Failed wanting %q received %q", "Ltd Cost: 15.75", msg.Message)
	}
}

func TestValidateWithFullMemberAllCoveragesHighRisk(t *testing.T) {
	testUser := rules.Users{}
	testUser.FirstName = "Aaron"
	testUser.LastName = "O'Brien"
	testUser.Birthday = "A Day"
	testUser.OcCode = "high risk"
	testUser.HasDental = true
	testUser.HasLife = true
	testUser.HasLtd = true
	t.Log("Running Test LTD Coverage ", testUser.FirstName)
	msg := ValidateWithCode(testUser)
	log.Print(msg.Message)
	if msg.Extra != "Total Cost: 37.20" {
		t.Fatalf("Test Failed wanting %q received %q", "Total Cost: 37.20", msg.Extra)
	}
}

func TestValidateWithFullMemberAllCoveragesLowRisk(t *testing.T) {
	testUser := rules.Users{}
	testUser.FirstName = "Aaron"
	testUser.LastName = "O'Brien"
	testUser.Birthday = "A Day"
	testUser.OcCode = "low risk"
	testUser.HasDental = true
	testUser.HasLife = true
	testUser.HasLtd = true
	t.Log("Running Test LTD Coverage ", testUser.FirstName)
	msg := ValidateWithCode(testUser)
	log.Print(msg.Message)
	if msg.Extra != "Total Cost: 32.55" {
		t.Fatalf("Test Failed wanting %q received %q", "Total Cost: 32.55", msg.Extra)
	}
}
