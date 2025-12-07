package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Claims struct {
	UserId int64 `json:"userId"`
	jwt.RegisteredClaims
}

//生成JWT
func GenerateJWT(userId int64)(string, error){
	claims := Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Hour)),
			Issuer: "Maomao",
		},
	}

	//创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	//使用密钥签名token
	signedToken, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}


//解析和验证JWT
func ParseJWT(tokenStr string)(*Claims, error){
	//解析token
	var err error
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		//确保使用的是HS256签名算法
		if t.Method!= jwt.SigningMethodHS256 {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(viper.GetString("jwt.secret")), nil
	})
	
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}