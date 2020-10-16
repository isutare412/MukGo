package api

import (
	"crypto/rsa"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// GoogleClaim struct for parse jwt
type GoogleClaim struct {
	jwt.StandardClaims
	Name string `json:"name"`
}

// JWKKeys used for parse public key
type JWKKeys struct {
	Keys []JWKKey `json:"keys"`
}

// JWKKey used for entries of JWTKeys
type JWKKey struct {
	Kid string `json:"kid"`
	N   string `json:"n"`
	E   string `json:"e"`
}

// CheckAuth read token from Authorization header and get user claim with that token
func CheckAuth(r *http.Request) (*GoogleClaim, error) {
	header := r.Header.Get("Authorization")
	tokens := strings.Split(header, " ")
	return ParseJWT(tokens[1])
}

// ParseJWT parse token to GoogleClaim struct
func ParseJWT(token string) (*GoogleClaim, error) {
	claim := &GoogleClaim{}
	// parse token string with custom claim
	_, err := jwt.ParseWithClaims(token, claim, func(parsedToken *jwt.Token) (interface{}, error) {
		kid := fmt.Sprintf("%v", parsedToken.Header["kid"])
		key, err := GetPEM("https://www.googleapis.com/oauth2/v3/certs", kid)
		return key, err
	})

	return claim, err
}

// GetPEM get pem of openid from url
func GetPEM(url string, kid string) (*rsa.PublicKey, error) {
	keys := &JWKKeys{}
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("on GetPEM: %v", err)
	}

	if err := json.NewDecoder(r.Body).Decode(keys); err != nil {
		return nil, fmt.Errorf("on GetPEM: %v", err)
	}

	keysArray := keys.Keys
	key := &JWKKey{}
	for i := 0; i < len(keysArray); i++ {
		if keysArray[i].Kid == kid {
			key = &keysArray[i]
			break
		}
	}

	publicKey := &rsa.PublicKey{}
	publicKey.E = int(binary.BigEndian.Uint32([]byte(key.E)))
	n := new(big.Int)
	n.SetBytes([]byte(key.N))
	publicKey.N = n

	return publicKey, nil
}
