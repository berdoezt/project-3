package helper

import (
	"project-tiga/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateID() string {
	return uuid.New().String()
}

func Hash(plain string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plain), 8)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func IsHashValid(hash string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))

	return err == nil
}

func GenerateAccessToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(1 * time.Minute).Unix(),
	}

	return generateToken(claims, "$uP3R___seC12E7")
}

func GenerateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
		"refresh": true,
	}

	return generateToken(claims, "$uP3R___jUUUN10Rrrr")
}

func generateToken(claims jwt.MapClaims, secret string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := jwtToken.SignedString([]byte(secret))
	return tokenString, err
}

func VerifyAccessToken(token string) (*jwt.Token, error) {

	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, model.ErrorInvalidToken
		}

		return []byte("$uP3R___seC12E7"), nil
	})

	return jwtToken, err
}

func VerifyRefreshToken(token string) (*jwt.Token, error) {

	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, model.ErrorInvalidToken
		}

		return []byte("$uP3R___jUUUN10Rrrr"), nil
	})

	return jwtToken, err
}
