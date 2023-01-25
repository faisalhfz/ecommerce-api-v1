package routes

import (
	"ecommerce-api/src/handler"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(g *echo.Group, transactionHandler *handler.TransactionHandler) {
	g.GET("/transactions", transactionHandler.GetTransactionsHandler)
}
