/*
 * @Author: 朱圣杰
 * @Date: 2022-09-03 15:16:32
 * @LastEditors: 朱圣杰
 * @LastEditTime: 2022-09-03 17:25:22
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

func setup() {
	fmt.Println("start")
}

func tearDown() {
	fmt.Println("over")
}
