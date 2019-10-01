package router

import (
	"jrnc-db/server/adapter/controllers"
	gin "github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := Router(false)
	r.Run()
}

// Router separate
func Router(gae bool) *gin.Engine {
	r := gin.New()

	v1 := r.Group("v1")
	{
		ctrl := controllers.Controllers{}

		// AutoMigration
		v1.GET("/autoMigration", ctrl.AutoMigration)

		// 購買取引先台帳登録
		v1.POST("/search", ctrl.GetByConditionsController)
		v1.POST("/insert", ctrl.CreateModelController)
		v1.PUT("/update/:id", ctrl.UpdateByIdController)
		v1.DELETE("/delete/:id", ctrl.DeleteByIdController)
		v1.GET("/csv", ctrl.CsvController)
	}
	return r
}
