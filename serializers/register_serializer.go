package serializers

type RegisterScanRequest struct {
	PhoneNum int64 `json:"phone_num"`
}

type RegisterItem struct {
	PhoneNum       int64           `json:"phone_num"`
	RegisterInfo   []*PlatformItem `json:"register_info"`
	Queries        int             `json:"queries"`
	Count          int             `json:"count"`
	RegisterUpdate int64           `json:"register_update"`
}
