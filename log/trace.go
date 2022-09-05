/*
 * @Author: 朱圣杰
 * @Date: 2022-09-05 16:12:10
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-05 16:17:45
 * @FilePath: /uploadTest/log/trace.go
 * @Description: 获取促发日志事件的文件以及行号
 *
 */
package log

import (
	"runtime"
)

/**
 * @description: 获取路径与行号
 * @return {*}
 */
func getTrace() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if ok {

		return file, line
	} else {
		return "", -1
	}
}
