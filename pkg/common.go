package pkg

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"time"
)

//DES加密函数
func DESEn(origText []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origText = PKCS7Padding(origText, block.BlockSize())
	cryptText := make([]byte, len(origText))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cryptText, origText)
	return cryptText, nil
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

//判断两个UnixMilli时间戳是否是同一天
func IsSameDay(timestamp1 int64, timestamp2 int64) bool {
	// 将时间戳转换为 time.Time 类型
	time1 := time.UnixMilli(timestamp1)
	time2 := time.UnixMilli(timestamp2)
	//转换为中国时区
	chinaTimeZone := time.FixedZone("CST", 8*60*60)
	time1InChina := time1.In(chinaTimeZone)
	time2InChina := time2.In(chinaTimeZone)
	// 比较年、月、日是否相同
	return time1InChina.Year() == time2InChina.Year() && time1InChina.Month() == time2InChina.Month() && time1InChina.Day() == time2InChina.Day()
}

//md5加密
func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	md5Str := hex.EncodeToString(hash[:])
	return md5Str
}
