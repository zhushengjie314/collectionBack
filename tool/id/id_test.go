/*
 * @Author: error: git config user.name && git config user.email & please set dead value or install git
 * @Date: 2022-08-29 16:58:45
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-08-29 19:59:09
 * @FilePath: /uploadTest/tool/id/id_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package id

import (
	"testing"
)

/**
 * @description: 测试生成uuid的程序，生成10000个不产生冲突
 * @param {*testing.T} t
 * @return {*}
 */
func TestGet(t *testing.T) {
	samples := make(map[string]int)
	count := 10000
	for i := 0; i < count; i++ {
		id := Get()
		if val, ok := samples[id]; ok {
			t.Log(val, i)
			t.Fail()
		} else {
			samples[id] = i
		}
	}
}
