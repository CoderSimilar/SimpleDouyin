package module

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const CurUserId = "uId"

// TokenExpireDuration token有效期
const TokenExpireDuration = time.Hour * 24 * 365

var mySecret = []byte("this is douyin")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录name和password字段，所以需要自定义结构体
type MyClaims struct {
	UserId   int64  `json:"userid"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenToken 生成 JWT
func GenToken(userId int64, password string) (string, error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		userId,
		password, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "groupwork",                                // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySecret)
}

// ParseToken 解析 JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid { // 检验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
