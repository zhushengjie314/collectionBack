/*
 * @Author: error: git config user.name && git config user.email & please set dead value or install git
 * @Date: 2022-08-29 20:46:59
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-02 15:16:15
 * @FilePath: /uploadTest/config/config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

import (
	"fmt"
	"strconv"
	"uploadTest/db"
	"uploadTest/db/mongo"
	log "uploadTest/log"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var Setting *viper.Viper

/**
 * @description: 读取配置文件
 * @return {*}
 */
func init() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("读取配置文件错误", err.Error())
		panic("停止")
	}
	Setting = viper.GetViper()
}

func Init() {
	InitDb()
	InitLog()
}

func Quit() {
	QuitDb()
}

func InitDb() {
	logrus.Infof("start inindb")
	m := Setting.Get("db")
	dbArr := m.([]interface{})
	for _, val := range dbArr {
		paramMap := val.(map[string]interface{})
		if paramMap["type"].(string) == "mongo" {
			dbMongo := &mongo.MongoDb{}
			dbMongo.Name = paramMap["name"].(string)
			dbMongo.Ip = paramMap["host"].(string)
			dbMongo.Port = strconv.Itoa(paramMap["port"].(int))
			dbMongo.User = paramMap["user"].(string)
			dbMongo.Pass = strconv.Itoa(paramMap["pass"].(int))
			dbMongo.Base = paramMap["dataBase"].(string)
			dbMongo.Connect()
			db.Conn[dbMongo.Name] = dbMongo
		}

	}
}

func InitLog() {
	//logrus.AddHook(&log.ErrHook{})
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func QuitDb() {
	db.Conn["appTest"].Close()
	fmt.Println("关闭数据库")
}
