package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/file-store-server/handler/resp"
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
	// single file
	f, _ := c.FormFile("file")
	if f == nil{
		c.JSON(http.StatusBadRequest, resp.Response{
			Code: resp.ERROR,
			Msg:  "请选择上传文件",
			Data: nil,
		})
		return
	}
	log.Println(f.Filename)

	// Upload the file to specific dst.
    dst := "tmp/"+f.Filename
	err := c.SaveUploadedFile(f, dst)
	if err != nil{
		c.JSON(http.StatusInternalServerError, resp.Response{
			Code: resp.ERROR,
			Msg:  "上传文件失败",
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, resp.Response{
		Code: resp.SUCCESS,
		Msg:  "上传文件成功",
		Data: dst,
	})
}

func UploadMultiHandler(c *gin.Context){
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, f := range files {
		log.Println(f.Filename)

		// Upload the file to specific dst.
		dst := "tmp/"+f.Filename
		c.SaveUploadedFile(f, dst)
	}

	c.JSON(http.StatusOK, resp.Response{
		Code: resp.SUCCESS,
		Msg:  "上传文件成功",
		Data: nil,
	})
}
