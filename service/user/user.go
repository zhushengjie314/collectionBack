/*
 * @Author: 朱圣杰
 * @Date: 2022-09-02 13:48:28
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-03 11:10:04
 * @FilePath: /uploadTest/service/user/user.go
 * @Description: 定义用户中心相关的结构体
 *
 */
package user

import (
	"encoding/json"
	"io"
	merr "uploadTest/err"
	"uploadTest/log"
)

type Id = string
type Name = string
type Pass = string
type error = *merr.MyErr

type Info struct {
	Id   Id   `bson:"id" json:"id"`
	Name Name `bson:"name" json:"name"`
	Pass Pass `bson:"pass" json:"pass"`
}

func (i *Info) Decode(read io.ReadCloser) (err error) {
	e := json.NewDecoder(read).Decode(i)
	if e != nil {
		log.Error(i, e.Error())
		err = merr.Err[1001]
	}
	return
}
