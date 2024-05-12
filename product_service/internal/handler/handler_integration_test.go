package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/valyala/fasthttp"

	"product_service/internal/domain/dto"
	"product_service/internal/domain/entity"
	"product_service/internal/domain/model"
	"product_service/internal/domain/service"
	"product_service/internal/handler"
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

func TestProductHandler_GetProducts(t *testing.T) {
	Convey("Given a product handler", t, func() {
		// Prepare mock repository
		mockRepo := &mockProductRepository{
			products: []entity.ProductEntity{
				{ID: 1, Title: "Product 1", Description: "Description 1", Price: 100.0, Category: "Category 1", Brand: "Brand 1", URL: "http://example.com/product1"},
				{ID: 2, Title: "Product 2", Description: "Description 2", Price: 200.0, Category: "Category 2", Brand: "Brand 2", URL: "http://example.com/product2"},
				{ID: 3, Title: "Product 3", Description: "Description 3", Price: 300.0, Category: "Category 3", Brand: "Brand 3", URL: "http://example.com/product3"},
				{ID: 4, Title: "Product 4", Description: "Description 4", Price: 400.0, Category: "Category 4", Brand: "Brand 4", URL: "http://example.com/product4"},
				{ID: 5, Title: "Product 5", Description: "Description 5", Price: 500.0, Category: "Category 5", Brand: "Brand 5", URL: "http://example.com/product5"},
				{ID: 6, Title: "Product 6", Description: "Description 6", Price: 600.0, Category: "Category 6", Brand: "Brand 6", URL: "http://example.com/product6"},
				{ID: 7, Title: "Product 7", Description: "Description 7", Price: 700.0, Category: "Category 7", Brand: "Brand 7", URL: "http://example.com/product7"},
				{ID: 8, Title: "Product 8", Description: "Description 8", Price: 800.0, Category: "Category 8", Brand: "Brand 8", URL: "http://example.com/product8"},
				{ID: 9, Title: "Product 9", Description: "Description 9", Price: 900.0, Category: "Category 9", Brand: "Brand 9", URL: "http://example.com/product9"},
				{ID: 10, Title: "Product 10", Description: "Description 10", Price: 1000.0, Category: "Category 10", Brand: "Brand 10", URL: "http://example.com/product10"},
			},
			err: nil,
		}

		service := service.NewProductService(mockRepo)
		productHandler := handler.NewProductHandler(service)

		Convey("When making a valid request", func() {
			ctx := &fasthttp.RequestCtx{}
			ctx.Request.SetRequestURI("/products?page=1&size=10")

			productHandler.GetProducts(ctx)
			Convey("Then the response should have a status code of 200", func() {
				So(ctx.Response.StatusCode(), ShouldEqual, fasthttp.StatusOK)
			})
			Convey("Then the response should have a content type of 'application/json'", func() {
				So(string(ctx.Response.Header.ContentType()), ShouldEqual, "application/json")
			})

			Convey("Then the response body should be a valid ProductListResponse", func() {
				var response dto.ProductsResponse
				err := json.Unmarshal(ctx.Response.Body(), &response)
				So(err, ShouldBeNil)

				So(response.Data, ShouldHaveLength, 10)
				for i, product := range response.Data {
					expectedProduct := model.Product{
						ID:          int64(i + 1),
						Title:       fmt.Sprintf("Product %d", i+1),
						Price:       float64((i + 1) * 100),
						Category:    fmt.Sprintf("Category %d", i+1),
						Brand:       fmt.Sprintf("Brand %d", i+1),
						URL:         fmt.Sprintf("http://example.com/product%d", i+1),
						Description: fmt.Sprintf("Description %d", i+1),
					}
					So(product, ShouldResemble, expectedProduct)
				}

				So(response.Page, ShouldEqual, 1)
				So(response.Size, ShouldEqual, 10)
			})
		})

		Convey("When making a request with an invalid page parameter", func() {
			ctx := &fasthttp.RequestCtx{}
			ctx.Request.SetRequestURI("/products?page=invalid&size=10")

			productHandler.GetProducts(ctx)

			Convey("Then the response should have a status code of 400", func() {
				So(ctx.Response.StatusCode(), ShouldEqual, fasthttp.StatusBadRequest)
			})

			Convey("Then the response should have a content type of 'application/json'", func() {
				So(string(ctx.Response.Header.ContentType()), ShouldEqual, "application/json")
			})

			Convey("Then the response body should contain an error message", func() {
				var response map[string]string
				err := json.Unmarshal(ctx.Response.Body(), &response)
				So(err, ShouldBeNil)
				So(response["error"], ShouldEqual, "Invalid page")
			})
		})
	})
}

func TestProductHandler_GetProduct(t *testing.T) {
	Convey("Given a product handler", t, func() {
		// Prepare mock repository
		mockRepo := &mockProductRepository{
			products: []entity.ProductEntity{
				{ID: 1, Title: "Product 1", Description: "Description 1", Price: 100.0, Category: "Category 1", Brand: "Brand 1", URL: "http://example.com/product1"},
			},
			err: nil,
		}

		service := service.NewProductService(mockRepo)
		productHandler := handler.NewProductHandler(service)

		Convey("When making a valid request", func() {
			ctx := &fasthttp.RequestCtx{}
			ctx.Request.SetRequestURI("/products/1")
			ctx.SetUserValue("id", "1")

			productHandler.GetProduct(ctx)

			Convey("Then the response should have a status code of 200", func() {
				So(ctx.Response.StatusCode(), ShouldEqual, fasthttp.StatusOK)
			})

			Convey("Then the response should have a content type of 'application/json'", func() {
				So(string(ctx.Response.Header.ContentType()), ShouldEqual, "application/json")
			})

			Convey("Then the response body should be a valid ProductDetailResponse", func() {
				var response dto.ProductDetailResponse
				err := json.Unmarshal(ctx.Response.Body(), &response)
				So(err, ShouldBeNil)

				expectedProduct := model.Product{
					ID:          1,
					Title:       "Product 1",
					Price:       100.0,
					Category:    "Category 1",
					Brand:       "Brand 1",
					URL:         "http://example.com/product1",
					Description: "Description 1",
				}
				So(response.Data, ShouldResemble, expectedProduct)
			})
		})

		Convey("When making a request with an invalid id parameter", func() {
			ctx := &fasthttp.RequestCtx{}
			ctx.Request.SetRequestURI("/products/invalid")
			ctx.SetUserValue("id", "invalid")

			productHandler.GetProduct(ctx)

			Convey("Then the response should have a status code of 400", func() {
				So(ctx.Response.StatusCode(), ShouldEqual, fasthttp.StatusBadRequest)
			})

			Convey("Then the response should have a content type of 'application/json'", func() {
				So(string(ctx.Response.Header.ContentType()), ShouldEqual, "application/json")
			})

			Convey("Then the response body should contain an error message", func() {
				var response map[string]string
				err := json.Unmarshal(ctx.Response.Body(), &response)
				So(err, ShouldBeNil)
				So(response["error"], ShouldEqual, "Invalid id")
			})
		})
	})
}
