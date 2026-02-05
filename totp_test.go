package totp

import (
	"strings"
	"testing"
)

func TestGenerateSecret(t *testing.T) {
	secret, err := GenerateSecret()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if strings.TrimSpace(secret) == "" {
		t.FailNow()
	}

	t.Log("Generated secret: ", secret)
}

func TestGenerateTotp(t *testing.T) {
	secret := "JBSWY3DPEHPK3PXP"

	otp, err := GenerateTotp(secret, 0)
	if err != nil {
		t.FailNow()
	}

	if otp != "282760" {
		t.FailNow()
	}
}
