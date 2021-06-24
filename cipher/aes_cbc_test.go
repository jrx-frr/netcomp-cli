package cipher

import (
	"crypto/aes"
	"testing"
)

const key string = "7465737474657374746573747465737474657374746573747465737474657374"
const value string = "test"
const encryptedValue string = "6cGOThTc2d8G4aAF+R34FQ=="

func TestEncryption(t *testing.T) {
	// Test for empty key
	_, err := EncryptAES("", value)
	if err == nil {
		t.Error("EncryptAES with empty key failed, expected error, got nil")
	}

	// Test for empty value
	_, err = EncryptAES(key, "")
	if err == nil {
		t.Error("EncryptAES with empty value failed, expected error, got nil")
	}

	// Test encryption result
	result, err := EncryptAES(key, value)
	if err != nil {
		t.Errorf("EncryptAES failed, expected [%s], got [%s]", encryptedValue, err)
	}

	if result != encryptedValue {
		t.Errorf("EncryptAES failed, expected [%s], got [%s]", encryptedValue, result)
	}
}

func TestDecryption(t *testing.T) {
	// Test for empty key
	_, err := DecryptAES("", value)
	if err == nil {
		t.Error("EncryptAES with empty key failed, expected error, got nil")
	}

	// Test for empty value
	_, err = DecryptAES(key, value)
	if err == nil {
		t.Error("EncryptAES with empty value failed, expected error, got nil")
	}
}

func TestAddPkcs7(t *testing.T) {
	// Test if given length is correct
	padded := AddPkcs7([]byte(value), aes.BlockSize)
	if len(padded) != aes.BlockSize {
		t.Errorf("AddPkcs7 failed, expected length of 16, got %d", len(padded))
	}
}

func TestRemovePkcs7(t *testing.T) {
	// Test if given length is correct
	unpadded := RemovePkcs7([]byte(value), aes.BlockSize)
	if len(unpadded) != len(value) {
		t.Errorf("RemovePkcs7 failed, expected length of %d, got %d", len(value), len(unpadded))
	}
}
