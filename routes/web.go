package routes

import (
	"gorbac/app/controller"
	"gorbac/app/repository"
	"gorbac/app/service"
	"gorbac/config"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func WebRouter(db config.Database) {
	// Repository Asset
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)
	roleMenuRepo := repository.NewRoleMenuRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)
	rolePermissionRepo := repository.NewRolePermissionRepository(db)
	statusRepo := repository.NewStatusRepository(db)
	branchAreaRepo := repository.NewAreaRepository(db)
	branchRepo := repository.NewBranchRepository(db)
	productRepo := repository.NewProductsRepository(db)

	// Service Asset
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	roleService := service.NewRoleService(roleRepo)
	menuService := service.NewMenuService(menuRepo)
	userRoleService := service.NewUserRoleService(userRoleRepo)
	roleMenuService := service.NewRoleMenuService(roleMenuRepo)
	permissionService := service.NewPermissionService(permissionRepo)
	rolePermissionService := service.NewRolePermissionService(rolePermissionRepo)
	statusService := service.NewStatusService(statusRepo)
	branchAreaService := service.NewAreaService(branchAreaRepo)
	branchService := service.NewBranchService(branchRepo)
	productService := service.NewProductsService(productRepo)

	//Controller Asset
	userController := controller.NewUserController(userService, userRoleService)
	authController := controller.NewAuthController(userService, authService)
	roleController := controller.NewRoleController(roleService, roleMenuService, rolePermissionService)
	menuController := controller.NewMenuController(menuService)
	permController := controller.NewPermissionController(permissionService)
	statusController := controller.NewStatusController(statusService)
	branchAreaController := controller.NewAreaController(branchAreaService)
	branchController := controller.NewBranchController(branchService)
	productController := controller.NewProductsController(productService)

	// Route
	httpRouter := gin.Default()

	// Register routing
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Testing  connection
	httpRouter.GET("/status-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Service ✅ API Up and Running"})
	})

	httpRouter.POST("/api/v1/auth/login", authController.Login)

	// Assets
	httpRouter.Static("/api/v1/assets/files", "./files")
	v1 := httpRouter.Group("/api/v1") // Grouping routes
	// v1.Use(security.JwtAuthMiddleware())

	// Auth routes
	v1.POST("/auth/token", authController.RefreshToken)

	// Users routes
	v1.GET("/users", userController.Index)
	v1.POST("/users", authController.Register)
	v1.GET("/users/:id", userController.Show)
	v1.PUT("/users/:id", userController.Update)
	v1.DELETE("/users/:id", userController.Delete)

	// User role routes
	v1.POST("/users/roles/assign", userController.AttachRole)
	// v1.POST("/users/roles/delete", userController.DetachRole)

	// Roles routes
	v1.GET("/roles", roleController.Index)
	v1.POST("/roles", roleController.Store)
	v1.GET("/roles/:id", roleController.Show)
	v1.PUT("/roles/:id", roleController.Update)
	v1.DELETE("/roles/:id", roleController.Delete)

	// Role menu routes
	v1.POST("/roles/menus/assign", roleController.AttachMenu)
	// v1.POST("/roles/menus/unassign", roleController.DetachMenu)

	// Permission routes
	v1.POST("/permissions", permController.Store)
	v1.GET("/permissions", permController.Index)
	v1.GET("/permissions/:id", permController.Show)
	v1.PUT("/permissions/:id", permController.Update)
	v1.DELETE("/permissions/:id", permController.Delete)

	// Role permission routes
	v1.POST("/roles/permissions/assign", roleController.AttachPermission)
	// v1.DELETE("/roles/permissions", roleController.DetachPermission)

	// Menus routes
	v1.GET("/menus", menuController.Index)
	v1.POST("/menus", menuController.Store)
	v1.GET("/menus/:id", menuController.Show)
	v1.PUT("/menus/:id", menuController.Update)
	v1.DELETE("/menus/:id", menuController.Delete)

	// Status routes
	v1.GET("/status", statusController.Index)
	v1.POST("/status", statusController.Store)
	v1.GET("/status/:id", statusController.Show)
	v1.PUT("/status/:id", statusController.Update)
	v1.DELETE("/status/:id", statusController.Delete)

	// Branch Area routes
	v1.GET("/area", branchAreaController.Index)
	v1.GET("/area/:id", branchAreaController.Show)

	// Branch routes
	v1.GET("/branch", branchController.Index)
	v1.GET("/branch/:id", branchController.Show)

	// products
	v1.GET("/products", productController.Index)

	httpRouter.Run(":" + os.Getenv("APP_PORT")) // rRun Routes with PORT
}
