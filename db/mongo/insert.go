/*
 * @Author: error: git config user.name && git config user.email & please set dead value or install git
 * @Date: 2022-08-29 20:04:34
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-08-29 20:44:38
 * @FilePath: /uploadTest/db/mongo/insert.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mongo

import (
	"context"

	"github.com/sirupsen/logrus"

	merr "uploadTest/err"
)

type error = *merr.MyErr

/**
 * @description: 插入数据
 * @param {Collection} coll
 * @param {interface{}} document
 * @return {*}
 */
func (m *MongoDb) InserOne(coll Collection, document interface{}) (err error) {
	c := m.DataBase.Collection(coll)
	id, e := c.InsertOne(context.TODO(), document)
	if e != nil {
		err = merr.Err[2002]
		logrus.Error(e)
		return
	} else {
		logrus.Infof("insert %s sunncess id %d", coll, id)
		return
	}
}
