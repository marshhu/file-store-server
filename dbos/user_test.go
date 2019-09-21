package dbos

import (
	"github.com/marshhu/file-store-server/util"
	"log"
	"testing"
)

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Check", testCheckUser)
	t.Run("Get", testGetUserInfo)
}
var username ="admin"
var password ="123456"

func testAddUser(t *testing.T) {
	var (
		avatar string = "https://ma-avatar.oss-cn-beijing.aliyuncs.com/avatar1.jpeg"
		phone string = ""
		email string = ""
		profile string = ""
	)
	err := AddUser(username,util.PwdMD5(password),avatar,phone,email,profile)

	if err != nil {
		t.Errorf("Error of AddFileInfo:%v", err)
	}
}

func testCheckUser(t *testing.T) {
	 result,err :=CheckUser(username,util.PwdMD5(password))
	 if err != nil{
		 t.Errorf("Error of CheckUser:%v", err)
	 }
	 if !result{
		t.Failed()
	 }
}

func testGetUserInfo(t *testing.T){
	user,err :=GetUserInfo(username)
	if err != nil{
		t.Errorf("Error of GetUserInfo:%v", err)
	}
	if user == nil{
		t.Failed()
	}
	log.Fatal(user)
}