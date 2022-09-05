/*
 * @Author: 朱圣杰
 * @Date: 2022-09-05 19:39:10
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-05 19:39:11
 * @FilePath: /userCenter/postman_pre.js
 * @Description: postman的pre脚本
 * 
 */
test = pm.iterationData.toObject();
pm.environment.set("want_result", test.want_result);
pm.environment.set("want_err_code", test.want_err_code);