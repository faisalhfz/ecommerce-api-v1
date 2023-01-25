package handler

import (
	"ecommerce-api/src/entity/response"
	"ecommerce-api/src/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	cUsecase *usecase.CartUsecase
}

func NewTransactionHandler(cUsecase *usecase.CartUsecase) *TransactionHandler {
	return &TransactionHandler{cUsecase: cUsecase}
}

func (tHandler TransactionHandler) GetTransactionsHandler(ctx echo.Context) error {
	transactions, err := tHandler.cUsecase.GetCompletedCarts()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to get transactions",
			Data:    nil,
		})
	}
	if len(transactions) == 0 {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "No transactions found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully get all transactions",
		Data:    transactions,
	})
}
