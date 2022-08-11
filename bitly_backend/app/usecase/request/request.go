package request

type (
	CreateUrl struct {
		DestinationUrl string `json:"destination_url" binding:"required"`
		Source string `json:"source,omitempty"`
		Medium string `json:"medium,omitempty"`
		Campaign string `json:"campaign,omitempty"`
		UserID float64 `json:"user_id"`
	}

	CreateUrlWithoutUser struct {
		DestinationUrl string `json:"destination_url" binding:"required"`
	}

	UpdateUrl struct {
		DestinationUrl string `json:"destination_url" binding:"required"`
		ModifyUrl string `json:"modify_url" binding:"required"`
		Title string `json:"title"`
		Hastag string `json:"hastag"`
		Source string `json:"source,omitempty"`
		Medium string `json:"medium,omitempty"`
		Campaign string `json:"campaign,omitempty"`
		ID int `json:"id" binding:"required"`
	}

	GetListProductByUserID struct {
		UserID int `json:"user_id" binding:"required"`
	}

	PatchGetClick struct {
		ID int `json:"id" binding:"required"`
	}

	Register struct {
		Username string `json:"username" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	Login struct {
		Find string `json:"find" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)