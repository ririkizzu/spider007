package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ririkizzu/spider007/routers/api/v1"
	"net/http"
)

func InitRouter(env string) *gin.Engine {
	gin.SetMode(env)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//分类
	class := r.Group("/api/class")
	{
		class.POST("/add", v1.AddClass)
		class.GET("/list", v1.GetClass)
	}

	//平台
	platform := r.Group("/api/platform")
	{
		platform.POST("/add", v1.AddPlatform)
		platform.GET("/list", v1.GetPlatform)
	}

	//手机号注册扫描
	register := r.Group("/api/register")
	{
		register.POST("/scan", v1.RegisterScan)
	}

	//用户
	user := r.Group("/api/user")
	{
		user.GET("/check-user", v1.CheckUserByOpenid)
		user.GET("/info", v1.GetUserByOpenid)
		user.GET("/add", v1.AddUser)
		user.GET("/signed", v1.UserSigned)
		user.POST("/update-credit", v1.UpdateUserCreditByOpenid)
		user.GET("/rewarded", v1.Rewarded)
	}

	//日志
	scanLog := r.Group("/api/log")
	{
		scanLog.GET("/list", v1.GetLog)
	}

	return r
}
