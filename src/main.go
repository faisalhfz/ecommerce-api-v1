package main

import (
	"ecommerce-api/src/config"
	"ecommerce-api/src/handler"
	"ecommerce-api/src/repository"
	"ecommerce-api/src/routes"
	"ecommerce-api/src/usecase"
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	config.AutoMigrate()

	productRepository := repository.NewProductRepository(config.DB)
	cartRepository := repository.NewCartRepository(config.DB)

	productUsecase := usecase.NewProductUsecase(productRepository, cartRepository)
	cartUsecase := usecase.NewCartUsecase(cartRepository, productRepository)

	productHandler := handler.NewProductHandler(productUsecase, cartUsecase)
	cartHandler := handler.NewCartHandler(cartUsecase)
	transactionHandler := handler.NewTransactionHandler(cartUsecase)

	e := echo.New()
	g := e.Group("/api/v1")

	g.GET("", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Welcome to PC Accessories Shop!")
	})

	routes.ProductRoutes(g, productHandler)
	routes.CartRoutes(g, cartHandler)
	routes.TransactionRoutes(g, transactionHandler)

	g.Static("/images", "src/files/images")

	e.Logger.Fatal(e.Start(fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("PORT"))))

}
