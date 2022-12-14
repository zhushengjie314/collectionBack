/*
 * @Author: 朱圣杰
 * @Date: 2022-08-28 21:37:21
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-05 19:19:38
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
	Err[1002] = &MyErr{1002, "任务名称为空"}
	Err[1003] = &MyErr{1003, "用户不存在"}
	Err[1004] = &MyErr{1004, "用户名/密码错误"}
	Err[1005] = &MyErr{1005, "日期早于今天"}
	Err[1006] = &MyErr{1006, "注册信息非法"}

	Err[1007] = &MyErr{1007, "没有token"}
	Err[1008] = &MyErr{1008, "token非法"}
	Err[1009] = &MyErr{1009, "token过期"}

	Err[1010] = &MyErr{1010, "参数解析失败"}

	//错误码20开头，代表后端业务代码错误
	Err[2001] = &MyErr{2001, "任务id生成失败"}
	Err[2002] = &MyErr{2002, "未知异常"}

	Err[2003] = &MyErr{2003, "更新token失败"}
	Err[2004] = &MyErr{2004, "生成token失败"}

	//错误码30开头，代表数据库错误
	Err[3001] = &MyErr{3001, "写入数据库失败"}
	Err[3002] = &MyErr{3002, "查询数据库失败"}
	Err[3003] = &MyErr{3003, "更新数据库失败"}
	Err[3004] = &MyErr{3004, "删除数据库失败"}
	Err[3005] = &MyErr{3005, "无记录"}

}
