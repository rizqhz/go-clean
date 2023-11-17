package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenInterface interface {
	GenerateToken(id uint) *string
	ExtractToken(token *jwt.Token) (*int, error)
}

type JwtToken struct {
	SignKey    string
	RefreshKey string
}

func NewJwt(signkey, refreshkey string) JwtTokenInterface {
	return &JwtToken{
		SignKey:    signkey,
		RefreshKey: refreshkey,
	}
}

func (j *JwtToken) GenerateToken(id uint) *string {
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := sign.SignedString([]byte(j.SignKey))
	if err != nil {
		return nil
	}
	return &token
}

func (j *JwtToken) ExtractToken(token *jwt.Token) (*int, error) {
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims := token.Claims
	exp, _ := claims.GetExpirationTime()
	if exp.Time.Compare(time.Now()) > 0 {
		claims := claims.(jwt.MapClaims)
		id := claims["id"].(int)
		return &id, nil
	}
	return nil, errors.New("expired token")
}
