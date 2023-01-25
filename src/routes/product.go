package routes

import (
	"ecommerce-api/src/handler"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(g *echo.Group, productHandler *handler.ProductHandler) {
	g.POST("/products", productHandler.PostProductHandler)
	g.GET("/products", productHandler.GetProductsHandler)
	g.GET("/products/:id", productHandler.GetProductByIdHandler)
	g.PUT("/products/:id", productHandler.PutProductByIdHandler)
	g.DELETE("/products/:id", productHandler.DeleteProductByIdHandler)
	g.POST("/products/:id/cart", productHandler.PostProductToCartByIdHandler)
	// g.PUT("/products/:id/cart", productHandler.PutProductToCartByIdHandler)
	// g.DELETE("/products/:id/cart", productHandler.DeleteProductToCartByIdHandler)
}
