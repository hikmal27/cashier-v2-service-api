package controller

import (
	"gorbac/app/request"
	"gorbac/app/service"
	"gorbac/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	service    service.RoleService
	menu       service.RoleMenuService
	permission service.RolePermissionService
}

func NewRoleController(s service.RoleService, m service.RoleMenuService, p service.RolePermissionService) RoleController {
	return RoleController{
		service:    s,
		menu:       m,
		permission: p,
	}
}

// @Summary Get list role
// @Description REST API Role
// @Author AzrielFatur
// @Success 200 {object} entity.User
// @Failure 404 {object} nil
// @method [GET]
// @Router /roles
func (controller RoleController) Index(ctx *gin.Context) {
	roles, err := controller.service.List()

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Role Found", http.StatusOK, roles)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one role
// @Description REST API role
// @Author AzrielFatur
// @Success 200 {object} entity.Role
// @Failure 404 {object} nil
// @method [GET]
// @Router /roles/:id
func (controller RoleController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	role, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.ErrorJSON(ctx, "Role Found", http.StatusOK, role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Insert role
// @Description REST API role
// @Author AzrielFatur
// @Success 200 {object} entity.Role
// @Failure 400 {object} nil
// @method [POST]
// @Router /roles
func (controller RoleController) Store(ctx *gin.Context) {
	var roleReq request.RoleRequest

	if err := ctx.ShouldBindJSON(&roleReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	role, err := controller.service.Insert(roleReq)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Role", http.StatusOK, role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update role
// @Description REST API role
// @Author AzrielFatur
// @Success 200 {object} entity.Role
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /roles/:id
func (controller RoleController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var updateReq request.RoleRequest

	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	role, err := controller.service.Update(updateReq, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Role", http.StatusOK, role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update role
// @Description REST API role
// @Author AzrielFatur
// @Success 200 {object} entity.Role
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /roles/:id
func (controller RoleController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Role not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	role, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Role", http.StatusOK, role)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Attach role has menu
// @Description REST API Role Menu
// @Author AzrielFatur
// @Success 200 {object} entity.RoleMenu
// @Failure 400, 404 {object} err.Error(), nil
// @method [POST]
// @Router /roles/:id/menus
func (controller RoleController) AttachMenu(ctx *gin.Context) {
	var req request.RoleHasMenu

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	attach, err := controller.menu.AttachMenu(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Attached Menu", http.StatusOK, attach)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Detach role has menus
// @Description REST API Role Menu
// @Author AzrielFatur
// @Success 200 {object} entity.RoleMenu
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /roles/menus
func (controller RoleController) DetachMenu(ctx *gin.Context) {
	var req request.RoleHasMenu

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Detach menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	detach, err := controller.menu.DetachMenu(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Detach menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Detached Role", http.StatusOK, detach)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Attach role has permissions
// @Description REST API Role Permission
// @Author AzrielFatur
// @Success 200 {object} entity.RolePermissions
// @Failure 400, 404 {object} err.Error(), nil
// @method [POST]
// @Router /roles/:id/permissions
func (controller RoleController) AttachPermission(ctx *gin.Context) {
	var req request.RoleHasPermission

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	attach, err := controller.permission.AttachPermission(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Attached Permission", http.StatusOK, attach)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Detach role has permissions
// @Description REST API Role Permission
// @Author AzrielFatur
// @Success 200 {object} entity.RolePermission
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /roles/permissions
func (controller RoleController) DetachPermission(ctx *gin.Context) {
	var req request.RoleHasPermission

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Detach permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	detach, err := controller.permission.DetachPermission(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Detach permission", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Detached Permission", http.StatusOK, detach)

	ctx.JSON(http.StatusOK, resp)
}
