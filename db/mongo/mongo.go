/*
 * @Author: 朱圣杰
 * @Date: 2022-08-31 22:01:06
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-08-31 23:54:23
 * @FilePath: /uploadTest/db/mongo/mongo.go
 * @Description:
 *
 * Copyright (c) 2022 by error: git config user.name && git config user.email & please set dead value or install git, All Rights Reserved.
 */
package mongo

import (
	"context"
	"errors"
	"fmt"
	"uploadTest/db"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection = string

type MongoDb struct {
	Name     db.Name
	Ip       db.Ip
	Port     db.Port
	Pass     db.Pass
	User     db.User
	Base     db.Base
	Client   *mongo.Client
	DataBase *mongo.Database
}

func (m *MongoDb) Connect() {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%v", m.Name, m.Pass, m.Ip, m.Port)
	fmt.Println(uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%v", client)
	m.Client = client
	m.DataBase = m.Client.Database(m.Base)
	if _, ok := db.Conn[m.Name]; !ok {
		db.Conn[m.Name] = m
	} else {
		panic(errors.New("插入数据库缓存池失败"))
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

}

func (m *MongoDb) Type() string {
	return "mongo"
}
