package serializers

type GetLogRequest struct {
	PhoneNum int64  `json:"phone_num" form:"phone_num"`
	Openid   string `json:"openid" form:"openid"`
	ListPageRequest
}

type GetLogResponse struct {
	List  []*LogItem `json:"list"`
	Total int        `json:"total"`
}

type LogItem struct {
	Id           int             `json:"id"`
	PhoneNum     int64           `json:"phone_num"`
	Openid       string          `json:"openid"`
	RegisterInfo []*PlatformItem `json:"register_info"`
}
