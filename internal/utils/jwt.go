package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type ClaimCustomJwt struct {
	Id uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uuid.UUID) (string, error) {

	godotenv.Load()
	mySecret := os.Getenv("SECRET_KEY")

	// claim isinya data-data yang disimpan di dalam JWT.
	// generate claim, claims = payload
	claims := ClaimCustomJwt{
		Id: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// token akan dikembalikan dengan newwithclaims
	// signingmethodHS256 = algoritma jwt
	// signedstring = signature, diberikan setelah token di generate
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(mySecret)) // memberikan ttd, secret key yang disimpan dari .env
	if err != nil {
		return "", err
	}
	// struktur jwt (Header.Payload.Signature)
	return tokenString, nil
}

func VerifyToken(tokenString string) (*ClaimCustomJwt, error) {

	godotenv.Load()
	mySecret := os.Getenv("SECRET_KEY")

	token, err := jwt.ParseWithClaims(
		tokenString,
		&ClaimCustomJwt{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(mySecret), nil
		},
	)

	// secret key harus byte slice
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*ClaimCustomJwt)
	if !ok {
		return nil, err
	}

	return claims, err
}
