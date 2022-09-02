/*
 * @Author: zhushengjie314 zhushengjie314@163.com
 * @Date: 2022-09-01 19:50:11
 * @LastEditors: zhushengjie314 zhushengjie314@163.com
 * @LastEditTime: 2022-09-01 20:50:59
 * @FilePath: /uploadTest/control/task/task.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package task

import (
	"fmt"
	"time"
	"uploadTest/service/task"

	"github.com/gin-gonic/gin"
)

func MainTask(c *gin.Context) {
	info := task.Info{}
	info.MainTaskCreator = "308752"
	info.MainCreateTime = time.Now()
	info.MainTaskName = "firstMainTask="
	info.MainTaskEndTime = time.Now().Add(3600 * 10)

	err := info.NewMain()
	if err != nil {
		fmt.Println("bad")
		return
	}
	fmt.Println("good")

}

func SubTask(c *gin.Context) {
	info := task.Sub{}
	info.MainTaskName = "firstMainTask="
	info.SubTaskName = "firstsubMainTask="
	info.FirstDataSource = "app"
	info.SecondDataSource = "ios"
	info.SubCreateTime = time.Now()

	err := info.NewSub()
	if err != nil {
		fmt.Println("bad")
		return
	}
	fmt.Println("good")

}
