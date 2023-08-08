package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ririkizzu/spider007/models"
	"github.com/ririkizzu/spider007/routers"
	"github.com/ririkizzu/spider007/setting"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Print("微信云托管服务启动成功")

	setting.InitSetting(gin.DebugMode)
	router := routers.InitRouter(gin.DebugMode)
	models.InitDBConn()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:      router,
		ReadTimeout:  setting.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout: setting.ServerSetting.WriteTimeout * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
