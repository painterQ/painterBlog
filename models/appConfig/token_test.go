package appConfig

import (
	"encoding/hex"
	"fmt"
	"net"
	"testing"
)


func Test_checkToken(t *testing.T) {
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiOjoxIiwiZXhwIjoxNTgwMzAxNzgwLjA5MjkzMiwibmJmIjoxNTgwNDc0NTgwLjA5MjkzMn0._G6QmRL6F_o8hV8Th7QzSFfbaofsToVjjpHS52qgP0I`
	key := `a1c2a9df558cfa23133c818f53be0ac31c626a1977fd28f29595f04d09e4fd72`
	ipStr := `::1`
	ip := net.ParseIP(ipStr)
	keyBytes, _ := hex.DecodeString(key)
	got := CheckToken(token, keyBytes, &ip)
	fmt.Println(got)
}
