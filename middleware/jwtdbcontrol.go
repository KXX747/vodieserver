package middleware

import (
	"github.com/dgrijalva/jwt-go"
)

/**
 jwt做用户权限拦截验证
 */

const(
	signedKey ="www.erc20.com-jsonjdkfhsf?!@##%$^%EFSFDXCXCVHFHRHWE@#@$#%%^*&^(*&)*__#@#%$^%&^*&(*)*()__"

)

type Claims struct {
	Appid string `json:"appid"`
	jwt.StandardClaims
}


func CreateAuthTokenOf5Minute(key string )(token string) {
	//expireToken := time.Now().Add(time.Minute * 5).Unix()

	return CreateAuthToken(key)
}

/**
 加密key
 */
func CreateAuthToken(key string)(token string) {
	claims := Claims{
		Appid: key,
		StandardClaims: jwt.StandardClaims{
		//	ExpiresAt: d,
			Issuer:    key,
		},
	}

	//使用claim创建一个token
	c_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//signs the token with a secret
	signedToken, _ := c_token.SignedString([]byte(signedKey))
	return signedToken
}

/**
解密key
 */
func TokenAuth(token string)  (string, error) {

	authToken, err := jwt.ParseWithClaims(token,&Claims{},
	 func(token *jwt.Token) (interface{}, error) {
		 return []byte(signedKey), nil
	})
	if claims, ok := authToken.Claims.(*Claims); ok && authToken.Valid {
		return claims.Appid, err
	}
	return "", err

}