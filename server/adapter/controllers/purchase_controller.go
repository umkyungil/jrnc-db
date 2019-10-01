package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jrnc-db/server/entity/service"
	"net/http"
)

// 条件検索: POST /search
func (pc Controllers) GetByConditionsController(c *gin.Context) {
	var s service.Service
	p, err := s.GetByConditionsService(c)

	if err != nil || p.ProcessResult == http.StatusNoContent{
		c.AbortWithStatus(p.ProcessResult)
	} else {
		c.JSON(p.ProcessResult, p)
	}
}

// 条件検索: GET /search/:id
func (pc Controllers) GetByIdController(c *gin.Context) {
	var s service.Service
	p, err := s.GetByIdService(c)

	if err != nil || p.ProcessResult == http.StatusNoContent{
		c.AbortWithStatus(p.ProcessResult)
	} else {
		c.JSON(p.ProcessResult, p)
	}
}

// 新規登録: POST /insert
func (pc Controllers) CreateModelController(c *gin.Context) {
	var s service.Service
	p, err := s.CreateModelService(c)

	if err != nil {
		c.AbortWithStatus(p.ProcessResult)
	} else {
		c.JSON(p.ProcessResult, p)
	}
}

// 更新: PUT /update/:id
func (pc Controllers) UpdateByIdController(c *gin.Context) {
	var s service.Service
	p, err := s.UpdateByIdService(c)

	if err != nil {
		c.AbortWithStatus(p.ProcessResult)
	} else {
		c.JSON(p.ProcessResult, p)
	}
}

// 削除: DELETE /delete/:id
func (pc Controllers) DeleteByIdController(c *gin.Context) {
	var s service.Service
	p, err := s.DeleteByIdService(c)

	if err != nil {
		c.AbortWithStatus(p.ProcessResult)
	} else {
		c.JSON(p.ProcessResult, p)
	}
}

// CSV: GET /csv
func (pc Controllers)CsvController (c *gin.Context) {
	var s service.Service
	p, err := s.CsvService(c)
	if err != nil || p.ProcessResult == http.StatusNoContent {
		c.AbortWithStatus(p.ProcessResult)
	}
}

// Batch Job Scheduller
func (pc Controllers)JobSchedullerController () {
	fmt.Println("test : ")
}