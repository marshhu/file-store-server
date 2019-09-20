package dbos

import "testing"

func clearTables() {
	dbConn.Exec("truncate file_info")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
}
var fileSha1 string = "erfegfb45t5y6u76ihgmu7"

func TestWorkFlow(t *testing.T) {
	t.Run("Add", testAddFileInfo)
	t.Run("Exist", testIsExistFileInfo)
	t.Run("Get", testGetFileInfo)
}
func testAddFileInfo(t *testing.T) {
	err := AddFileInfo(fileSha1, "avatar.jpg",23400,"/tmp/avatar.jpg")
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