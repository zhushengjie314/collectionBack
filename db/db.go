/*
 * @Author: zhushengjie314 zhushengjie314@163.com
 * @Date: 2022-09-01 20:06:38
 * @LastEditors: zhushengjie314 zhushengjie314@163.com
 * @LastEditTime: 2022-09-01 20:06:39
 * @FilePath: /uploadTest/db/db.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package db

import (
	"fmt"
)

type Name = string
type Ip = string
type Port = string
type Pass = string
type User = string
type Base = string

type Db interface {
	Type() string
	Connect()
	Close()
}

type DbMap map[Name]Db

var Conn DbMap

func init() {
	Conn = make(map[Name]Db)
	fmt.Println(Conn)
}
