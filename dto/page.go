/*
 * @Author: zhushengjie314 zhushengjie314@163.com
 * @Date: 2022-09-01 20:58:41
 * @LastEditors: zhushengjie314 zhushengjie314@163.com
 * @LastEditTime: 2022-09-02 11:22:18
 * @FilePath: /uploadTest/dto/page.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dto

type Page struct {
	Start int `json:"start"`
	Size  int `json:"size"`
}
