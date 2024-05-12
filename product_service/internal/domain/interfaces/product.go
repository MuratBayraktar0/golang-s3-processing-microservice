package interfaces

import (
	"product_service/internal/domain/entity"
)

// We will be able to change the database without being dependent on MongoDB and without making any code changes at the domain, or core layer, thanks to this interface.
type ProductRepository interface {
	GetProducts(page, pageSize int64) ([]entity.ProductEntity, error)
	GetProduct(id int64) (entity.ProductEntity, error)
}
