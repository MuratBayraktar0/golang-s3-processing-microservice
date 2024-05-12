package handler

import (
	"encoding/json"
	"strconv"

	"product_service/internal/domain/dto"
	"product_service/internal/domain/service"

	"github.com/valyala/fasthttp"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProducts(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	params := []string{"page", "size"}
	values := make([]int64, len(params))

	for i, param := range params {
		str := string(ctx.QueryArgs().Peek(param))
		val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			h.sendErrorResponse(ctx, fasthttp.StatusBadRequest, "Invalid "+param)
			return
		}
		values[i] = val
	}

	products, err := h.service.GetProducts(values[0], values[1])
	if err != nil {
		h.sendErrorResponse(ctx, fasthttp.StatusInternalServerError, "Failed to retrieve products")
		return
	}

	response := dto.ProductsResponse{
		Data: products,
		Page: values[0],
		Size: values[1],
	}

	h.sendSuccessResponse(ctx, response)
}

func (h *ProductHandler) GetProduct(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	id := string(ctx.UserValue("id").(string))
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.sendErrorResponse(ctx, fasthttp.StatusBadRequest, "Invalid id")
		return
	}

	product, err := h.service.GetProduct(idInt)
	if err != nil {
		h.sendErrorResponse(ctx, fasthttp.StatusInternalServerError, "Failed to retrieve product")
		return
	}

	response := dto.ProductDetailResponse{
		Data: product,
	}

	h.sendSuccessResponse(ctx, response)
}

func (h *ProductHandler) sendSuccessResponse(ctx *fasthttp.RequestCtx, response interface{}) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	jsonResponse, _ := json.Marshal(response)
	ctx.Write(jsonResponse)
}

func (h *ProductHandler) sendErrorResponse(ctx *fasthttp.RequestCtx, statusCode int, message string) {
	ctx.SetStatusCode(statusCode)
	jsonResponse, _ := json.Marshal(map[string]string{"error": message})
	ctx.Write(jsonResponse)
}
