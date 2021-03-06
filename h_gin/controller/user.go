package controller

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// @Summary 测试专用
// @tags test
// @Produce  json
// @Param Token header string true "Token"
// @Param test query string true "Test"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /test/info [get]
func GetTest(ctx *gin.Context) {
	fmt.Println("test", ctx.Query("test"))
	fmt.Println("token", ctx.Request.Header.Get("Token"))

	urlPath := ctx.Request.URL.Path
	fmt.Printf("urlPath %v \n", urlPath)

	ctx.JSON(200, gin.H{
		"message": "test",
	})
}

// @Summary 测试专用
// @tags test
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param test formData string true "Test"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /test/add [post]
func PostTest(ctx *gin.Context) {
	fmt.Println(ctx.PostForm("test"))
	fmt.Println("token", ctx.Request.Header.Get("Token"))

	urlPath := ctx.Request.URL.Path
	fmt.Printf("urlPath %v \n", urlPath)

	ctx.JSON(200, gin.H{
		"message": "test",
	})
}

// @Summary 获取用户信息
// @tags user
// @Produce  json
// @Param Token header string true "Token"
// @Param user_id query int true "userId"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /user [get]
func GetUser(ctx *gin.Context) {
	fmt.Println(ctx.Query("user_id"))

	ctx.JSON(200, gin.H{
		"message": "getUser",
	})
}

type User struct {
	NickName string `form:"nick_name" valid:"required~昵称不能为空"`
	Email    string `form:"email" valid:"required~昵称不能为空,email~邮件账号不合法"`
	Phone    string `form:"phone" valid:"required~昵称不能为空,numeric~号码应为数字"`
}

// @Summary 新增用户
// @tags user
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param nick_name formData string true "NickName"
// @Param email formData string true "Email"
// @Param phone formData int false "Phone"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user [post]
func PostUser(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBind(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	result, err := govalidator.ValidateStruct(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)

	fmt.Println("user_name", user.NickName)
	fmt.Println("email", user.Email)
	fmt.Println("phone", user.Phone)

	ctx.JSON(200, gin.H{
		"message": "test",
	})
}

// @Summary 修改用户
// @tags user
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param user_id formData int true "UserId"
// @Param nick_name formData string true "NickName"
// @Param email formData string true "Email"
// @Param phone formData int true "Phone"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user [put]
func PutUser(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println("put____", string(body))
	fmt.Println("put____", string(body))
}

// @Summary 删除用户
// @tags user
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param user_id formData int true "UserId"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user [delete]
func DeleteUser(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println("delete____", string(body))
}
