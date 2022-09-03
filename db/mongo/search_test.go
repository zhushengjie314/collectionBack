/*
 * @Author: 朱圣杰
 * @Date: 2022-09-03 13:51:49
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-03 14:56:14
 * @FilePath: /uploadTest/db/mongo/search_test.go
 * @Description:
 *
 */
package mongo

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestSearchOne(t *testing.T) {
	db := MongoDb{
		Name: "appTest",
		Ip:   "127.0.0.1",
		Port: "27017",
		User: "root",
		Pass: "123456",
		Base: "appTest",
	}
	db.Connect()
	query := bson.M{"id": "308753"}
	res, err := db.SearchOne("user", query)

	if err != nil {
		t.Fail()
	}
	if res == nil {
		//t.Error("空的")
		t.Fail()
	}
	t.Logf("返回%v/n", res)

}

func TestSearch(t *testing.T) {
	db := MongoDb{
		Name: "appTest",
		Ip:   "127.0.0.1",
		Port: "27017",
		User: "root",
		Pass: "123456",
		Base: "appTest",
	}
	db.Connect()
	res, err := db.Search("taskTest")

	if err != nil {
		t.Fail()
	}
	if res == nil {
		//t.Error("空的")
		t.Fail()
	}
	t.Logf("返回%v/n", res)
	bson.Unmarshal()
}
