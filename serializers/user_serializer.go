package serializers

type UpdateUserCreditRequest struct {
	Credit int `json:"credit"`
}

type UserItem struct {
	Id            int    `json:"id"`
	Openid        string `json:"openid"`
	AvatarUrl     string `json:"avatar_url"`
	NickName      string `json:"nick_name"`
	Credit        int    `json:"credit"`
	SignedAt      int64  `json:"signed_at"`
	RewardedCount int    `json:"rewarded_count"`
	Limited       int8   `json:"limited"`
}
