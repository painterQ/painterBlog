package internal

import (
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"net"
	"time"
)

const (
	expiry    = 172800 //48hour
)

//painterClaim custom claim
type painterClaim struct {
	IP string `json:"address"`
	jwt.StandardClaims
}

//GenToken json web token
func GenToken(ip net.IP, key []byte) string {
	now := time.Now()
	next := now.Add(expiry * time.Second)
	claims := &painterClaim{
		IP: ip.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(next),
			NotBefore: jwt.At(now),
		},
	}
	fmt.Println("###key",hex.EncodeToString(key))
	fmt.Println("###ip",ip.String())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(key)
	if err !=nil{
		panic(err)
	}
	return result
}

//CheckToken 校验token是否有效, 包括IP
func CheckToken(token string, key []byte, fromIP *net.IP) bool {
	tokenObject, err := jwt.ParseWithClaims(token, &painterClaim{},
		func(*jwt.Token) (interface{}, error) {
			return key, nil
		},
	)
	if err != nil {
		return false
	}
	claim := tokenObject.Claims.(*painterClaim)
	return fromIP.Equal(net.ParseIP(claim.IP))
}
