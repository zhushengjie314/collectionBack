/*
 * @Author: 朱圣杰
 * @Date: 2022-09-03 14:46:21
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-03 14:46:21
 * @FilePath: /uploadTest/tool/time/time.go
 * @Description: 生成各种时间
 *
 */
package time

import "time"

/**
 * @description: 生成未来的时间戳
 * @param {int64} add
 * @return {*}
 */
func FutureTime(add int64) int64 {
	return time.Now().Unix() + add
}
