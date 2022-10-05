package web

type ActivityRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	User_id     uint64 `json:"user_id" binding:"required"`
	UpdatedBy   string `json:"updated_by"`
	DeletedBy   string `json:"deleted_by"`
}

type ActivityUpdateRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	User_id     uint64 `json:"user_id" binding:"required"`
	UpdatedBy   string `json:"updated_by" binding:"required"`
	DeletedBy   string `json:"deleted_by"`
}
