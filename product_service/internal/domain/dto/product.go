package dto

import (
	"product_service/internal/domain/model"
)

type ProductsResponse struct {
	Data []model.Product `json:"data"`
	Page int64           `json:"page"`
	Size int64           `json:"size"`
}

type ProductDetailResponse struct {
	Data model.Product `json:"data"`
}
