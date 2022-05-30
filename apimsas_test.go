package apimsas

import (
	"testing"
	"time"
)

func checkErr(err error, t *testing.T) {
	if err != nil {
		t.Fatalf("Failed to complete testing: %v", err)
	}
}

// Test that attempting to obtain two tokens within the expiry duration returns the same token twice
func TestSameToken(t *testing.T) {
	sas := NewApimSasProvider("id", "key")

	token1, err := sas.GetSasToken()
	checkErr(err, t)

	// make sure there's some time between token requests
	time.Sleep(2 * time.Second)

	token2, err := sas.GetSasToken()
	checkErr(err, t)

	if token1 != token2 {
		t.Errorf("Expected first and second token to be the same\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

// Test that we can create a provider with a specific duration and that waiting long enough causes two
// different tokens to be generated
func TestProviderWithDuration(t *testing.T) {
	sas := NewApimSasProviderDuration("id", "key", time.Second)

	if sas.duration != time.Second {
		t.Errorf("Token duration is %v, expected %v", sas.duration, time.Second)
	}

	token1, err := sas.GetSasToken()
	checkErr(err, t)

	// make sure there's some time between token requests
	time.Sleep(2 * time.Second)

	token2, err := sas.GetSasToken()
	checkErr(err, t)

	if token1 == token2 {
		t.Errorf("Expected first and second token to be different\nToken 1: %v\nToken 2: %v", token1, token2)
	}

}
