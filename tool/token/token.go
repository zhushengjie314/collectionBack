/*
 * @Author: 朱圣杰
 * @Date: 2022-09-02 16:28:53
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-02 16:31:14
 * @FilePath: /uploadTest/tool/token/token.go
 * @Description:
 *
 */
package token

import (
	"github.com/dgrijalva/jwt-go"
)

var key = []byte("Abfjds#$#$#%$BGFJYT&*^^%RfFfg")

type MyClaims struct {
	UserName string `json:"userName"`
	UserId   string `json:"userId"`
	Field    string `json:"field"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

/**
 * @description: 生成token
 * @param {*MyClaims} claims
 * @return {返回token、错误}
 */
func New(claims *MyClaims) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err = token.SignedString(key)
	return
}

/**
 * @description: 解析token
 * @param {string} token
 * @return {结构体}
 */
func Decode(token string) (*MyClaims, error) {
	claims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	return claims.Claims.(*MyClaims), err
}
