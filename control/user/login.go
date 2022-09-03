/*
 * @Author: 朱圣杰
 * @Date: 2022-09-03 11:18:29
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-03 15:08:44
 * @FilePath: /uploadTest/control/user/login.go
 * @Description: 登录的逻辑
 *
 */
package user

import (
	"net/http"
	"uploadTest/service/user"
	"uploadTest/vo"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// token, err := c.Cookie("token")
	// if err != nil {
	// 	log.Error(c.Request.URL, err.Error())
	// 	c.AsciiJSON(http.StatusOK, vo.NewErr(1007, nil))
	// 	return
	// }
	// data, exists := c.Get(token)
	// if !exists {
	// 	log.Error(c.Request.URL, err.Error())
	// 	c.AsciiJSON(http.StatusOK, vo.NewErr(1007, nil))
	// 	return
	// }
	body := &user.Info{}
	body.Decode(c.Request.Body)
	token, err := body.Login()
	if err != nil {
		c.AsciiJSON(http.StatusOK, vo.NewErr(err.Code, nil))
	} else {
		c.SetCookie("token", token, 0, "", "", false, false)
		c.AsciiJSON(http.StatusOK, vo.NewOk("登录成功"))
	}
	// defer func() {
	// 	if err == nil {
	// 		c.AsciiJSON(http.StatusOK, vo.NewErr(2002, nil))
	// 	}
	// 	log.Error(c, "未知异常")
	// 	recover()
	// }()

}
