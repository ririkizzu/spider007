package plugins

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ririkizzu/spider007/pkg"
	"io/ioutil"
	"net/http"
	"strings"
)

type RegisterScan func(phoneNum int64) (registered bool, err error)

func InitRegisterScanMap() map[string]RegisterScan {
	var registerScanMap map[string]RegisterScan
	registerScanMap = make(map[string]RegisterScan)
	registerScanMap["微博"] = ScanWeibo
	registerScanMap["自如"] = ScanZiRoom
	registerScanMap["百度"] = ScanBaidu
	return registerScanMap
}

func ScanWeibo(phoneNum int64) (registered bool, err error) {
	registered = false
	url := fmt.Sprintf("https://weibo.com/signup/v5/formcheck?type=mobilesea&zone=0086&value=%d", phoneNum)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	request.Header.Add("Referer", "https://weibo.com/signup/signup.php")
	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var bodyMap map[string]interface{}
	_ = json.Unmarshal(body, &bodyMap)
	code := bodyMap["code"]
	if code == "600001" {
		registered = true
	}
	return registered, err
}

func ScanZiRoom(phoneNum int64) (registered bool, err error) {
	registered = false
	//先得到DES的key和iv
	desUrl := "https://passport.ziroom.com/account/des/index.html"
	desRequest, _ := http.NewRequest("GET", desUrl, nil)
	desRequest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:102.0) Gecko/20100101 Firefox/102.0")
	desRequest.Header.Add("Cookie", "PHPSESSID=l20br6erqh7e8etqbhfkqlitm3")
	desRequest.Header.Add("Connection", "close")
	desResponse, err := http.DefaultClient.Do(desRequest)
	defer desResponse.Body.Close()
	desBody, _ := ioutil.ReadAll(desResponse.Body)
	var desBodyMap map[string]map[string]interface{}
	_ = json.Unmarshal(desBody, &desBodyMap)
	secretKey := desBodyMap["data"]["secret_key"]
	secretIv := desBodyMap["data"]["secret_iv"]
	//fmt.Println(secretKey, secretIv)

	//DES对data进行加密
	url := "https://passport.ziroom.com/account/login/login.html"
	data := fmt.Sprintf("{\"phone\":\"%d\",\"password\":\"111111\",\"seven\":0}", phoneNum)
	//fmt.Println(data)
	dataByte, _ := pkg.DESEn([]byte(data), []byte(fmt.Sprintf("%s", secretKey)), []byte(fmt.Sprintf("%s", secretIv)))
	data = hex.EncodeToString(dataByte)
	//fmt.Println(data)
	payload := strings.NewReader(fmt.Sprintf("data=%s&", data))
	request, _ := http.NewRequest("POST", url, payload)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:102.0) Gecko/20100101 Firefox/102.0")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Cookie", "PHPSESSID=l20br6erqh7e8etqbhfkqlitm3")
	request.Header.Add("Connection", "close")
	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var bodyMap map[string]interface{}
	_ = json.Unmarshal(body, &bodyMap)
	errorMessage := bodyMap["error_message"]
	//fmt.Println(errorMessage)
	if errorMessage == "密码错误" {
		registered = true
	}
	return registered, err
}

func ScanBaidu(phoneNum int64) (registered bool, err error) {
	registered = false
	url := fmt.Sprintf("https://passport.baidu.com/v2/?regphonecheck&phone=%d", phoneNum)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:102.0) Gecko/20100101 Firefox/102.0")
	request.Header.Add("Connection", "close")
	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	body_ := strings.Replace(strings.Replace(string(body), "(", "", -1), ")", "", -1)
	//fmt.Println(body_)
	var bodyMap map[string]interface{}
	_ = json.Unmarshal([]byte(body_), &bodyMap)
	errmsg := bodyMap["errmsg"]
	if errmsg == "手机号已被注册，请直接登录或更换手机号注册" {
		registered = true
	}
	return registered, err
}

// ScanTemplate 扫描函数模版
func ScanTemplate(phoneNum int64) (registered bool, err error) {
	registered = false
	return registered, err
}
