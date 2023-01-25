package handler

import (
	"ecommerce-api/src/entity/request"
	"ecommerce-api/src/entity/response"
	"ecommerce-api/src/usecase"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	pUsecase *usecase.ProductUsecase
	cUsecase *usecase.CartUsecase
}

func NewProductHandler(pUsecase *usecase.ProductUsecase, cUsecase *usecase.CartUsecase) *ProductHandler {
	return &ProductHandler{pUsecase: pUsecase, cUsecase: cUsecase}
}

func (pHandler ProductHandler) PostProductHandler(ctx echo.Context) error {
	name := ctx.FormValue("name")
	image, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid file header",
			Data:    nil,
		})
	}
	imgPath, err := FileHandler(image)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Error opening file",
			Data:    nil,
		})
	}
	price, err := strconv.Atoi(ctx.FormValue("price"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid price",
			Data:    nil,
		})
	}
	category := ctx.FormValue("category")
	description := ctx.FormValue("description")

	productRequest := request.CreateProductRequest{Name: name, Image: imgPath, Price: price, Category: category, Description: description}
	productId, err := pHandler.pUsecase.CreateProduct(productRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to create product",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Product created successfully",
		Data:    productId,
	})
}

func (pHandler ProductHandler) GetProductsHandler(ctx echo.Context) error {
	category := ctx.FormValue("category")
	minPrice, _ := strconv.Atoi(ctx.FormValue("minprice"))
	maxPrice, _ := strconv.Atoi(ctx.FormValue("maxprice"))
	filter := request.ProductFilterRequest{Category: category, MinPrice: minPrice, MaxPrice: maxPrice}
	products, _ := pHandler.pUsecase.GetProducts(filter)
	if len(products) == 0 {
		return ctx.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "No products found",
			Data:    products,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully get all products",
		Data:    products,
	})
}

func (pHandler ProductHandler) GetProductByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
		})
	}
	product, err := pHandler.pUsecase.GetProductById(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Product id not found",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully get product",
		Data:    product,
	})
}

func (pHandler ProductHandler) PutProductByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
		})
	}
	product, err := pHandler.pUsecase.GetProductById(id)
	_ = product
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Product id not found",
			Data:    nil,
		})
	}

	name := ctx.FormValue("name")
	image, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid file header",
			Data:    nil,
		})
	}
	imgPath, err := FileHandler(image)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Error opening file",
			Data:    nil,
		})
	}
	price, err := strconv.Atoi(ctx.FormValue("price"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid price",
			Data:    nil,
		})
	}
	category := ctx.FormValue("category")
	description := ctx.FormValue("description")

	productRequest := request.CreateProductRequest{Name: name, Image: imgPath, Price: price, Category: category, Description: description}
	if err := pHandler.pUsecase.EditProductById(id, productRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to update product",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully updated product",
		Data:    nil,
	})
}

func (pHandler ProductHandler) DeleteProductByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
		})
	}
	product, err := pHandler.pUsecase.GetProductById(id)
	_ = product
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Product id not found",
			Data:    nil,
		})
	}

	if err := pHandler.pUsecase.DeleteProductById(id); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to delete product",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully deleted product",
		Data:    nil,
	})
}

func (pHandler ProductHandler) PostProductToCartByIdHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid product id",
			Data:    nil,
		})
	}
	_, err = pHandler.pUsecase.GetProductById(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Product id not found",
			Data:    nil,
		})
	}
	quantity, err := strconv.Atoi(ctx.FormValue("quantity"))
	if err != nil {
		quantity = 1
	}

	_, err = pHandler.cUsecase.GetCart()
	if err != nil {
		_, err := pHandler.cUsecase.CreateCart()
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "Failed to create cart",
				Data:    nil,
			})
		}
	}
	cartId, err := pHandler.pUsecase.AddProductToCart(id, quantity)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to add product to cart",
			Data:    nil,
		})
	}
	return ctx.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully added product to cart",
		Data:    cartId,
	})
}

func FileHandler(image *multipart.FileHeader) (string, error) {
	src, err := image.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	path := filepath.Join("src", "files", "images", image.Filename)
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return fmt.Sprintf("http://%v:%v/api/v1/images/%v", os.Getenv("HOST"), os.Getenv("PORT"), image.Filename), nil
}
