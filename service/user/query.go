/*
 * @Author: 朱圣杰
 * @Date: 2022-09-05 20:44:11
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-05 21:08:39
 * @FilePath: /uploadTest/service/user/query.go
 * @Description: 构建查询user表的api
 *
 */
package user

import (
	"uploadTest/config"
	"uploadTest/db"
	"uploadTest/db/mongo"
	merr "uploadTest/err"
	"uploadTest/log"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2/bson"
)

func IdQuery(id string) (info Info, err error) {
	query := bson.M{"id": id}
	mode := config.Setting.Get("mode")
	coll := "dev"
	if mode == "dev" {
		coll = "userDev"
	} else if mode == "test" {
		coll = "userTest"
	} else if mode == "online" {
		coll = "user"
	}
	res, err := db.Conn["appTest"].(*mongo.MongoDb).SearchOne(coll, query)
	if err != nil {
		return
	}
	err2 := mapstructure.Decode(res, &info)
	if err2 != nil {
		log.Error(res, err2.Error())
		err = merr.Err[1004]
		return
	}
	return

}
