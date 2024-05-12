package service

import (
	"errors"
	"product_service/internal/domain/entity"
	"product_service/internal/domain/model"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type mockProductRepository struct {
	products []entity.ProductEntity
	err      error
}

func (m *mockProductRepository) GetProducts(page, pageSize int64) ([]entity.ProductEntity, error) {
	return m.products, m.err
}

func (m *mockProductRepository) GetProduct(id int64) (entity.ProductEntity, error) {
	for _, product := range m.products {
		if product.ID == id {
			return product, nil
		}
	}
	return entity.ProductEntity{}, errors.New("product not found")
}

func TestProductService_GetProducts(t *testing.T) {
	Convey("Given a ProductService", t, func() {
		// Prepare mock repository
		mockRepo := &mockProductRepository{
			products: []entity.ProductEntity{
				{ID: 1, Title: "Product 1", Description: "Description 1", Price: 100.0, Category: "Category 1", Brand: "Brand 1", URL: "http://example.com/product1"},
				{ID: 2, Title: "Product 2", Description: "Description 2", Price: 200.0, Category: "Category 2", Brand: "Brand 2", URL: "http://example.com/product2"},
			},
			err: nil,
		}

		service := NewProductService(mockRepo)

		Convey("When calling GetProducts", func() {
			products, err := service.GetProducts(1, 10)

			Convey("Then the returned products should match the expected ones", func() {
				expectedProducts := []model.Product{
					{ID: 1, Title: "Product 1", Description: "Description 1", Price: 100.0, Category: "Category 1", Brand: "Brand 1", URL: "http://example.com/product1"},
					{ID: 2, Title: "Product 2", Description: "Description 2", Price: 200.0, Category: "Category 2", Brand: "Brand 2", URL: "http://example.com/product2"},
				}
				So(len(products), ShouldEqual, len(expectedProducts))
				for i, p := range products {
					So(p, ShouldResemble, expectedProducts[i])
				}
			})

			Convey("Then the error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestProductService_GetProduct(t *testing.T) {
	Convey("Given a ProductService", t, func() {
		// Prepare mock repository
		mockRepo := &mockProductRepository{
			products: []entity.ProductEntity{
				{ID: 1, Title: "Product 1", Description: "Description 1", Price: 100.0, Category: "Category 1", Brand: "Brand 1", URL: "http://example.com/product1"},
				{ID: 2, Title: "Product 2", Description: "Description 2", Price: 200.0, Category: "Category 2", Brand: "Brand 2", URL: "http://example.com/product2"},
			},
			err: nil,
		}

		service := NewProductService(mockRepo)

		Convey("When calling GetProduct", func() {
			product, err := service.GetProduct(1)

			Convey("Then the returned product should match the expected one", func() {
				expectedProduct := model.Product{
					ID:          1,
					Title:       "Product 1",
					Description: "Description 1",
					Price:       100.0,
					Category:    "Category 1",
					Brand:       "Brand 1",
					URL:         "http://example.com/product1",
				}
				So(product, ShouldResemble, expectedProduct)
			})

			Convey("Then the error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
