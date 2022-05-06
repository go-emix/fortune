package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"strings"
	"time"
)

var TokenMalformedErr = errors.New("token malformed error")
var TokenExpiredErr = errors.New("token expired")
var TokenNotValidYetErr = errors.New("token not valid yet")
var TokenInvalidErr = errors.New("token invalid")

const saltSource string = "adcdefghjkmnpqrstuvwxyz1234567890ABCDEFGHJKLMNPQRSTUVWXYZ9876543210"

var signKey = ""
var expire time.Duration

var slen = len(saltSource)

func randGenSalt() string {
	builder := strings.Builder{}
	for i := 0; i < 4; i++ {
		intn := rand.Intn(slen)
		builder.WriteByte(saltSource[intn])
	}
	return builder.String()
}

func Initialize(key string, ex int) {
	signKey = key
	expire = time.Duration(ex) * time.Second
	rand.Seed(time.Now().UnixNano())
}

//认证对象
type Certification struct {
	jwt.StandardClaims
	// 自定义字段，添加自己需要的信息
	UserId   int
	Username string
	//盐，随机数
	Salt string
}

func GenToken(claim Certification) (string, error) {
	claim.Salt = randGenSalt()
	claim.ExpiresAt = time.Now().Add(expire).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(signKey))
}

func ParseToken(token string) (*Certification, error) {
	parse, err := jwt.ParseWithClaims(token, &Certification{}, keyFunc)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			e := ve.Errors
			if e&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformedErr
			} else if e&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpiredErr
			} else if e&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYetErr
			} else {
				return nil, TokenInvalidErr
			}
		}
		return nil, TokenInvalidErr
	}
	if claims, ok := parse.Claims.(*Certification); ok && parse.Valid {
		return claims, nil
	}
	return nil, TokenInvalidErr
}

func keyFunc(*jwt.Token) (interface{}, error) {
	return []byte(signKey), nil
}
