/*
 * @Author: 朱圣杰
 * @Date: 2022-08-29 16:41:21
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-08-29 16:57:16
 * @FilePath: /uploadTest/tool/id/id.go
 * @Description: 生成唯一id的封装
 */
package id

import (
	uuid "github.com/satori/go.uuid"
)

func Get() (uid string) {
	u := uuid.NewV4()
	return u.String()
}
