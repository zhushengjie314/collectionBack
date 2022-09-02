/*
 * @Author: 朱圣杰
 * @Date: 2022-07-26 10:36:29
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-02 14:41:54
 * @FilePath: /uploadTest/router/router.go
 * @Description:
 *
 * Copyright (c) 2022 by error: git config user.name && git config user.email & please set dead value or install git, All Rights Reserved.
 */
package router

import (
	"fmt"
	"uploadTest/control/task"
	"uploadTest/control/user"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.POST("/newMainTask", task.MainTask)
	r.POST("/newSubTask", task.SubTask)
	r.POST("/register", user.Register)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "https://pan.baidu.com/rest/2.0/xpan/file?method=list&access_token=121.243aaff966e07b0415b84ed5cbc230df.YaH98rpdBLTGv1DocL5UNviuOfneyHUB4dQiA-D.kBm_WQ",
		})
	})

	r.GET("/collection/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "collectionhello",
		})
	})

	r.POST("/collection/images", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		dst := "./" + file.Filename
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})

		//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
}
