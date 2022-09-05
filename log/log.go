/*
 * @Author: zhushengjie314 zhushengjie314@163.com
 * @Date: 2022-09-02 10:49:11
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-05 16:19:48
 * @FilePath: /uploadTest/log/log.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package log

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func Info(field interface{}, message string) {

	// t := reflect.TypeOf(field)
	// type v reflect.Value
	// if t.Kind() == reflect.Ptr {
	// 	t = t.Elem()
	// 	v = v(reflect.ValueOf()
	// }
	// if t.Kind() != reflect.Struct {
	// 	log.Info("Check type error not Struct")
	// 	return
	// }

	// fieldNum := t.NumField()
	// result := make(map[string]reflect.Value)
	// for i := 0; i < fieldNum; i++ {
	// 	field := t.Field(i)

	// 	result[field.Name] = v.FieldByName(field.Name)
	// }

	log.WithFields(log.Fields{
		"data": field,
	}).Info(message)
}

func Error(field interface{}, message string) {
	file, line := getTrace()
	log.WithFields(log.Fields{
		"data": field,
		"file": fmt.Sprintf("%s:%d", file, line),
	}).Error(message)
}

func Fatal(field interface{}, message string) {
	file, line := getTrace()
	log.WithFields(log.Fields{
		"data": field,
		"file": fmt.Sprintf("%s:%d", file, line),
	}).Fatal(message)
}
