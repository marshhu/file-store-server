package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func MD5(data string) string {
	_md5 := md5.New()
	_md5.Write([]byte(data))
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func FileMD5(file *os.File) string {
	_md5 := md5.New()
	io.Copy(_md5, file)
	return hex.EncodeToString(_md5.Sum(nil))
}