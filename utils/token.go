package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(ttl time.Duration, payload any, privateKey string) (string, error) {

	privKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf(" Could not decode key : %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(privKey)

	if err != nil {
		return "", fmt.Errorf("Error: Create Parse Key : %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		return "", fmt.Errorf("Create: Sign Token Error: %v", err)
	}

	return token, nil

}

func VerifyToken(token string, publicKey string) (any, error) {
	pbey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("Verify: Decoding Pub Key : %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(pbey)
	if err != nil {
		return nil, fmt.Errorf("Validate: parse key : %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); ok {
			return nil, fmt.Errorf("unexpected method %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {

		return nil, fmt.Errorf("Validate: %w", err)
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("Validate : Invalid Token")
	}

	return claims["sub"], nil

}
