package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ririkizzu/spider007/models"
	"github.com/ririkizzu/spider007/pkg/e"
	"github.com/ririkizzu/spider007/serializers"
	"net/http"
)

func AddPlatform(c *gin.Context) {
	code := e.SUCCESS

	requestParam := &serializers.AddPlatformRequest{}
	err := c.ShouldBind(requestParam)
	if err != nil {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": false,
		})
		return
	}

	tagString, _ := json.Marshal(requestParam.Tag)

	platform := &models.TPlatform{
		ClassId:   requestParam.ClassId,
		Name:      requestParam.Name,
		Icon:      requestParam.Icon,
		Developer: requestParam.Developer,
		Desc:      requestParam.Desc,
		Link:      requestParam.Link,
		Tag:       string(tagString),
	}
	data, err := models.AddPlatform(platform)
	if err != nil {
		code = e.MYSQL_ERR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
	return
}

func GetPlatform(c *gin.Context) {
	code := e.SUCCESS

	requestParam := &serializers.GetPlatformRequest{}
	err := c.ShouldBind(requestParam)
	if err != nil {
		code = e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "",
		})
		return
	}

	platformList, count, err := models.GetPlatform(requestParam)

	data := &serializers.GetPlatformResponse{
		List: []*serializers.PlatformItem{},
	}

	var className string

	for _, platform := range platformList {
		className, err = models.GetClassNameById(platform.ClassId)
		if err != nil {
			code = e.MYSQL_ERR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": "",
			})
			return
		}

		var tag []string
		_ = json.Unmarshal([]byte(platform.Tag), &tag)

		item := &serializers.PlatformItem{
			Id:        platform.Id,
			ClassName: className,
			Name:      platform.Name,
			Icon:      platform.Icon,
			Developer: platform.Developer,
			Desc:      platform.Desc,
			Link:      platform.Link,
			Tag:       tag,
		}
		data.List = append(data.List, item)
	}

	data.Total = count

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
	return
}
