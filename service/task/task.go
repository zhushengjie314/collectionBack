/*
 * @Author: 朱圣杰
 * @Date: 2022-08-28 21:21:04
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-08-29 20:03:17
 * @FilePath: /uploadTest/service/task/task.go
 * @Description: 处理任务的新建
 *
 * Copyright (c) 2022 by error: git config user.name && git config user.email & please set dead value or install git, All Rights Reserved.
 */
package task

import (
	"time"
	merr "uploadTest/err"
	"uploadTest/tool/date"
	"uploadTest/tool/id"
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
	SubTaskName      Name   `json:"subTaskName"`
	SubTaskId        Id     `json:"subTaskId"`
	FirstDataSource  Source `json:"firstDataSource"`
	SecondDataSource Source `json:"secondDataSorce"`
	SubCreateTime    Time   `json:"-"`
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

	// 保存到mmongoDb

	return
}

/**
 * @description: 检查新建主任务参数的合法性
 * @return {错误}
 */
func (i *Info) checkNewMain() (err error) {
	if i.MainTaskName == "" {
		err = merr.Err[1001]
		return
	}
	if date.BeforeToday(i.MainTaskEndTime) {
		err = merr.Err[1005]
		return
	}
	return
}
