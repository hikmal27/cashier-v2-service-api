package controller

import (
	"gorbac/app/request"
	"gorbac/app/service"
	"gorbac/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
	role    service.UserRoleService
}

func NewUserController(s service.UserService, r service.UserRoleService) UserController {
	return UserController{
		service: s,
		role:    r,
	}
}

// @Summary Get list user
// @Description REST API User
// @Author AzrielFatur
// @Success 200 {object} entity.User
// @Failure 404 {object} nil
// @method [GET]
// @Router /users
func (controller UserController) Index(ctx *gin.Context) {
	users, err := controller.service.List()

	if err != nil {
		resp := helper.ErrorJSON(ctx, "No Data Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Data Found", http.StatusOK, users)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one user
// @Description REST API User
// @Author AzrielFatur
// @Success 200 {object} entity.User
// @Failure 404 {object} nil
// @method [GET]
// @Router /users/:id
func (controller UserController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	user, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "User not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "User Found", http.StatusOK, user)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update user
// @Description REST API User
// @Author AzrielFatur
// @Success 200 {object} entity.User
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /users/:id
func (controller UserController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "User not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	var updateReq request.UserUpdateRequest

	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update user", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	user, err := controller.service.Update(updateReq, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update user", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Updated User", http.StatusOK, user)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete user
// @Description REST API User
// @Author AzrielFatur
// @Success 200 {object} entity.User
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /users/:id
func (controller UserController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "User not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	user, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete user", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Deleted User", http.StatusOK, user)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Attach user has role
// @Description REST API User Role
// @Author AzrielFatur
// @Success 200 {object} entity.UserRole
// @Failure 400, 404 {object} err.Error(), nil
// @method [POST]
// @Router /users/roles
func (controller UserController) AttachRole(ctx *gin.Context) {
	var req request.UserHasRole

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	attach, err := controller.role.AttachRole(req)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to Attach role", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Attached Role", http.StatusOK, attach)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Detach user has role
// @Description REST API User Role
// @Author AzrielFatur
// @Success 200 {object} entity.UserRole
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /users/roles
// func (controller UserController) DetachRole(ctx *gin.Context) {
// 	var req request.UserHasRole

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		resp := helper.ErrorJSON(ctx, "Failed to Detach role", http.StatusBadRequest, err.Error())

// 		ctx.JSON(http.StatusBadRequest, resp)

// 		return
// 	}

// 	detach, err := controller.role.DetachRole(req)

// 	if err != nil {
// 		resp := helper.ErrorJSON(ctx, "Failed to Detach role", http.StatusBadRequest, err.Error())

// 		ctx.JSON(http.StatusBadRequest, resp)

// 		return
// 	}

// 	resp := helper.SuccessJSON(ctx, "Successfully Detached Role", http.StatusOK, detach)

// 	ctx.JSON(http.StatusOK, resp)
// }
