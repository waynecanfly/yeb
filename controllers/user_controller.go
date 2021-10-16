package controllers

import (
	"fmt"
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

func Login(c *gin.Context)  {
	DB := common.GetDB()
	//获取参数
	json := make(map[string]interface{})
	c.BindJSON(&json)
	name := json["username"]
	pwd := json["password"]
	code := json["code"]
	if CaptchaVerify(c, code.(string)) != true {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "验证码验证失败，请重新输入",
		})
		return
	}
	fmt.Println(code)
	//数据验证
	if len(pwd.(string)) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}

	//判断姓名是否存在
	var user model.User
	DB.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "User does not exits",
		})
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(pwd.(string))); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "wrong password",
		})
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("token generate error :%v", err)
		return
	}

	//返回结果
	response.Success(c, gin.H{"token": token, "tokenHead": "Bearer"}, "登录成功")
}