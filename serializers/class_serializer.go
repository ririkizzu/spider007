package serializers

type AddClassRequest struct {
	Name string `json:"name"`
}

type GetClassRequest struct {
	ListPageRequest
}

type GetClassResponse struct {
	List  []*ClassItem `json:"list"`
	Total int64        `json:"total"`
}

type ClassItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
