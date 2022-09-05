/*
 * @Author: 朱圣杰
 * @Date: 2022-09-05 19:39:51
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-05 19:39:52
 * @FilePath: /userCenter/postman_text.js
 * @Description: login接口Tests
 * 
 */
var want_result = pm.environment.get("want_result")
var want_err_code = pm.environment.get("want_err_code")

pm.test("状态码200", function () {
    pm.response.to.have.status(200);
})
if (want_result) {
    pm.test("不返回错误", function () {
        pm.response.to.not.have.jsonBody('code');
        pm.response.to.not.have.jsonBody('msg');
    })
    pm.test("返回true的result", function () {
        pm.response.to.have.jsonBody('result');
        var jsonData = pm.response.json();
        pm.expect(jsonData.result).to.eql(true);
    })
    pm.test("返回data", function () {
        pm.response.to.have.jsonBody('data');
        var jsonData = pm.response.json();
        pm.expect(jsonData.data).to.eql("登录成功");
    })
} else {
    pm.test("返回code和msg", function () {
        pm.response.to.have.jsonBody('code');
        pm.response.to.have.jsonBody('msg');
    })
    pm.test("返回正确的code", function () {
        var jsonData = pm.response.json();
        pm.expect(jsonData.code).to.eql(want_err_code);
    })
    pm.test("返回false的result", function () {
        pm.response.to.have.jsonBody('result');
        var jsonData = pm.response.json();
        pm.expect(jsonData.result).to.eql(false);
    })
    pm.test("不返回data", function () {
        pm.response.to.not.have.jsonBody('data');
    })
}
