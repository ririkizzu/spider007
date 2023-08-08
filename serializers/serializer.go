package serializers

type ListPageRequest struct {
	Offset   int `json:"offset" form:"offset"`
	PageSize int `json:"page_size" form:"page_size"`
}
