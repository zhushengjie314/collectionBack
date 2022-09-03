/*
 * @Author: 朱圣杰
 * @Date: 2022-09-03 10:10:28
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-03 10:23:57
 * @FilePath: /uploadTest/tool/time/compare.go
 * @Description: 提供时间的比较
 *
 */
package time

import "time"

/**
 * @description: 判断输入时间和当前时间差是不是小于N秒
 * @param {time.Time} t
 * @param {float64} N
 * @return {*}
 */
func EarlyNowLestN(t time.Time, N float64) bool {
	gap := time.Since(t).Seconds()
	if gap > 0 && gap < N {
		return true
	} else {
		return false
	}
}

/**
 * @description: 根据时间搓判断输入时间和当前时间差是不是小于N秒
 * @param {int64} t
 * @param {int64} N
 * @return {*}
 */
func NumEarlyNowLestN(t int64, N int64) bool {
	if t-time.Now().Unix() < N {
		return true
	} else {
		return false
	}
}

/**
 * @description: 判断输入时间是不是在未来
 * @param {time.Time} t
 * @return {*}
 */
func BigThenNow(t time.Time) bool {
	gap := time.Since(t).Seconds()
	if gap < 0 {
		return true
	} else {
		return false
	}
}

/**
 * @description: 根据时间搓判断输入时间是不是在未来
 * @param {int64} t
 * @return {*}
 */
func NumBigThenNow(t int64) bool {
	gap := t - time.Now().Unix()
	if gap < 0 {
		return true
	} else {
		return false
	}
}
