package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ririkizzu/spider007/models"
	"github.com/ririkizzu/spider007/pkg/e"
	"github.com/ririkizzu/spider007/pkg/plugins"
	"github.com/ririkizzu/spider007/serializers"
	"log"
	"net/http"
	"sync"
	"time"
)

func RegisterScan(c *gin.Context) {
	code := e.SUCCESS

	requestParam := &serializers.RegisterScanRequest{}

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

	count, _ := models.GetRegisterCountByPhoneNum(requestParam.PhoneNum)

	var platformInfo []*serializers.PlatformItem
	platformInfo = []*serializers.PlatformItem{}
	data := &serializers.RegisterItem{
		PhoneNum: requestParam.PhoneNum,
		Queries:  1,
	}

	scanMap := plugins.InitRegisterScanMap()
	timeNow := time.Now().UnixMilli()

	//如果被扫描的手机号码已经入库
	if count > 0 {
		//请求次数先加+1
		err = models.AddRegisterQueries(requestParam.PhoneNum)
		registerModel, err := models.GetRegisterByPhoneNum(requestParam.PhoneNum)
		if err != nil {
			code = e.MYSQL_ERR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": false,
			})
			return
		}

		registerUpdate := registerModel.RegisterUpdate
		//如果数据大于等于1天未更新，则重新扫描
		if (timeNow-registerUpdate)/86400000 >= 1 {
			//协程提高扫描速率
			//response
			responseChannel := make(chan *serializers.PlatformItem, 10)
			go func() {
				for registerInfo := range responseChannel {
					platformInfo = append(platformInfo, registerInfo)
				}
			}()

			wg := &sync.WaitGroup{}
			//控制并发数为10
			limiter := make(chan bool, 5)
			//request
			for name, scanFunc := range scanMap {
				wg.Add(1)
				limiter <- true
				go func(name string, scanFunc plugins.RegisterScan, wg *sync.WaitGroup) {
					defer wg.Done()
					registered, _ := scanFunc(requestParam.PhoneNum)
					if registered {
						platform, _ := models.GetPlatformByName(name)
						className, _ := models.GetClassNameById(platform.ClassId)
						var tag []string
						_ = json.Unmarshal([]byte(platform.Tag), &tag)
						registerInfo := &serializers.PlatformItem{
							Id:        platform.Id,
							ClassName: className,
							Name:      platform.Name,
							Icon:      platform.Icon,
							Developer: platform.Developer,
							Desc:      platform.Desc,
							Link:      platform.Link,
							Tag:       tag,
						}
						responseChannel <- registerInfo
					}
					<-limiter
				}(name, scanFunc, wg)
			}
			wg.Wait()
			data.RegisterInfo = platformInfo
			data.Queries = registerModel.Queries
			data.Count = len(platformInfo)
			data.RegisterUpdate = timeNow
			platformInfoMarshal, err := json.Marshal(data.RegisterInfo)
			updateMap := make(map[string]interface{}, 0)
			updateMap["register_info"] = string(platformInfoMarshal)
			updateMap["register_update"] = timeNow
			err = models.UpdateRegisterByPhoneNum(data.PhoneNum, updateMap)
			if err != nil {
				code = e.MYSQL_ERR
				c.JSON(http.StatusOK, gin.H{
					"code": code,
					"msg":  e.GetMsg(code),
					"data": false,
				})
				return
			}
			//若未超过5天，则直接返回数据库中的结果
		} else {
			err = json.Unmarshal([]byte(registerModel.RegisterInfo), &platformInfo)
			data.RegisterInfo = platformInfo
			data.PhoneNum = registerModel.PhoneNum
			data.Queries = registerModel.Queries
			data.Count = len(platformInfo)
			data.RegisterUpdate = registerModel.RegisterUpdate
		}
		//如果被扫描的手机号码未入库，是第一次被扫描，则重新扫描入库
	} else {
		//协程提高扫描速率
		//response
		responseChannel := make(chan *serializers.PlatformItem, 10)
		go func() {
			for registerInfo := range responseChannel {
				platformInfo = append(platformInfo, registerInfo)
			}
		}()

		wg := &sync.WaitGroup{}
		//控制并发数为10
		limiter := make(chan bool, 5)
		//request
		for name, scanFunc := range scanMap {
			wg.Add(1)
			limiter <- true
			go func(name string, scanFunc plugins.RegisterScan, wg *sync.WaitGroup) {
				defer wg.Done()
				registered, _ := scanFunc(requestParam.PhoneNum)
				if registered {
					platform, _ := models.GetPlatformByName(name)
					className, _ := models.GetClassNameById(platform.ClassId)
					var tag []string
					_ = json.Unmarshal([]byte(platform.Tag), &tag)
					registerInfo := &serializers.PlatformItem{
						Id:        platform.Id,
						ClassName: className,
						Name:      platform.Name,
						Icon:      platform.Icon,
						Developer: platform.Developer,
						Desc:      platform.Desc,
						Link:      platform.Link,
						Tag:       tag,
					}
					responseChannel <- registerInfo
				}
				<-limiter
			}(name, scanFunc, wg)
		}
		wg.Wait()
		data.RegisterInfo = platformInfo
		data.Count = len(platformInfo)
		data.RegisterUpdate = timeNow

		platformInfoMarshal, err := json.Marshal(data.RegisterInfo)
		register := &models.TRegister{
			PhoneNum:       data.PhoneNum,
			RegisterInfo:   string(platformInfoMarshal),
			Queries:        data.Queries,
			RegisterUpdate: timeNow,
		}
		//入库
		err = models.AddRegister(register)
		if err != nil {
			code = e.MYSQL_ERR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": false,
			})
			return
		}
	}

	//写入扫描日志
	platformInfoMarshal, _ := json.Marshal(data.RegisterInfo)
	scanLog := &models.TLog{
		PhoneNum:     data.PhoneNum,
		Openid:       c.GetHeader("X-WX-OPENID"),
		RegisterInfo: string(platformInfoMarshal),
	}
	err = models.AddLog(scanLog)
	if err != nil {
		log.Fatalf("write log err: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
	return
}
