package define

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cherry-neko-key"

var CodeLength = 4

// 单位s
var CodeExpire = 300
