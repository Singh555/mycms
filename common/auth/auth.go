package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var jwtKey = []byte("supersecretkeyvdjwbdhwjdbiwuhdqwihdiq")

type JWTClaim struct {
	Id     uint   `json:"id"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(id uint, mobile string, email string) (tokenString string, err error) {
	expTime := viper.Get("JWT_EXP_HR")
	fmt.Println("jwt expiration time from env ", expTime)
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Id:     id,
		Email:  email,
		Mobile: mobile,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "",
			Audience:  nil,
			ExpiresAt: expirationTime.Date(),
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

var UserJwtData *JWTClaim

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		log.Fatal("error occurred during parsing the token", err.Error())
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		log.Fatal("error occurred during parsing the token", err.Error())
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		log.Warning("error token expired", err.Error())
		return
	}
	if token.Valid {
		UserJwtData = claims
	}

	return
}
