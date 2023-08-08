package models

import (
	"gorm.io/gorm"
)

type TUser struct {
	Openid        string
	AvatarUrl     string
	NickName      string
	Credit        int
	SignedAt      int64
	RewardedCount int
	RewardedAt    int64
	Limited       int8
	Model
}

func GetUserCountByOpenid(openid string) (int64, error) {
	var count int64
	err := db.Model(&TUser{}).Where("openid = ?", openid).Count(&count).Error
	return count, err
}

func GetUserByOpenid(openid string) (*TUser, error) {
	var userModel *TUser
	err := db.Model(&TUser{}).Where("openid = ?", openid).First(&userModel).Error
	return userModel, err
}

func AddUser(user *TUser) (bool, error) {
	err := db.Model(&TUser{}).Create(user).Error
	if err != nil {
		return false, err
	}
	return true, err
}

func UpdateUserCreditByOpenid(openid string, value int) error {
	err := db.Model(&TUser{}).Where("openid = ?", openid).Update("credit", gorm.Expr("credit + ?", value)).Error
	return err
}

func UpdateUserSignedAtByOpenid(openid string, timestamp int64) error {
	err := db.Model(&TUser{}).Where("openid = ?", openid).Update("signed_at", timestamp).Error
	return err
}

func UpdateUserRewardedAtByOpenid(openid string, timestamp int64) error {
	err := db.Model(&TUser{}).Where("openid = ?", openid).Update("rewarded_at", timestamp).Error
	return err
}

func UpdateUserRewardedCountByOpenid(openid string) error {
	err := db.Model(&TUser{}).Where("openid = ?", openid).Update("rewarded_count", gorm.Expr("rewarded_count + ?", 1)).Error
	return err
}

func UpdateUserRewardedInfoByOpenid(openid string, updateMap map[string]interface{}) error {
	err := db.Model(&TUser{}).Where("openid = ?", openid).Updates(updateMap).Error
	return err
}
