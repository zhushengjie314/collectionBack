/*
 * @Author: zhushengjie314 zhushengjie314@163.com
 * @Date: 2022-09-02 12:03:20
 * @LastEditors: zhushengjie314 zhushengjie314@163.com
 * @LastEditTime: 2022-09-02 12:03:20
 * @FilePath: /uploadTest/db/mongo/insert_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mongo

import (
	"testing"
	"time"
)

type document struct {
	StringData string
	IntData    int
	FloatData  float64
	TimeData   time.Time
}

func TestInsertOne(t *testing.T) {
	d := document{
		StringData: "中文字符&&！@！",
		IntData:    65536,
		FloatData:  3.14444,
		TimeData:   time.Now(),
	}
	db := MongoDb{
		Name: "appTest",
		Ip:   "127.0.0.1",
		Port: "27017",
		User: "root",
		Pass: "123456",
		Base: "appTest",
	}
	db.Connect()
	err := db.InserOne("taskTest", d)
	if err != nil {
		t.Fail()
	}

}
