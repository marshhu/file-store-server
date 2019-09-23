package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marshhu/file-store-server/conf"
	"github.com/marshhu/file-store-server/dbos"
	"github.com/marshhu/file-store-server/handler/resp"
	"github.com/marshhu/file-store-server/util"
	"log"
	"net/http"
)

func PingHandler(c *gin.Context){
	c.JSON(http.StatusOK, resp.Response{
		Code: resp.SUCCESS,
		Msg:  "ping success",
		Data: nil,
	})
}

func UploadSingleHandler(c *gin.Context){
	maxUploadSize := conf.AppSetting.MaxUploadSize
	c.Request.Body = http.MaxBytesReader(c.Writer,c.Request.Body,maxUploadSize)
    if err := c.Request.ParseMultipartForm(maxUploadSize);err != nil{
		c.JSON(http.StatusBadRequest, resp.Response{
			Code: resp.INVALID_PARAMS,
			Msg:  fmt.Sprintf("文件太大，超过了%dM",maxUploadSize),
			Data: nil,
		})
		return
	}

	// single file
	f, _ := c.FormFile("file")
	if f == nil{
		c.JSON(http.StatusBadRequest, resp.Response{
			Code: resp.INVALID_PARAMS,
			Msg:  "请选择上传文件",
			Data: nil,
		})
		return
	}
	log.Println(f.Filename)

	//// Upload the file to specific dst.
    //dst := "upload/"+f.Filename
	//err := c.SaveUploadedFile(f, dst)
	//if err != nil{
	//	c.JSON(http.StatusInternalServerError, resp.Response{
	//		Code: resp.ERROR,
	//		Msg:  "上传文件失败",
	//		Data: nil,
	//	})
	//	return
	//}

	src, _ := f.Open()
	defer src.Close()
	fileSha1 := util.FileSha1(src)
	fileSize := f.Size
	// 游标重新回到文件头部
	src.Seek(0,0)
	fileName := fileSha1 + util.GetExt(f.Filename) //文件名唯一
	//如果文件已存在 获取文件信息并返回
	if dbos.IsExistFileInfo(fileSha1){
		fileInfo,_ := dbos.GetFileInfo(fileSha1)
		c.JSON(http.StatusOK, resp.Response{
			Code: resp.SUCCESS,
			Msg:  "上传文件成功",
			Data: map[string]string{"file_sha1":fileInfo.FileSha1,"file_name":fileInfo.FileName,"file_address":fileInfo.FileAddress},
		})
		return
	}

	//上传到云OSS
	fileAddress,err := util.PutObjectToOSS(fileName,src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Response{
			Code: resp.ERROR,
			Msg:  "上传到云服务oss失败",
			Data: nil,
		})
	}

   //保存文件信息
	err = dbos.AddFileInfo(fileSha1,fileName,fileSize,fileAddress)
	if err != nil{
		c.JSON(http.StatusInternalServerError, resp.Response{
			Code: resp.ERROR,
			Msg:  "保存文件信息到数据库失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, resp.Response{
		Code: resp.SUCCESS,
		Msg:  "上传文件成功",
		Data: map[string]string{"file_sha1":fileSha1,"file_name":fileName,"file_address":fileAddress},
	})
}

func UploadMultiHandler(c *gin.Context){
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, f := range files {
		log.Println(f.Filename)

		// Upload the file to specific dst.
		dst := "upload/"+f.Filename
		c.SaveUploadedFile(f, dst)
	}

	c.JSON(http.StatusOK, resp.Response{
		Code: resp.SUCCESS,
		Msg:  "上传文件成功",
		Data: nil,
	})
}

func GetFileInfoHandler(c *gin.Context){
	fileSha1 := c.Param("fileSha1")
	if len(fileSha1) <=0{
		c.JSON(http.StatusBadRequest, resp.Response{
			Code: resp.INVALID_PARAMS,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}
	fileInfo,err := dbos.GetFileInfo(fileSha1)
	if err != nil{
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, resp.Response{
			Code: resp.ERROR,
			Msg:  "获取文件信息发生错误",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, resp.Response{
		Code: resp.SUCCESS,
		Msg:  "OK",
		Data: fileInfo,
	})
}
