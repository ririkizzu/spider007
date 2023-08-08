package models

import (
	"encoding/json"
	"github.com/ririkizzu/spider007/serializers"
)

type TPlatform struct {
	ClassId   int
	Name      string
	Icon      string
	Developer string
	Desc      string
	Link      string
	Tag       string
	Model
}

func AddPlatform(platform *TPlatform) (*serializers.PlatformItem, error) {
	err := db.Model(&TPlatform{}).Create(platform).Error
	if err != nil {
		return nil, err
	}

	var tag []string
	_ = json.Unmarshal([]byte(platform.Tag), &tag)

	className, err := GetClassNameById(platform.ClassId)
	data := &serializers.PlatformItem{
		Id:        platform.Id,
		ClassName: className,
		Name:      platform.Name,
		Icon:      platform.Icon,
		Developer: platform.Developer,
		Desc:      platform.Desc,
		Link:      platform.Link,
		Tag:       tag,
	}
	return data, err
}

func GetPlatform(request *serializers.GetPlatformRequest) ([]*TPlatform, int64, error) {
	var platformList []*TPlatform
	var count int64

	dbModel := db.Model(&TPlatform{})

	dbModel.Count(&count)

	if request.PageSize > 0 {
		dbModel.Offset(request.Offset).Limit(request.PageSize)
	}

	err := dbModel.Scan(&platformList).Error

	return platformList, count, err
}

func GetPlatformByName(name string) (*TPlatform, error) {
	var platform *TPlatform
	err := db.Model(&TPlatform{}).Where("name = ?", name).First(&platform).Error
	return platform, err
}
