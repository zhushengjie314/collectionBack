/*
 * @Author: zhushengjie314 zhushengjie314@163.com
 * @Date: 2022-09-02 11:31:52
 * @LastEditors: zhushengjie314 zhushengjie314@163.com
 * @LastEditTime: 2022-09-02 11:44:01
 * @FilePath: /uploadTest/db/mongo/mongo_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mongo

import "testing"

func TestConnect(t *testing.T) {
	db := MongoDb{
		Name: "appTest",
		Ip:   "127.0.0.1",
		Port: "27017",
		User: "root",
		Pass: "123456",
		Base: "appTest",
	}
	db.Connect()
}
