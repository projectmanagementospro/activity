package web

type ActivityRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:="description"`
	CreatedBy   string `json:="createdby" `
	UpdateBy    string `json:="updatedby" `
}

type ActivityUpdateRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:="description"`
	CreatedBy   string `json:="createdby" `
	UpdateBy    string `json:="updatedby" `
}
