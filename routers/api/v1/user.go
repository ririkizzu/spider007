package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ririkizzu/spider007/models"
	"github.com/ririkizzu/spider007/pkg"
	"github.com/ririkizzu/spider007/pkg/e"
	"github.com/ririkizzu/spider007/serializers"
	"net/http"
	"time"
)

func CheckUserByOpenid(c *gin.Context) {
	code := e.SUCCESS

	openid := c.GetHeader("X-WX-OPENID")
	count, err := models.GetUserCountByOpenid(openid)
	if err != nil {
		code = e.MYSQL_ERR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": count,
	})
	return
}

func GetUserByOpenid(c *gin.Context) {
	code := e.SUCCESS

	openid := c.GetHeader("X-WX-OPENID")
	userModel, err := models.GetUserByOpenid(openid)
	if err != nil {
		code = e.MYSQL_ERR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "",
		})
		return
	}

	rewardedAt := userModel.RewardedAt
	isSameDay := pkg.IsSameDay(time.Now().UnixMilli(), rewardedAt)

	if !isSameDay {
		updateMap := map[string]interface{}{
			"rewarded_count": 0,
			"limited":        2,
		}
		err = models.UpdateUserRewardedInfoByOpenid(openid, updateMap)
		if err != nil {
			code = e.MYSQL_ERR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": "",
			})
			return
		}
	}

	userModel, err = models.GetUserByOpenid(openid)
	user := &serializers.UserItem{
		Id:            userModel.Id,
		Openid:        userModel.Openid,
		AvatarUrl:     userModel.AvatarUrl,
		NickName:      userModel.NickName,
		Credit:        userModel.Credit,
		SignedAt:      userModel.SignedAt,
		RewardedCount: userModel.RewardedCount,
		Limited:       userModel.Limited,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": user,
	})
	return
}

func AddUser(c *gin.Context) {
	code := e.SUCCESS

	openid := c.GetHeader("X-WX-OPENID")
	user := &models.TUser{
		Openid:    openid,
		AvatarUrl: "https://mmbiz.qpic.cn/mmbiz/icTdbqWNOwNRna42FI242Lcia07jQodd2FJGIYQfG0LAJGFxM4FbnQP6yfMxBgJ0F3YRqJCJ1aPAK2dQagdusBZg/0",
		NickName:  "微信用户",
		Credit:    5,
	}
	addUser, err := models.AddUser(user)
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
		"data": addUser,
	})
	return
}

//TODO 后台判断用户是否签到
func UserSigned(c *gin.Context) {
	code := e.SUCCESS

	openid := c.GetHeader("X-WX-OPENID")
	timestamp := time.Now().UnixMilli()
	err := models.UpdateUserSignedAtByOpenid(openid, timestamp)
	if err != nil {
		code = e.MYSQL_ERR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": false,
		})
		return
	}
	err = models.UpdateUserCreditByOpenid(openid, 2)
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
		"data": true,
	})
	return
}

func UpdateUserCreditByOpenid(c *gin.Context) {
	code := e.SUCCESS

	openid := c.GetHeader("X-WX-OPENID")

	requestParam := &serializers.UpdateUserCreditRequest{}
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

	err = models.UpdateUserCreditByOpenid(openid, requestParam.Credit)
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
		"data": true,
	})
}

func Rewarded(c *gin.Context) {
	code := e.SUCCESS

	openid := c.GetHeader("X-WX-OPENID")

	timestamp := time.Now().UnixMilli()
	err := models.UpdateUserRewardedAtByOpenid(openid, timestamp)
	if err != nil {
		code = e.MYSQL_ERR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": false,
		})
		return
	}

	err = models.UpdateUserRewardedCountByOpenid(openid)
	if err != nil {
		code = e.MYSQL_ERR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": false,
		})
		return
	}

	err = models.UpdateUserCreditByOpenid(openid, 3)
	if err != nil {
		code = e.MYSQL_ERR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": false,
		})
		return
	}

	userModel, err := models.GetUserByOpenid(openid)
	if err != nil {
		code = e.MYSQL_ERR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "",
		})
		return
	}

	rewardedCount := userModel.RewardedCount
	if rewardedCount == 10 {
		updateMap := map[string]interface{}{
			"limited": 1,
		}
		err = models.UpdateUserRewardedInfoByOpenid(openid, updateMap)
		if err != nil {
			code = e.MYSQL_ERR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": "",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": true,
	})
	return
}
