/*
 * @Author: 朱圣杰
 * @Date: 2022-09-02 13:49:44
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-03 15:04:08
 * @FilePath: /uploadTest/service/user/login.go
 * @Description: 用户登录功能
 *
 */
package user

import (
	"uploadTest/db"
	"uploadTest/db/mongo"
	merr "uploadTest/err"
	"uploadTest/log"
	"uploadTest/tool/time"
	token2 "uploadTest/tool/token"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2/bson"
)

/**
 * @description: 登录
 * @return {*}
 */
func (i *Info) Login() (token string, err error) {
	// 查user表

	query := bson.M{"id": i.Id}
	res, err := db.Conn["appTest"].(*mongo.MongoDb).SearchOne("user", query)
	if err != nil {
		return "", err
	}
	getInfo := Info{}
	err2 := mapstructure.Decode(res, &getInfo)
	if err2 != nil {
		//t.Error("空的")
		log.Error(res, err2.Error())
		err = merr.Err[1004]
	}
	if getInfo.Id != i.Id || getInfo.Pass != i.Pass {
		log.Error(i, "用户名/密码不正确")
		err = merr.Err[1004]
	}

	// 生成token
	token, err3 := token2.NewUserToken(i.Name, i.Id, time.FutureTime(3600*24*15))
	if err3 != nil {
		//t.Error("空的")
		log.Error(i, err3.Error())
		err = merr.Err[2004]
	}
	return
}
