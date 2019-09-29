package dbos

import (
	"testing"
)

func clearTables() {
	dbConn.Exec("truncate file_info")
	dbConn.Exec("truncate user")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
}

var fileSha1 string = "erfegfb45t5y6u76ihgmu7"

func TestFileInfoWorkFlow(t *testing.T) {
	t.Run("Add", testAddFileInfo)
	t.Run("Exist", testIsExistFileInfo)
	t.Run("Get", testGetFileInfo)
	t.Run("List", testGetFileList)
}
func testAddFileInfo(t *testing.T) {
	err := AddFileInfo(fileSha1, "avatar.jpg",23400,"tmp/avatar.jpg")
	if err != nil {
		t.Errorf("Error of AddFileInfo:%v", err)
	}
}

func testIsExistFileInfo(t *testing.T) {
   isExist := IsExistFileInfo(fileSha1)
   if !isExist{
   	 t.Failed()
   }
}

func testGetFileInfo(t *testing.T) {
   fileInfo,err := GetFileInfo(fileSha1)
   if err != nil{
   	 t.Errorf("test GetFileInfo falied:%v",err)
   }
   if fileInfo == nil{
   	 t.Failed()
   }
}

func testGetFileList(t *testing.T) {
	fileList, err := GetFileList()
	if err != nil {
		t.Errorf("test GetFileList falied:%v", err)
	}
	if len(fileList) == 0 {
		t.Failed()
	}
}