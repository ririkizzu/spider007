package models

import "github.com/ririkizzu/spider007/serializers"

type TClass struct {
	Name string
	Model
}

func AddClass(class *TClass) (*serializers.ClassItem, error) {
	err := db.Model(&TClass{}).Create(class).Error
	data := &serializers.ClassItem{
		Id:   class.Id,
		Name: class.Name,
	}
	return data, err
}

func GetClassList(request *serializers.GetClassRequest) ([]*TClass, int64, error) {
	var class []*TClass
	var count int64

	dbModel := db.Model(&TClass{})

	dbModel.Count(&count)

	if request.PageSize > 0 {
		dbModel.Offset(request.Offset).Limit(request.PageSize)
	}

	err := dbModel.Scan(&class).Error

	return class, count, err
}

func GetClassNameById(id int) (string, error) {
	var className string
	err := db.Model(&TClass{}).Where("id = ?", id).Select("name").Scan(&className).Error
	return className, err
}
