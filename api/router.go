package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.POST("/api/register", UserRegister)             // 注册
	r.POST("/api/namelogin", UserNameLogin)           //账号密码登录
	r.POST("/api/emaillogin", EmailLogin)             //邮箱密码登录
	r.POST("/api/phonenumberlogin", PhoneNumberLogin) //手机号密码登录
	r.Run(":8090")                                    // 跑在 8090 端口上
}
