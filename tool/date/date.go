/*
 * @Author: 朱圣杰
 * @Date: 2022-08-28 22:13:18
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-08-28 22:15:04
 * @FilePath: /uploadTest/tool/date/date.go
 * @Description: 时间和日期相关工具
 *
 * Copyright (c) 2022 by error: git config user.name && git config user.email & please set dead value or install git, All Rights Reserved.
 */
package date

import "time"

/**
 * @description: 判断时间是否比今天早
 * @param {time.Time} in
 * @return {*}
 */
func BeforeToday(in time.Time) bool {
	now := time.Now()
	if now.Year() > in.Year() {
		return true
	}
	if now.Month() > in.Month() {
		return true
	}
	if now.Day() > in.Day() {
		return true
	}
	return false
}
