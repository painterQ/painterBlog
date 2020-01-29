package models

import (
	"github.com/dgrijalva/jwt-go/v4"
	"net"
	"net/http"
	"time"
)

const (
	expiry = 172800 //48hour
	tokenName = "painterBlog"
)

//painterClaim custom claim
type painterClaim struct {
	IP string `json:"address"`
	jwt.StandardClaims
}

//GenToken json web token
func genToken(ip net.IP, key []byte) string {
	now := time.Now()
	next := now.Add(expiry * time.Second)
	claims := &painterClaim{
		IP: ip.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(now),
			NotBefore: jwt.At(next),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, _ := token.SignedString(key)
	return result
}

//checkToken 校验token是否有效, 包括IP
func checkToken(token string, key []byte, fromIP *net.IP) bool {
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

//SetToken set token to client
func SetToken(token string, w http.ResponseWriter) string {
	cookie := http.Cookie{Name: tokenName, Value: token, Path: "/", MaxAge: int(expiry)}
	http.SetCookie(w, &cookie)
	return token
}

//ClearToken clear token
func ClearToken(w http.ResponseWriter) {
	cookie := http.Cookie{Name: tokenName, Value: "clear", Path: "/", MaxAge: 0}
	http.SetCookie(w, &cookie)
}
