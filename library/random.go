package library

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateRandomSymbols(length int) (string, error) {
	symbols := "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"

	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return "", err
		}
		result[i] = symbols[index.Int64()]
	}

	return string(result), nil
}

func GenerateRandom(length int) (string, error) {
	symbols := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return "", err
		}
		result[i] = symbols[index.Int64()]
	}

	return string(result), nil
}
