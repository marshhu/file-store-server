package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/marshhu/file-store-server/dbos"
	"github.com/marshhu/file-store-server/handler/resp"
	"github.com/marshhu/file-store-server/util"
	"log"
	"net/http"
)

type AuthRequest struct{
	UserName    string `form:"username"`
	Password string `form:"password"`
}

func AuthHandler(c *gin.Context){
	auth := AuthRequest{}
	err := c.ShouldBind(&auth)
	if err!= nil{
		c.JSON(http.StatusBadRequest, resp.Response{
			Code: resp.INVALID_PARAMS,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}
	result,err :=dbos.CheckUser(auth.UserName, auth.Password)
	if err != nil{
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, resp.Response{
			Code: resp.ERROR,
			Msg:  "服务器发生异常",
			Data: nil,
		})
		return
	}
	if !result{
		c.JSON(http.StatusOK, resp.Response{
			Code: resp.ERROR,
			Msg:  "账号或密码错误",
			Data: nil,
		})
		return
	}

	token, err := util.GenerateToken(auth.UserName, auth.Password)
	if err != nil{
		c.JSON(http.StatusInternalServerError, resp.Response{
			Code: resp.ERROR,
			Msg:  "生成Token失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, resp.Response{
		Code: resp.SUCCESS,
		Msg:  "OK",
		Data: map[string]string{"token":token},
	})
}

func GetUserByTokenHandler(c *gin.Context){
	token := c.Query("token")
	if len(token) <= 0{
		c.JSON(http.StatusBadRequest, resp.Response{
			Code: resp.INVALID_PARAMS,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	tokenInfo,err := util.ParseToken(token)
	if err != nil{
		c.JSON(http.StatusBadRequest, resp.Response{
			Code: resp.INVALID_PARAMS,
			Msg:  "token失效",
			Data: nil,
		})
		return
	}
	user,err := dbos.GetUserInfo(tokenInfo.Username)
	if err != nil{
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, resp.Response{
			Code: resp.ERROR,
			Msg:  "获取用户信息发生错误",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, resp.Response{
		Code: resp.SUCCESS,
		Msg:  "OK",
		Data: map[string]string{
			"name": user.UserName,
			"avatar":user.Avatar,
		},
	})

}