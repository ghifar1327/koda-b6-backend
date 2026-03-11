package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateSecureOTP() (int, error) {
    n, err := rand.Int(rand.Reader, big.NewInt(9000))
    if err != nil {
        return 0, err
    }
    return int(n.Int64()) + 1000, nil
}