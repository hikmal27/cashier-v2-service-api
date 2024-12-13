package controller

import (
	"gorbac/app/service"
	"gorbac/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AreaController struct {
	service service.AreaService
}

func NewAreaController(s service.AreaService) AreaController {
	return AreaController{
		service: s,
	}
}

// @Summary Get area
// @Description REST API area by id
// @Author Rasmad Ibnu
// @Success 200 {object} entity.BranchArea
// @Failure 404 {object} nil
// @method [GET]
// @Router /area
func (controller AreaController) Index(ctx *gin.Context) {
	area, err := controller.service.FindAll()

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Branch Area not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Branch Area Found", http.StatusOK, area)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one branch area
// @Description REST API branch area
// @Author Rasmad Ibnu
// @Success 200 {object} entity.BranchArea
// @Failure 404 {object} nil
// @method [GET]
// @Router /area/:id
func (controller AreaController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	area, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Branch Area not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Branch Area Found", http.StatusOK, area)

	ctx.JSON(http.StatusOK, resp)
}
