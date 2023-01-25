package handler

import (
	"ecommerce-api/src/entity/response"
	"ecommerce-api/src/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cUsecase *usecase.CartUsecase
}

func NewCartHandler(cUsecase *usecase.CartUsecase) *CartHandler {
	return &CartHandler{cUsecase: cUsecase}
}

func (cHandler CartHandler) GetCartHandler(ctx echo.Context) error {
	cartResponse, err := cHandler.cUsecase.GetCart()
	if err != nil {
		return ctx.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "Cart empty",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully get all products in cart",
		Data:    cartResponse,
	})
}

func (cHandler CartHandler) PostCartHandler(ctx echo.Context) error {
	_, err := cHandler.cUsecase.GetCart()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "No carts found",
			Data:    nil,
		})
	}
	transactionPIN := ctx.FormValue("pin")
	if len(transactionPIN) != 6 {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid transaction PIN",
			Data:    nil,
		})
	}
	cartResponse, err := cHandler.cUsecase.CheckoutCart()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to checkout cart",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Transaction completed",
		Data:    cartResponse,
	})
}

func (cHandler CartHandler) DeleteCartHandler(ctx echo.Context) error {
	_, err := cHandler.cUsecase.GetCart()
	// _ = cart
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "No carts found",
			Data:    nil,
		})
	}
	if err := cHandler.cUsecase.ClearCart(); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to delete cart",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully deleted cart",
		Data:    nil,
	})
}

// func (cHandler CartHandler) DeleteProductFromCartHandler(ctx echo.Context) error {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
// 			Code:    http.StatusBadRequest,
// 			Message: "Invalid product id",
// 			Data:    nil,
// 		})
// 	}
// 	cart, err := cHandler.cUsecase.GetCart()
// 	if err != nil {
// 		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
// 			Code:    http.StatusNotFound,
// 			Message: "No carts found",
// 			Data:    nil,
// 		})
// 	}
// 	if cHandler.cUsecase.IsProductInCart(id, cart) == false {
// 		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
// 			Code:    http.StatusNotFound,
// 			Message: "Product not found in cart",
// 			Data:    nil,
// 		})
// 	}
// 	if err := cHandler.cUsecase.RemoveProductFromCart(id); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
// 			Code:    http.StatusBadRequest,
// 			Message: "Failed to remove product from cart",
// 			Data:    nil,
// 		})
// 	}
// 	return ctx.JSON(http.StatusOK, response.BaseResponse{
// 		Code:    http.StatusOK,
// 		Message: "Successfully removed product from cart",
// 		Data:    nil,
// 	})
// }
