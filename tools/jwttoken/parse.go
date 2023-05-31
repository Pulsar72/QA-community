package jwttoken

import (
	"QA_community/model"
	"QA_community/tools/middleware"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

// 解析Token
func ParseToken(tokenString string) (*model.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return middleware.SecretSignKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}
