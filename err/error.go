/*
 * @Author: 朱圣杰
 * @Date: 2022-08-28 21:37:21
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-02 14:21:09
 * @FilePath: /uploadTest/err/error.go
 * @Description:统一的错误处理，err的key为错误码，前两位为错误类型，后两位为具体错误
 *
 * Copyright (c) 2022 by error: git config user.name && git config user.email & please set dead value or install git, All Rights Reserved.
 */
package err

type MyErr struct {
	Code int    `json:"errCode"`
	Msg  string `json:"errMsg`
}

var Err map[int]*MyErr

/**
 * @description: 初始化错误
 * @return {*}
 */
func init() {
	Err = make(map[int]*MyErr)

	//错误码10开头，代表前端传入参数错误
	Err[1001] = &MyErr{1001, "任务名称不符合格式"}
	Err[1002] = &MyErr{1001, "任务名称为空"}
	Err[1003] = &MyErr{1001, "用户不存在"}
	Err[1004] = &MyErr{1001, "用户密码错误"}
	Err[1005] = &MyErr{1001, "日期早于今天"}
	Err[1006] = &MyErr{1001, "注册信息非法"}

	//错误码20开头，代表后端业务代码错误
	Err[2001] = &MyErr{1001, "任务id生成失败"}
	Err[2002] = &MyErr{1001, "写入数据库失败"}

}
