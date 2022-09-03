/*
 * @Author: 朱圣杰
 * @Date: 2022-09-03 11:06:46
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-03 14:34:00
 * @FilePath: /uploadTest/db/mongo/search.go
 * @Description: 查询mongo
 *
 */
package mongo

import (
	merr "uploadTest/err"
	"uploadTest/log"
)

/**
 * @description: 查询一个结果
 * @param {Collection} coll
 * @param {map[string]interface{}} query bsonmap
 * @return {*}
 */
func (m *MongoDb) SearchOne(coll Collection, query map[string]interface{}) (res map[string]interface{}, err error) {
	c := m.DataBase.C(coll)
	err2 := c.Find(query).One(&res)
	if err2 != nil {
		log.Error(query, err2.Error())
		if err2.Error() == "not found" {
			err = merr.Err[3005]
		} else {
			err = merr.Err[3002]
		}
		return
	}
	if res == nil {
		log.Error(query, "返回空指针")
		err = merr.Err[3005]
	}
	return
}

func (m *MongoDb) Search(coll Collection) (res []interface{}, err error) {
	c := m.DataBase.C(coll)
	err2 := c.Find(nil).All(&res)
	if err2 != nil {
		log.Error(coll, err2.Error())
		if err2.Error() == "not found" {
			err = merr.Err[3005]
		} else {
			err = merr.Err[3002]
		}
	}
	if res == nil {
		log.Error(coll, "返回空指针")
		err = merr.Err[3005]
	}
	return
}
