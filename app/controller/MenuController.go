package controller

import (
	"gorbac/app/request"
	"gorbac/app/service"
	"gorbac/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	service service.MenuService
}

func NewMenuController(s service.MenuService) MenuController {
	return MenuController{
		service: s,
	}
}

// @Summary Get list menu
// @Description REST API Menu
// @Author AzrielFatur
// @Success 200 {object} entity.User
// @Failure 404 {object} nil
// @method [GET]
// @Router /menus
func (controller MenuController) Index(ctx *gin.Context) {
	param, _ := strconv.ParseBool(ctx.Request.URL.Query().Get("parent_id"))

	menus, err := controller.service.List(param)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Menu not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Menu Found", http.StatusOK, menus)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get one menu
// @Description REST API menu
// @Author AzrielFatur
// @Success 200 {object} entity.Meny
// @Failure 404 {object} nil
// @method [GET]
// @Router /menus/:id
func (controller MenuController) Show(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	menu, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Menu not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Menu Found", http.StatusOK, menu)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Insert menu
// @Description REST API menu
// @Author AzrielFatur
// @Success 200 {object} entity.Menu
// @Failure 400 {object} nil
// @method [POST]
// @Router /menus
func (controller MenuController) Store(ctx *gin.Context) {
	var menuReq request.MenuRequest

	if err := ctx.ShouldBindJSON(&menuReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	menu, err := controller.service.Insert(menuReq)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to create menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Created Menu", http.StatusOK, menu)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Update menu
// @Description REST API menu
// @Author AzrielFatur
// @Success 200 {object} entity.Menu
// @Failure 400, 404 {object} err.Error(), nil
// @method [PUT]
// @Router /menus/:id
func (controller MenuController) Update(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Menu not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusOK, resp)

		return
	}

	var updateReq request.MenuRequest

	if err := ctx.ShouldBindJSON(&updateReq); err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	menu, err := controller.service.Update(updateReq, ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to update menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)
	}

	resp := helper.SuccessJSON(ctx, "Successfully to Update Menu", http.StatusOK, menu)

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete menu
// @Description REST API menu
// @Author AzrielFatur
// @Success 200 {object} entity.Menu
// @Failure 400, 404 {object} err.Error(), nil
// @method [DELETE]
// @Router /menus/:id
func (controller MenuController) Delete(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id")) // Get Param ID

	_, err := controller.service.FindById(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Menu not Found", http.StatusNotFound, nil)

		ctx.JSON(http.StatusNotFound, resp)

		return
	}

	menu, err := controller.service.Delete(ID)

	if err != nil {
		resp := helper.ErrorJSON(ctx, "Failed to delete menu", http.StatusBadRequest, err.Error())

		ctx.JSON(http.StatusBadRequest, resp)

		return
	}

	resp := helper.SuccessJSON(ctx, "Successfully Delete Menu", http.StatusOK, menu)

	ctx.JSON(http.StatusOK, resp)
}
