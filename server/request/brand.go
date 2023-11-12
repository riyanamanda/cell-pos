package request

type StoreBrandRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateBrandRequest struct {
	Name string `json:"name" binding:"required"`
}