package serializers

type AddPlatformRequest struct {
	ClassId   int      `json:"class_id"`
	Name      string   `json:"name"`
	Icon      string   `json:"icon"`
	Developer string   `json:"developer"`
	Desc      string   `json:"desc"`
	Link      string   `json:"link"`
	Tag       []string `json:"tag"`
}

type GetPlatformRequest struct {
	ListPageRequest
}

type GetPlatformResponse struct {
	List  []*PlatformItem `json:"list"`
	Total int64           `json:"total"`
}

type PlatformItem struct {
	Id        int      `json:"id"`
	ClassName string   `json:"class_name"`
	Name      string   `json:"name"`
	Icon      string   `json:"icon"`
	Developer string   `json:"developer"`
	Desc      string   `json:"desc"`
	Link      string   `json:"link"`
	Tag       []string `json:"tag"`
}
