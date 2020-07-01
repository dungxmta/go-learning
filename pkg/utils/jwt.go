package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const (
	SECRET_KEY = "c3VwZXJfc2VjcmV0X2tleQo="
)

const (
	defaultExp = 24 // hour
	issuer     = "test"
)

type TokenType uint8

const (
	PortalToken TokenType = iota
	ApiKeyToken
)

var tokenTypes []string

func init() {
	tokenTypes = []string{"portal", "api_key"}
}

func (idx TokenType) String() string {
	return tokenTypes[idx]
}

type TimeUnit uint8

const (
	Second TimeUnit = iota + 1
	Minute
	Hour
	Day
	Week
	Month
	Year
)

func (idx TimeUnit) FromNow(duration time.Duration) (int64, error) {
	switch idx {
	case Second:
		return time.Now().Add(time.Second * duration).Unix(), nil
	case Minute:
		return time.Now().Add(time.Minute * duration).Unix(), nil
	case Hour:
		return time.Now().Add(time.Hour * duration).Unix(), nil
	case Day:
		return time.Now().Add(time.Hour * 24 * duration).Unix(), nil
	case Week:
		return time.Now().Add(time.Hour * 24 * 7 * duration).Unix(), nil
	case Month:
		return time.Now().Add(time.Hour * 24 * 7 * 30 * duration).Unix(), nil
	case Year:
		return time.Now().Add(time.Hour * 24 * 365 * duration).Unix(), nil
	default:
		return 0, errors.New("invalid unit")
	}
}

type ClaimsApiKey struct {
	KeyId     string      `json:"key_id"`
	TokenType string      `json:"typ"`
	Others    interface{} `json:"others,omitempty"`
	jwt.StandardClaims
}

func (m ClaimsApiKey) Valid() error {
	// TODO: do validate here, e.g. check perm, scope, ...
	return nil
}

type ClaimsPortal struct {
	UserId    string      `json:"user_id"`
	TenantId  string      `json:"tenant_id"`
	TokenType string      `json:"typ"`
	Others    interface{} `json:"others,omitempty"`
	jwt.StandardClaims
}

type TokenOpts struct {
	KeyId     string
	TokenType string

	Duration time.Duration
	Unit     TimeUnit
}

func GenToken(opts *TokenOpts) (token string, err error) {
	now := time.Now()

	claims := ClaimsApiKey{
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Hour * defaultExp).Unix(),
		},
	}

	if opts != nil {
		claims.KeyId = opts.KeyId
		claims.TokenType = opts.TokenType

		exp, err := opts.Unit.FromNow(opts.Duration)
		if err != nil {
			return "", err
		}
		claims.ExpiresAt = exp
	}

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tk.SignedString([]byte(SECRET_KEY))
	return
}

func IsValidToken(s string) (valid bool) {
	claims := &ClaimsApiKey{}

	// token, err := jwt.Parse(s, func(token *jwt.Token) (i interface{}, err error) {
	token, err := jwt.ParseWithClaims(s, claims, func(token *jwt.Token) (i interface{}, err error) {
		// check alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		log.Println("invalid token,", err)
		return false
	}

	// log.Println("claims?", token.Claims)
	valid = token.Valid
	log.Println("valid?", valid)
	if valid {
		b, _ := json.Marshal(claims)
		log.Println("claims?", string(b))
	}

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	// check perms...
	// 	log.Println("token ok,", claims)
	// 	return true
	// }

	return
}
