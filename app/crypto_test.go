package app

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	plaintext := "ZAHIRDBA"
	encrypted, err := Encrypt(plaintext)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	decrypted, err := Decrypt(encrypted)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	if decrypted != plaintext {
		t.Errorf("Expected decrypted [%v], got [%v]", plaintext, decrypted)
	}
}

func TestBackwardCompatibilityDecrypt(t *testing.T) {
	encrypted := "85I1it9lcLpTe2s4OFW/jA=="
	expected := "ZAHIRDBA"
	decrypted, err := Decrypt(encrypted)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	if decrypted != expected {
		t.Errorf("Expected decrypted [%v], got [%v]", expected, decrypted)
	}
}
