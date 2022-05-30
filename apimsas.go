package apimsas

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

// Provider is a struct which gives access to SAS generation methods.
// Will cache last token generated until it expires.
// Use NewApimSasToken or NewApimSasTokenDuration to correctly instantiate this structure
type Provider struct {
	id       string
	key      string
	duration time.Duration
	token    string
	expiry   time.Time
}

// NewApimSasProvider returns a new ApimSasToken with a default lifetime of 2 hours
func NewApimSasProvider(id, key string) *Provider {
	return &Provider{
		id:       id,
		key:      key,
		duration: 2 * time.Hour,
	}
}

// NewApimSasProviderDuration returns a new ApimSasToken with a defined lifetime
func NewApimSasProviderDuration(id, key string, duration time.Duration) *Provider {
	return &Provider{
		id:       id,
		key:      key,
		duration: duration,
	}
}

// GetSasToken generates a new Sas Token.
// If a token has been previously generated and it's still valid, then the previous token will be returned
func (p *Provider) GetSasToken() (string, error) {
	if p.isValid() {
		return p.token, nil
	}
	return p.generateSasToken()
}

// Returns true if the token is still usable
func (p *Provider) isValid() bool {
	return !p.expiry.IsZero() && time.Now().Before(p.expiry)
}

// Generates a Sas Token for use with Azure APIM
func (p *Provider) generateSasToken() (string, error) {
	encoder := hmac.New(sha512.New, []byte(p.key))

	p.expiry = time.Now().Add(p.duration).Round(time.Second).UTC()
	fExpiry := p.expiry.Format(time.RFC3339Nano)
	expiry := strings.ReplaceAll(fExpiry, "Z", ".0000000Z")

	dataToSign := p.id + "\n" + expiry

	_, err := encoder.Write([]byte(dataToSign))
	if err != nil {
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(encoder.Sum(nil))

	p.token = fmt.Sprintf("SharedAccessSignature uid=%s&ex=%s&sn=%s", p.id, expiry, signature)

	return p.token, nil
}

// Implements the Stringer interface so that you can:
//
// ```
// sas := NewApimSasProvider("id", "key")
// fmt.Printf("%v", sas)
// ```
func (p *Provider) String() string {
	return fmt.Sprintf("{ Token: %s, Expires: %v, IsValid: %v }", p.token, p.expiry, p.isValid())
}
