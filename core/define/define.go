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

var TokenExpire = 3600
var RefreshTokenExpire = 7200

// 邮箱
var MailPWD = "AQSMQZUTGAEFGSDH"

// 腾讯云
var CosBucket = "https://cherryneko-1312558494.cos.ap-nanjing.myqcloud.com"
var TencentSecretID = "AKIDgdUDi03sTpzvd8RWNenncFgvyohjhYPB"
var TencentSecretKEY = "iHhxSqmV1KQJyQl04va2HJ8mCPPCl7FA"
