package ResponseDTOs

type GetAllApiTokensResponseDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Result []struct {
		Description    string `json:"description"`
		ExpireAtInMs   int64  `json:"expireAtInMs"`
		Id             int    `json:"id"`
		Name           string `json:"name"`
		Token          string `json:"token"`
		UpdatedAt      string `json:"updatedAt"`
		UserId         int    `json:"userId"`
		UserIdentifier string `json:"userIdentifier"`
	} `json:"result"`
}
