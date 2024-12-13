package controller

import (
	"gorbac/app/entity"
	"gorbac/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

type ProductsController struct {
	service service.ProductsService
}

func NewProductsController(s service.ProductsService) ProductsController {
	return ProductsController{
		service: s,
	}
}

func (controller *ProductsController) Index(ctx *gin.Context) {
	pg := paginate.New()
	products := controller.service.FindAll()

	pageData := pg.With(products).Request(ctx.Request).Response(&[]entity.Products{})

	ctx.JSON(http.StatusOK, pageData)
}
