package controller

import (
	"gorbac/app/request"
	"gorbac/app/service"
	"gorbac/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatusController struct {
	service service.StatusService
}

func NewStatusController(s service.StatusService) StatusController {
	return StatusController{
		service: s,
	}
}

// @Summary Get status
// @Description REST API Status
// @Author Rasmad Ibnu
// @Success 200 {object} entity.Status
// @Failure 404 {object} nil
// @method [GET]
// @Router /status
func (controller StatusController) Index(ctx *gin.Context) {
	status, err := controller.service.List()

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Status not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Status Found", http.StatusOK, status)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary insert status
// @Description REST API Status
// @Author Rasmad Ibnu
// @Success 200 {object} entity.Status
// @Failure 400 {object} err.Error()
// @method [POST]
// @Router /status
func (controller StatusController) Store(ctx *gin.Context) {
	var req request.Status

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to crate status", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	status, err := controller.service.Insert(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create status", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created status", http.StatusOK, status)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one status
// @Description REST API status
// @Author Rasmad Ibnu
// @Success 200 {object} entity.Status
// @Failure 404 {object} nil
// @method [GET]
// @Router /status/:id
func (controller StatusController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	status, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Status not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Status Found", http.StatusOK, status)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update status
// @Description REST API status
// @Author Rasmad Ibnu
// @Success 200 {object} entity.Status
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /status/:id
func (controller StatusController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Status not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var req request.Status

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Status", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	status, err := controller.service.Update(req, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update Status", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Status", http.StatusOK, status)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete status
// @Description REST API status
// @Author Rasmad Ibnu
// @Success 200 {object} entity.Status
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /status/:id
func (controller StatusController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Status not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	status, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete Status", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Status", http.StatusOK, status)

	ctx.JSON(http.StatusOK, resp)
}
