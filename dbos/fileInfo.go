package dbos

import (
	"errors"
	"time"
)

type FileInfo struct {
	ID         int64      `json:"id"`
	FileSha1    string    `json:"file_sha1"`
	FileName    string    `json:"file_name"`
	FileSize    int64     `json:"file_size"`
	FileAddress string    `json:"file_address"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	Status      int       `json:"status"`
}

func AddFileInfo(fileSha1 string,fileName string,fileSize int64,fileAddress string) error{
    if dbConn != nil{
		stmtIns, err := dbConn.Prepare("Insert into file_info(file_sha1,file_name,file_size,file_address,create_at,update_at,status) values(?,?,?,?,?,?,?);")
		if err != nil {
			return err
		}
		_,err = stmtIns.Exec(fileSha1, fileName,fileSize,fileAddress,time.Now(),time.Now(),1)
		if err != nil{
			return  err
		}
		defer stmtIns.Close()
		return nil
	}
	return  errors.New("数据库连接失败")
}