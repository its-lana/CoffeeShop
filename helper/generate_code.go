package helper

import (
	"bytes"
	"crypto/rand"
)

const (
	charset     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	comboLength = 4
)

func GenerateSecretOrderCode() (string, error) {
	var buffer bytes.Buffer
	for i := 0; i < comboLength; i++ {
		b := make([]byte, 1)
		_, err := rand.Read(b)
		if err != nil {
			return "", err
		}
		index := int(b[0]) % len(charset) // Adjust for potential negative values
		buffer.WriteByte(charset[index])
	}

	code := buffer.String()
	return code, nil
}
