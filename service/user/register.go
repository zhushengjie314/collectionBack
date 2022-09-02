package user

import (
	"uploadTest/db"
	"uploadTest/db/mongo"
	merr "uploadTest/err"
	"uploadTest/log"
	"uploadTest/tool/check"
)

func (i *Info) Register() (err error) {
	err = i.checkRegister()
	if err != nil {
		return
	}
	err = db.Conn["appTest"].(*mongo.MongoDb).InserOne("user", i)
	return
}

/**
 * @description: 检查注册功能输入合法性
 * @return {*}
 */
func (i *Info) checkRegister() (err error) {
	// 检查输入
	if check.EmptyStr(i.Id) || check.EmptyStr(i.Name) || check.EmptyStr(i.Pass) {
		log.Error(i, "注册信息非法")
		err = merr.Err[1006]
		return
	}

	// 检查名字是否已被注册
	return
}
