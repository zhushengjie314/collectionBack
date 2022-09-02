/*
 * @Author: 朱圣杰
 * @Date: 2022-08-28 21:21:04
 * @LastEditors: zhushengjie314 zhushengjie314@163.com
 * @LastEditTime: 2022-09-01 21:06:43
 * @FilePath: /uploadTest/service/task/task.go
 * @Description: 处理任务的新建
 *
 * Copyright (c) 2022 by error: git config user.name && git config user.email & please set dead value or install git, All Rights Reserved.
 */
package task

import (
	"io"
	"time"
	"uploadTest/db"
	"uploadTest/db/mongo"
	"uploadTest/dto"
	merr "uploadTest/err"
	"uploadTest/tool/date"
	"uploadTest/tool/id"

	"github.com/sirupsen/logrus"
)

type Name = string    // task name
type Id = string      // main task id
type Time = time.Time // task end time
type User = string    // id of users who tack part in this task
type Source = string
type error = *merr.MyErr

type Info struct {
	MainTaskName    Name  `json:"mainTaskName"`
	MainTaskId      Id    `json:"mainTaskId"`
	MainTaskEndTime Time  `json:"taskEndTime"`
	MainCreateTime  Time  `json:"-"`
	MainStartTime   Time  `json:"-"`
	MainStopTime    Time  `json:"-"`
	MainTaskCreator Name  `json:"creator"`
	SubTasks        []Sub `json:"subTasks"`
}

type Sub struct {
	MainTaskName     Name   `json:"mainTaskName"`
	SubTaskName      Name   `json:"subTaskName"`
	SubTaskId        Id     `json:"subTaskId"`
	FirstDataSource  Source `json:"firstDataSource"`
	SecondDataSource Source `json:"secondDataSorce"`
	SubCreateTime    Time   `json:"-"`
}

type ListMain struct {
	Page dto.Page
}

type ListSub struct {
	Page       dto.Page
	MainTaskId Id `json:"minTaskId"`
}

/**
 * @description: 新建任务会在info中新增任务id和创建时间
 * @return {*}
 */
func (i *Info) NewMain() (err error) {
	// 校验参数正确性
	if err = i.checkNewMain(); err != nil {
		return
	}

	// 生成主任务id
	i.MainTaskId = id.Get()

	//保存到mmongoDb
	mdb := db.Conn["appTest"].(*mongo.MongoDb)
	print("save")
	err = mdb.InserOne("task", i)
	print(err)
	return
}

func (i *Info) ListMain() (err error) {
	// 校验参数正确性
	if err = i.checkNewMain(); err != nil {
		return
	}

	// 生成主任务id
	i.MainTaskId = id.Get()

	//保存到mmongoDb
	mdb := db.Conn["appTest"].(*mongo.MongoDb)
	print("save")
	err = mdb.InserOne("task", i)
	print(err)
	return
}

/**
 * @description: 检查新建主任务参数的合法性
 * @return {错误}
 */
func (i *Info) checkNewMain() (err error) {
	if i.MainTaskName == "" {
		logrus.Errorf("主任务名字为空")
		err = merr.Err[1001]
		return
	}
	if date.BeforeToday(i.MainTaskEndTime) {
		logrus.Errorf("主任务结束时间非法")
		err = merr.Err[1005]
		return
	}
	return
}

/**
 * @description: 检查新建子任务参数的合法性
 * @return {错误}
 */
func (s *Sub) checkNewSub() (err error) {
	if s.MainTaskName == "" {
		logrus.Errorf("name emepty")
		err = merr.Err[1001]
		return
	}
	if s.SubTaskName == "" {
		logrus.Errorf("name emepty")
		err = merr.Err[1001]
		return
	}
	if s.FirstDataSource == "" {
		logrus.Errorf("source emepty")
		err = merr.Err[1001]
		return
	}
	if s.SecondDataSource == "" {
		logrus.Errorf("source emepty")
		err = merr.Err[1001]
		return
	}
	return
}

func (i *Info) Transform(body io.ReadCloser) (err error) {

	return
}

/**
 * @description: 新建任务会在info中新增任务id和创建时间
 * @return {*}
 */
func (s *Sub) NewSub() (err error) {
	// 校验参数正确性
	if err = s.checkNewSub(); err != nil {
		return
	}

	// 生成主任务id
	s.SubTaskId = id.Get()

	//更新到mmongoDb
	mdb := db.Conn["appTest"].(*mongo.MongoDb)
	print("save")
	err = mdb.InserOne("task", s)
	print(err)
	return
}

func (lm *ListMain) List() (data []Info, err error) {
	// 查询mogo
	return
}

func (ls *ListSub) List() (data []Sub, err error) {
	// 查询mogo
	return
}
