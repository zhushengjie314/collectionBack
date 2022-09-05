/*
 * @Author: 朱圣杰
 * @Date: 2022-09-03 09:26:09
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-05 20:35:55
 * @FilePath: /uploadTest/middleware/auth.go
 * @Description: token校验和刷新
 *
 */
package middleware

import (
	"net/http"
	merr "uploadTest/err"
	"uploadTest/log"
	"uploadTest/tool/time"
	"uploadTest/tool/token"
	"uploadTest/vo"

	"github.com/gin-gonic/gin"
)

type error = *merr.MyErr

type Info struct {
	NewToken string // 如果token快过期，则生成newToken
	token.MyClaims
}

/**
 * @description: token中间件
 * @param {*gin.Context} c
 * @return {*}
 */
func Auth(c *gin.Context) {
	userToken, err := c.Cookie("token")
	if err != nil {
		log.Error(c.Request.URL, err.Error())
		c.AsciiJSON(http.StatusOK, vo.NewErr(1007, nil))
		c.Abort()
		return
	}
	newInfo, merr := tokenValid(userToken)
	if merr != nil {
		c.AsciiJSON(http.StatusOK, vo.NewErr(merr.Code, nil))
		c.Abort()
		return
	}

	c.Set(userToken, newInfo)

	return

	data, err2 := token.Decode(userToken)
	if err2 != nil {
		log.Error(userToken, err.Error())
		c.AsciiJSON(http.StatusOK, vo.NewErr(1008, nil))
		c.Abort()
		return
	}
	if time.NumBigThenNow(data.ExpiresAt) {
		log.Error(userToken, "token过期")
		c.AsciiJSON(http.StatusOK, vo.NewErr(1009, nil))
		c.Abort()
		return
	}
	if time.NumEarlyNowLestN(data.ExpiresAt, 3600) {
		data.ExpiresAt += 3600 * 24 * 15
		tokenString, err3 := token.New(data)
		if err3 != nil {
			log.Error(data, err3.Error())
			c.AsciiJSON(http.StatusOK, vo.NewErr(2003, nil))
			c.Abort()
			return
		} else {
			newInfo := Info{
				NewToken: tokenString,
				MyClaims: *data,
			}
			c.Set(userToken, newInfo)
		}
	} else {
		newInfo := Info{
			NewToken: "",
			MyClaims: *data,
		}
		c.Set(userToken, newInfo)
	}
}

/**
 * @description: 验证token的合法性
 * @param {string} userToken
 * @return {*}
 */
func tokenValid(userToken string) (newInfo Info, err error) {
	data, err2 := token.Decode(userToken)
	if err2 != nil {
		log.Error(userToken, err2.Error())
		err = merr.Err[1008]
		return
	}
	if time.NumBigThenNow(data.ExpiresAt) {
		log.Error(userToken, "token过期")
		err = merr.Err[1009]
		return
	}
	if time.NumEarlyNowLestN(data.ExpiresAt, 3600) {
		data.ExpiresAt += 3600 * 24 * 15
		tokenString, err3 := token.New(data)
		if err3 != nil {
			log.Error(data, err3.Error())
			err = merr.Err[2003]
			return
		} else {
			newInfo = Info{
				NewToken: tokenString,
				MyClaims: *data,
			}
			return
		}
	} else {
		newInfo = Info{
			NewToken: "",
			MyClaims: *data,
		}
		return
	}
}

/**
 * @description: 测试中间件的使用
 * @param {*gin.Context} c
 * @return {*}
 */
func Test(c *gin.Context) {
	c.AsciiJSON(http.StatusOK, "hello")
}
