package controllers

import (
	"jrnc-db/server/driver/mysql"
	"github.com/gin-gonic/gin"
)

// AutoMigration action: GET /autoMigration
func (pc Controllers) AutoMigration(c *gin.Context) {
	mysql.AutoMigration()
}
