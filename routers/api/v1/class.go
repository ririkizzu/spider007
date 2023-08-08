package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ririkizzu/spider007/models"
	"github.com/ririkizzu/spider007/pkg/e"
	"github.com/ririkizzu/spider007/serializers"
	"net/http"
)

func AddClass(c *gin.Context) {
	code := e.SUCCESS

	requestParam := &serializers.AddClassRequest{}
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

	class := &models.TClass{
		Name: requestParam.Name,
	}
	data, err := models.AddClass(class)
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

func GetClass(c *gin.Context) {
	code := e.SUCCESS

	requestParam := &serializers.GetClassRequest{}
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

	classModelArr, total, err := models.GetClassList(requestParam)
	if err != nil {
		code = e.MYSQL_ERR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "",
		})
		return
	}

	data := &serializers.GetClassResponse{
		List: []*serializers.ClassItem{},
	}

	for _, class := range classModelArr {
		item := &serializers.ClassItem{
			Id:   class.Id,
			Name: class.Name,
		}
		data.List = append(data.List, item)
	}

	data.Total = total

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
	return
}
