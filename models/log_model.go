package models

import "github.com/ririkizzu/spider007/serializers"

type TLog struct {
	PhoneNum     int64
	Openid       string
	RegisterInfo string
	Model
}

func AddLog(log *TLog) error {
	err := db.Model(&TLog{}).Create(log).Error
	return err
}

func GetLog(request *serializers.GetLogRequest) ([]*TLog, error) {
	var logList []*TLog

	dbModel := db.Model(&TLog{})

	if request.PhoneNum != 0 {
		dbModel.Where("phone_num = ?", request.PhoneNum)
	}

	if request.Openid != "" {
		dbModel.Where("openid = ?", request.Openid)
	}

	if request.PageSize > 0 {
		dbModel.Offset(request.Offset).Limit(request.PageSize)
	}

	err := dbModel.Scan(&logList).Error
	return logList, err
}
