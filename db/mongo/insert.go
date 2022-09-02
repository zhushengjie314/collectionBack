/*
 * @Author: error: git config user.name && git config user.email & please set dead value or install git
 * @Date: 2022-08-29 20:04:34
 * @LastEditors: zhushengjie314 zhushengjie314@163.com
 * @LastEditTime: 2022-09-02 11:37:53
 * @FilePath: /uploadTest/db/mongo/insert.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mongo

import (
	merr "uploadTest/err"
	"uploadTest/log"
)

type error = *merr.MyErr

/**
 * @description: 插入数据
 * @param {Collection} coll
 * @param {interface{}} document
 * @return {*}
 */
func (m *MongoDb) InserOne(coll Collection, document interface{}) (err error) {
	c := m.DataBase.C(coll)
	e := c.Insert(document)
	if e != nil {
		err = merr.Err[2002]
		log.Error(document, e.Error())
		return
	} else {
		return
	}

}
