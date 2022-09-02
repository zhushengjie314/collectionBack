/*
 * @Author: 朱圣杰
 * @Date: 2022-08-31 22:01:06
 * @LastEditors: zhushengjie314 zhushengjie314@163.com
 * @LastEditTime: 2022-09-02 12:06:46
 * @FilePath: /uploadTest/db/mongo/mongo.go
 * @Description:
 *
 * Copyright (c) 2022 by error: git config user.name && git config user.email & please set dead value or install git, All Rights Reserved.
 */
package mongo

import (
	"time"
	"uploadTest/db"

	"uploadTest/log"

	"gopkg.in/mgo.v2"
)

type Collection = string

type MongoDb struct {
	Name     db.Name
	Ip       db.Ip
	Port     db.Port
	Pass     db.Pass
	User     db.User
	Base     db.Base
	DataBase *mgo.Database
	Session  *mgo.Session
}

// func (m *MongoDb) Connect2() {
// 	uri := fmt.Sprintf("mongodb://%s:%s@%s:%v/?authSource=admin", m.User, m.Pass, m.Ip, m.Port)
// 	fmt.Println(uri)
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
// 	if err != nil {
// 		panic(err)
// 	}

// 	//fmt.Printf("%v", client)
// 	m.Client = client
// 	m.DataBase = m.Client.Database(m.Base)
// 	if _, ok := db.Conn[m.Name]; !ok {
// 		db.Conn[m.Name] = m
// 	} else {
// 		panic(errors.New("插入数据库缓存池失败"))
// 	}
// 	// defer func() {
// 	// 	logrus.Info("取消连接")
// 	// 	if err = client.Disconnect(context.TODO()); err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// }()
// }

func (m *MongoDb) Connect() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{m.Ip + ":" + m.Port},
		Timeout:  60 * time.Second,
		Username: m.User,
		Password: m.Pass,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatal(mongoDBDialInfo, err.Error())
		return
	}
	session.SetMode(mgo.Monotonic, true)
	m.Session = session
	m.DataBase = m.Session.DB(m.Base)
}

func (m *MongoDb) Type() string {
	return "mongo"
}

func (m *MongoDb) Close() {
	m.Session.Close()
	log.Info(m.Name, "数据库关闭")
}
