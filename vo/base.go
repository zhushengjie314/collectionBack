/*
 * @Author: 朱圣杰
 * @Date: 2022-09-03 09:45:57
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-03 14:40:28
 * @FilePath: /uploadTest/vo/base.go
 * @Description: 定义基础的返回结构体
 *
 */
package vo

import (
	merr "uploadTest/err"
)

type Base struct {
	Code   int         `json:"code,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Result bool        `json:"result,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

/**
 * @description: 返回正确的结构体
 * @param {interface{}} data
 * @return {*}
 */
func NewOk(data interface{}) Base {
	return Base{
		Data:   data,
		Result: true,
	}
}

/**
 * @description: 返回错误的结构体
 * @param {int} code 错误码
 * @param {interface{}} data
 * @return {*}
 */
func NewErr(code int, data interface{}) Base {

	return Base{
		Code:   merr.Err[code].Code,
		Msg:    merr.Err[code].Msg,
		Result: false,
		Data:   data,
	}
}
