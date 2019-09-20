package dbos

import "testing"

func clearTables() {
	dbConn.Exec("truncate file_info")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
}

func TestAddFileInfo(t *testing.T) {
	err := AddFileInfo("erfegfb45t5y6u76ihgmu7", "avatar.jpg",23400,"/tmp/avatar.jpg")
	if err != nil {
		t.Errorf("Error of AddFileInfo:%v", err)
	}
}