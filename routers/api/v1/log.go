package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ririkizzu/spider007/models"
	"github.com/ririkizzu/spider007/pkg/e"
	"github.com/ririkizzu/spider007/serializers"
	"net/http"
)

func GetLog(c *gin.Context) {
	code := e.SUCCESS

	requestParam := &serializers.GetLogRequest{}
	err := c.ShouldBind(requestParam)
	if err != nil {
		code := e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "",
		})
		return
	}

	logModelArr, err := models.GetLog(requestParam)
	if err != nil {
		code = e.MYSQL_ERR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "",
		})
		return
	}

	data := &serializers.GetLogResponse{
		List: []*serializers.LogItem{},
	}

	for _, log := range logModelArr {
		var registerInfo []*serializers.PlatformItem
		_ = json.Unmarshal([]byte(log.RegisterInfo), &registerInfo)
		item := &serializers.LogItem{
			Id:           log.Id,
			PhoneNum:     log.PhoneNum,
			Openid:       log.Openid,
			RegisterInfo: registerInfo,
		}
		data.List = append(data.List, item)
	}

	data.Total = len(data.List)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
