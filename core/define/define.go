package define

import (
	"github.com/dgrijalva/jwt-go"
)

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

var TokenExpire = 36000
var RefreshTokenExpire = 72000

// 腾讯云
var CosBucket = "https://cherryneko-1312558494.cos.ap-nanjing.myqcloud.com"

// 查询参数
var PageSize = 10

var Datetime = "2006-01-02 15:04:05"
