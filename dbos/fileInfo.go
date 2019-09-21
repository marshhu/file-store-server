package dbos

import (
	"database/sql"
	"log"
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

func IsExistFileInfo(fileSha1 string) bool{
	stmtOut, err := dbConn.Prepare("select id from file_info where file_sha1 = ?;")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%s", err)
		return false
	}
    var id int64
	err = stmtOut.QueryRow(fileSha1).Scan(&id)
	if err != nil || err == sql.ErrNoRows{
		return false
	}
    if id > 0 {
		return true
	}
	return false
}

func GetFileInfo(fileSha1 string) (*FileInfo,error){
	stmtOut, err := dbConn.Prepare("select id,file_sha1,file_name,file_size,file_address,create_at,update_at,status from file_info where file_sha1 = ?;")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%s", err)
		return nil,err
	}
	fileInfo := FileInfo{}
	var createTime string
	var updateTime string
	err = stmtOut.QueryRow(fileSha1).Scan(&fileInfo.ID,&fileInfo.FileSha1,&fileInfo.FileName,&fileInfo.FileSize,&fileInfo.FileAddress,&createTime,&updateTime,&fileInfo.Status)
	if err !=nil && err != sql.ErrNoRows{
		return nil, err
	}
	if err == sql.ErrNoRows{
		return nil,nil
	}
    fileInfo.CreateAt,_ = time.Parse("2006-01-02 15:04:05",createTime)
	fileInfo.UpdateAt,_ = time.Parse("2006-01-02 15:04:05",updateTime)
	return &fileInfo,nil
}