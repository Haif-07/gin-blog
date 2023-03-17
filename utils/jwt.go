package utils

import (
	"gin-blog/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKeY = []byte("nihaoshijie")

type Cliams struct {
	SocialUserId string
	jwt.StandardClaims
}

func GetToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	cliams := &Cliams{
		SocialUserId: user.SocialUserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "ahong",
			Subject:   "user token",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := t.SignedString(jwtKeY)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenstring string) (*jwt.Token, *Cliams, error) {
	cliams := &Cliams{}
	token, err := jwt.ParseWithClaims(tokenstring, cliams, func(t *jwt.Token) (interface{}, error) {
		return jwtKeY, nil
	})
	return token, cliams, err
}
