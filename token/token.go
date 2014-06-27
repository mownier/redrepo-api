package token

import (
	"encoding/base64"
	"code.google.com/p/go-uuid/uuid"
	)
func GenerateAccessToken() string {
	accessToken := uuid.New()
	return base64.StdEncoding.EncodeToString([]byte(accessToken))
}

func GenerateRefreshToken()  string {
	refreshToken := uuid.New()
	return base64.StdEncoding.EncodeToString([]byte(refreshToken))
}

func GenerateClientKey() string {
	key := uuid.New()
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func GenerateClientSecret() string {
	secret := uuid.New()
	return base64.StdEncoding.EncodeToString([]byte(secret))
}