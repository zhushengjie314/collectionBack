/*
 * @Author: error: git config user.name && git config user.email & please set dead value or install git
 * @Date: 2022-08-29 16:57:42
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-08-29 20:54:31
 * @FilePath: /uploadTest/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	//"uploadTest/db"
	//mongo "uploadTest/db/mongo"
	"uploadTest/config"
	m "uploadTest/router"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Init()
	// a := &mongo.MongoDb{
	// 	Name: "main",
	// 	Ip:   "127.0.0.1",
	// 	Port: 27017,
	// 	User: "root",
	// 	Pass: "123456",
	// }
	// a.Connect()
	// con := db.Conn["main"]
	// baidu := config.Setting.GetStringMap("baidu")
	// fmt.Println(baidu)
	// url := "https://pan.baidu.com/rest/2.0/xpan/file?method=list&access_token=121.243aaff966e07b0415b84ed5cbc230df.YaH98rpdBLTGv1DocL5UNviuOfneyHUB4dQiA-D.kBm_WQ"
	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Println(err)

	// }
	// body, err := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(body))

	r := gin.Default()
	m.Router(r)

	r.Run(":8079")

}
