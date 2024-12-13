package controller

import (
	"gorbac/app/request"
	"gorbac/app/service"
	"gorbac/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	service service.PermissionService
}

func NewPermissionController(s service.PermissionService) PermissionController {
	return PermissionController{
		service: s,
	}
}

// @Summary Get list permission
// @Description REST API permission
// @Author AzrielFatur
// @Success 200 {object} entity.Permission
// @Failure 404 {object} nil
// @method [GET]
// @Router /permissions
func (controller PermissionController) Index(ctx *gin.Context) {
	permissions, err := controller.service.List()

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Permission not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Permission Found", http.StatusOK, permissions)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one permission
// @Description REST API permission
// @Author AzrielFatur
// @Success 200 {object} entity.Permission
// @Failure 404 {object} nil
// @method [GET]
// @Router /permissions/:id
func (controller PermissionController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	permission, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Permission not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.ErrorJSON(ctx, "Permission Found", http.StatusOK, permission)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Insert permission
// @Description REST API permission
// @Author AzrielFatur
// @Success 200 {object} entity.Permission
// @Failure 400 {object} nil
// @method [POST]
// @Router /permissions
func (controller PermissionController) Store(ctx *gin.Context) {
	var permissionReq request.PermissionRequest

	if err := ctx.ShouldBindJSON(&permissionReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	permission, err := controller.service.Insert(permissionReq)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Permission", http.StatusOK, permission)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update permission
// @Description REST API permission
// @Author AzrielFatur
// @Success 200 {object} entity.Permission
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /permissions/:id
func (controller PermissionController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Permission not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var updateReq request.PermissionRequest

	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	permission, err := controller.service.Update(updateReq, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Permission", http.StatusOK, permission)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update permission
// @Description REST API permission
// @Author AzrielFatur
// @Success 200 {object} entity.Permission
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /permissions/:id
func (controller PermissionController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Permission not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	permission, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Permission", http.StatusOK, permission)

	ctx.JSON(http.StatusOK, resp)
}
