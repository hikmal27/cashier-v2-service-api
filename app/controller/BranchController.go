package controller

import (
	"gorbac/app/service"
	"gorbac/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BranchController struct {
	service service.BranchService
}

func NewBranchController(s service.BranchService) BranchController {
	return BranchController{
		service: s,
	}
}

// @Summary Get branch
// @Description REST API all branch
// @Author Rasmad Ibnu
// @Success 200 {object} entity.Branch
// @Failure 404 {object} nil
// @method [GET]
// @Router /branch
func (controller BranchController) Index(ctx *gin.Context) {
	param := ctx.Request.URL.Query()

	m := make(map[string]interface{})
	for k, v := range param {
		m[k] = v
	}

	branch, err := controller.service.FindAll(m)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Branch not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Branch Found", http.StatusOK, branch)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one branch
// @Description REST API branch by ID
// @Author Rasmad Ibnu
// @Success 200 {object} entity.Branch
// @Failure 404 {object} nil
// @method [GET]
// @Router /branch/:id
func (controller BranchController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	branch, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Branch not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Branch Found", http.StatusOK, branch)

	ctx.JSON(http.StatusOK, resp)
}
