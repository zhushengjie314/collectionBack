/*
 * @Author: 朱圣杰
 * @Date: 2022-09-03 15:16:32
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-05 20:24:28
 * @FilePath: /uploadTest/tool/token/token_test.go
 * @Description:
 *
 */
package token

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestUserToken(t *testing.T) {
	_, err := NewUserToken("朱圣杰", "308752", 3434324)
	if err != nil {
		t.Errorf("生成token错误%v\n", err)
		t.Fatal()
	}
}

func TestNew(t *testing.T) {
	m := MyClaims{
		UserName: "111",
		UserId:   "1111",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix(),
			Issuer:    "dataMange",
		},
	}
	tokenString, err := New(&m)
	if err != nil {
		t.Errorf("生成token错误%v\n", err)
	} else {
		t.Log(tokenString)
	}
}

func TestMain(m *testing.M) {
	setup()
	run := m.Run()
	tearDown()
	os.Exit(run)
}

var genToken map[string]MyClaims

func setup() {
	fmt.Println("start")
}

func tearDown() {
	fmt.Println("over")
}

/**
 * @description: 生成各种token
 * @return {*}
 */
func initGenToken() {
	genToken = make(map[string]MyClaims)
	// 过期token
	genToken["overdue"] = MyClaims{
		UserName: "朱圣杰",
		UserId:   "308753",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() - 3600*24,
			Issuer:    "dataMange",
		},
	}

	// 正常token
	genToken["normal"] = MyClaims{
		UserName: "朱圣杰",
		UserId:   "308753",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600*24*365,
			Issuer:    "dataMange",
		},
	}

	// 临期token
	genToken["nearingOverdue"] = MyClaims{
		UserName: "朱圣杰",
		UserId:   "308753",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() - 60*30,
			Issuer:    "dataMange",
		},
	}

	// 没有UserName的token
	genToken["noUserName"] = MyClaims{
		UserId: "308753",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600*24*365,
			Issuer:    "dataMange",
		},
	}

	// 没有UserId的token
	genToken["noUserId"] = MyClaims{
		UserName: "朱圣杰",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600*24*365,
			Issuer:    "dataMange",
		},
	}

	// useName和userId不在数据库
	genToken["noRecord"] = MyClaims{
		UserName: "周曼",
		UserId:   "307714",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600*24*365,
			Issuer:    "dataMange",
		},
	}

	// UserName和UserId不匹配的token
	genToken["noMatch"] = MyClaims{
		UserName: "赖书年",
		UserId:   "308753",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600*24*365,
			Issuer:    "dataMange",
		},
	}

	// 伪造的token
	genToken["fakeToken"] = MyClaims{
		UserName: "朱圣杰",
		UserId:   "308753",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600*24*365,
			Issuer:    "dataMange",
		},
	}
}
