/*
 * @Author: 朱圣杰
 * @Date: 2022-09-02 14:39:49
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-02 14:46:11
 * @FilePath: /uploadTest/control/user/register.go
 * @Description: 注册的控制层
 *
 */
package user

import (
	"uploadTest/log"
	"uploadTest/service/user"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	body := &user.Info{}
	body.Decode(c.Request.Body)
	body.Register()

	defer func() {
		log.Info("", "xxxxxxxxxxxxxxxxxxxxxxxx")
		recover()
	}()

}
