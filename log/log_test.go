/*
 * @Author: zhushengjie314 zhushengjie314@163.com
 * @Date: 2022-09-02 10:56:46
 * @LastEditors: zhushengjie314 zhushengjie314@163.com
 * @LastEditTime: 2022-09-02 11:24:40
 * @FilePath: /uploadTest/log/log_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package log

import "testing"

type hello struct {
	A string
	B int
	c float64
}

func TestInfo(t *testing.T) {
	in := &hello{
		A: "good",
		B: 10,
		c: 4.4,
	}
	Info(in, "a")
}

func TestError(t *testing.T) {
	in := &hello{
		A: "good",
		B: 10,
		c: 4.4,
	}
	Error(in, "a")
}

func TestFailt(t *testing.T) {
	in := &hello{
		A: "good",
		B: 10,
		c: 4.4,
	}
	Fatal(in, "a")
}
