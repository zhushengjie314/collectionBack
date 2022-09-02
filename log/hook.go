/*
 * @Author: 朱圣杰
 * @Date: 2022-09-02 14:11:13
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-02 15:14:36
 * @FilePath: /uploadTest/log/hook.go
 * @Description: 日志存储
 *
 */
package log

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type ErrHook struct {
}

func (m *ErrHook) Levels() []log.Level {
	return []log.Level{
		log.ErrorLevel,
		log.PanicLevel,
	}
}

func (m *ErrHook) Fire(entry *log.Entry) error {
	f, err := os.OpenFile("err.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// var result bytes.Buffer
	// //encoder为实际转换者 gob.NewEncoder可以理解为制造一个转换者
	// encoder := gob.NewEncoder(&result)
	// //将结构体b转换
	// err = encoder.Encode(&entry)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("res", result)
	// outData, err := json.Marshal(&entry)
	// fmt.Printf("entry,%v", bytes.NewReader(outData))
	// if err != nil {
	// 	return err
	// }

	if _, err := f.Write([]byte(entry.Message)); err != nil {
		return err
	}
	return nil
}
