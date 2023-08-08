package models

import "gorm.io/gorm"

type TRegister struct {
	PhoneNum       int64
	RegisterInfo   string
	Queries        int
	RegisterUpdate int64
	Model
}

func AddRegister(register *TRegister) error {
	err := db.Model(&TRegister{}).Create(register).Error
	return err
}

func GetRegisterCountByPhoneNum(phoneNum int64) (int64, error) {
	var count int64
	err := db.Model(&TRegister{}).Where("phone_num = ?", phoneNum).Count(&count).Error
	return count, err
}

func GetRegisterByPhoneNum(phoneNum int64) (*TRegister, error) {
	var registerModel *TRegister
	err := db.Model(&TRegister{}).Where("phone_num = ?", phoneNum).First(&registerModel).Error
	return registerModel, err
}

func AddRegisterQueries(phoneNum int64) error {
	err := db.Model(&TRegister{}).Where("phone_num = ?", phoneNum).Update("queries", gorm.Expr("queries + ?", 1)).Error
	return err
}

func UpdateRegisterByPhoneNum(phoneNum int64, updateMap map[string]interface{}) error {
	err := db.Model(&TRegister{}).Where("phone_num = ?", phoneNum).Updates(updateMap).Error
	return err
}
