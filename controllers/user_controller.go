package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"yeb/common"
	"yeb/model"
	"yeb/response"
	"yeb/util"
)

//func Login(c *gin.Context){
//	//获取参数
//	name := c.PostForm("username")
//	password := c.PostForm("password")
//	code := c.PostForm("code")
//
//
//	response.Response(c, http.StatusOK, 200, )
//}

// Register 用户注册
func Register(c *gin.Context){
	DB := common.GetDB()
	requestUser := model.User{}
	//gin提供了一种方式Bind
	err := c.ShouldBind(&requestUser)
	if err != nil {
		return
	}
	//获取参数
	name := requestUser.Name
	pwd := requestUser.Pwd

	if len(pwd) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, pwd)

	//创建用户
	//为用户密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name: name,
		Pwd:  string(hasedPassword),
	}
	DB.Create(&newUser)

	response.Success(c, nil, "用户注册成功")
}