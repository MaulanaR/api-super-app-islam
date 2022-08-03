package app

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"

	"grest.dev/grest/crypto"
)

func configureCrypto() {
	crypto.Configure(CRYPTO_KEY)
}

func Encrypt(text string) (string, error) {
	encrypted, err := crypto.Encrypt(text)
	if err != nil {
		return encrypted, err
	}
	return CRYPTO_PREFIX + encrypted, nil
}

func Decrypt(text string) (string, error) {
	prefixLength := len([]rune(CRYPTO_PREFIX))
	if text[:prefixLength] != CRYPTO_PREFIX {
		return decryptOld(text)
	}
	return crypto.Decrypt(text[prefixLength:])
}

func decryptOld(text string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return text, err
	}

	h := sha256.New()
	h.Write([]byte(CRYPTO_KEY))
	key := h.Sum(nil)[:32]
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", crypto.ErrCtUnpadded
	}
	iv := []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = pkcs5Unpadding(ciphertext)
	return string(ciphertext), nil
}

func pkcs5Unpadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	length := len(encrypt) - int(padding)
	if length > 0 {
		return encrypt[:length]
	}
	return encrypt
}
