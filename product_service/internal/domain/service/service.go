package service

import (
	"product_service/internal/domain/entity"
	"product_service/internal/domain/interfaces"
	"product_service/internal/domain/model"
)

type ProductService struct {
	repo interfaces.ProductRepository
}

func NewProductService(repo interfaces.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(page, pageSize int64) ([]model.Product, error) {
	productEntity, err := s.repo.GetProducts(page, pageSize)

	if err != nil {
		return nil, err
	}

	var products []model.Product
	for _, product := range productEntity {
		products = append(products, s.EntityToModel(product))
	}

	return products, nil
}

func (s *ProductService) GetProduct(id int64) (model.Product, error) {
	productEntity, err := s.repo.GetProduct(id)
	if err != nil {
		return model.Product{}, err
	}

	return s.EntityToModel(productEntity), nil
}

func (s *ProductService) EntityToModel(entity entity.ProductEntity) model.Product {
	return model.Product{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Price:       entity.Price,
		Category:    entity.Category,
		Brand:       entity.Brand,
		URL:         entity.URL,
	}
}
