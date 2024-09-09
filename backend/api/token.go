package secretnote

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) JWTMaker {
	return JWTMaker{secretKey}
}

func (JWTMaker *JWTMaker) CreateToken(email string, duration time.Duration) (string, error) {

	tokenID, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":        tokenID,
		"Email":     email,
		"IssuedAt":  time.Now(),
		"ExpiredAt": time.Now().Add(duration),
	})
	token, err := jwtToken.SignedString([]byte(JWTMaker.secretKey))
	return token, err
}

func (JWTMaker *JWTMaker) VerifyToken(tokenString string) (interface{}, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTMaker.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, err
	} else {
		return claims, nil
	}

}
