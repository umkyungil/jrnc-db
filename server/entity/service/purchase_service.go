package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jrnc-db/server/entity/models"
	"jrnc-db/server/entity/repository"
	"jrnc-db/server/entity/util"
	"net/http"
	"strconv"
)

// 複数行レスポンス
type PurchaseResults struct {
	Purchase				[]models.Purchase 	`json:"purchase"`
	ProcessResult int                   		`json:"process_result"`
}

// 単1行レスポンス
type PurchaseResult struct {
	Purchase	       		models.Purchase 	`json:"purchase"`
	ProcessResult int                   		`json:"process_result"`
}

// 検索条件
type PurchaseRequest struct {
	CustomerCode			string				`json:"customer_code"`
	CustomerName			string				`json:"customer_name"`
	SpecialiedField			string				`json:"specialied_field"`
	EvaluateDateFrom		string				`json:"evaluate_date_from"`
	EvaluateDateTo			string				`json:"evaluate_date_to"`
	ContinuedEvaluateResult	string				`json:"continued_evaluate_result"`
}

// 全件検索
func (s Service) GetAllModelService() (PurchaseResults, error) {
	var rep repository.Repository
	empty_res := []models.Purchase{}
	res, err := rep.GetAllModelRepository()

	// 戻り値設定
	if err != nil {
		var r =  PurchaseResults{empty_res,http.StatusNotFound}
		return r, err
	} else {
		var r= PurchaseResults{res, http.StatusOK}
		return r, err
	}
}

// 条件検索
func (s Service) GetByConditionsService(c *gin.Context) (PurchaseResults, error) {
	var rep repository.Repository
	var req PurchaseRequest
	empty_res := []models.Purchase{}
	c.BindJSON(&req)

	// 日付整合性チェック
	if err := util.CheckDate(req.EvaluateDateFrom); err != nil {
		var r =  PurchaseResults{empty_res,http.StatusBadRequest}
		return r, err
	}
	if err := util.CheckDate(req.EvaluateDateTo); err != nil {
		var r =  PurchaseResults{empty_res,http.StatusBadRequest}
		return r, err
	}
	if req.EvaluateDateFrom != "" && req.EvaluateDateTo != ""{
		f, _ := strconv.Atoi(req.EvaluateDateFrom)
		t, _ := strconv.Atoi(req.EvaluateDateTo)

		if f > t {
			var r =  PurchaseResults{empty_res,http.StatusBadRequest}
			return r, fmt.Errorf("Error: %s", "Date not valid")
		}
	}

	// 検索処理
	res, err := rep.GetByConditionsRepository(req.CustomerCode, req.CustomerName, req.SpecialiedField, req.EvaluateDateFrom, req.EvaluateDateTo, req.ContinuedEvaluateResult)

	// 戻り値処理
	if err != nil {
		var r = PurchaseResults{empty_res, http.StatusNotFound}
		return r, err
	} else if len(res) == 0 {
		var r =  PurchaseResults{res,http.StatusNoContent}
		return r, err
	} else {
		var r =  PurchaseResults{res,http.StatusOK}
		return r, err
	}
}

// ID検索
func (s Service) GetByIdService(c *gin.Context) (PurchaseResult, error) {
	var rep repository.Repository
	empty_res := models.Purchase{}
	res, err := rep.GetByIdRepository(c.Param("id"))

	// 戻り値設定
	if err != nil {
		var r =  PurchaseResult{empty_res,http.StatusNotFound}
		return r, err
	} else {
		var r= PurchaseResult{res, http.StatusOK}
		return r, err
	}
}

// 新規登録
func (s Service) CreateModelService(c *gin.Context) (PurchaseResult, error) {
	var rep repository.Repository
	var p models.Purchase
	empty_res := models.Purchase{}

	// TODO:特殊文字チェック
	// TODO:かなチェック

	// リクエストバインディングチェック
	if err := c.BindJSON(&p); err != nil {
		var r =  PurchaseResult{empty_res,http.StatusBadRequest}
		return r, err
	}
	// 日付整合性チェック
	if err := util.CheckDate(p.RequestDate); err != nil {
		var r =  PurchaseResult{empty_res,http.StatusBadRequest}
		return r, err
	}
	if err := util.CheckDate(p.EvaluateDate); err != nil {
		var r =  PurchaseResult{empty_res,http.StatusBadRequest}
		return r, err
	}
	if p.RequestDate != "" && p.EvaluateDate != ""{
		f, _ := strconv.Atoi(p.RequestDate)
		t, _ := strconv.Atoi(p.EvaluateDate)

		if f > t {
			var r =  PurchaseResult{empty_res,http.StatusBadRequest}
			return r, fmt.Errorf("Error: %s", "Date not valid")
		}
	}

	// 新規登録処理
	res, err := rep.CreateModelRepository(&p)

	// 戻り値処理
	if err != nil {
		var r =  PurchaseResult{empty_res,http.StatusInternalServerError}
		return r, err
	} else {
		var r= PurchaseResult{*res, http.StatusCreated}
		return r, err
	}
}

// 更新登録
func (s Service) UpdateByIdService(c *gin.Context) (PurchaseResult, error) {
	var rep repository.Repository
	var p models.Purchase
	empty_res := models.Purchase{}

	// リクエストバインディングチェック
	if err := c.BindJSON(&p); err != nil {
		var r =  PurchaseResult{empty_res,http.StatusBadRequest}
		return r, err
	}
	// 日付整合性チェック
	if err := util.CheckDate(p.RequestDate); err != nil {
		var r =  PurchaseResult{empty_res,http.StatusBadRequest}
		return r, err
	}
	if err := util.CheckDate(p.EvaluateDate); err != nil {
		var r =  PurchaseResult{empty_res,http.StatusBadRequest}
		return r, err
	}
	// TODO 特殊文字チェック
	// TODO かなチェック
	// 更新登録処理
	p, err := rep.UpdateByIdRepository(c.Param("id"), c)

	// 戻り値処理
	if err != nil {
		var r= PurchaseResult{p, http.StatusInternalServerError}
		return r, err
	} else {
		var r= PurchaseResult{p, http.StatusOK}
		return r, err
	}
}

// 削除
func (s Service) DeleteByIdService(c *gin.Context) (PurchaseResult, error) {
	var rep repository.Repository
	empty_res := models.Purchase{}

	// 削除処理
	err := rep.DeleteByIdRepository(c.Param("id"), c)

	// 戻り値処理
	if err != nil {
		var r= PurchaseResult{empty_res, http.StatusInternalServerError}
		return r, err
	} else {
		var r= PurchaseResult{empty_res, http.StatusNoContent}
		return r, err
	}
}

// CSV処理
func (s Service) CsvService(c *gin.Context) (PurchaseResults, error) {
	var rep repository.Repository
	var req PurchaseRequest
	empty_res := []models.Purchase{}

	// パラメータ取得
	customer_code := c.DefaultQuery("customer_code", "")
	customer_name := c.DefaultQuery("customer_name", "")
	specialied_field := c.DefaultQuery("specialied_field", "")
	evaluate_date_from := c.DefaultQuery("evaluate_date_from", "")
	evaluate_date_to := c.DefaultQuery("evaluate_date_to", "")
	continued_evaluate_result := c.DefaultQuery("continued_evaluate_result", "")

	// 日付整合性チェック
	if err := util.CheckDate(req.EvaluateDateFrom); err != nil {
		var r =  PurchaseResults{empty_res,http.StatusBadRequest}
		return r, err
	}
	if err := util.CheckDate(req.EvaluateDateTo); err != nil {
		var r =  PurchaseResults{empty_res,http.StatusBadRequest}
		return r, err
	}
	if req.EvaluateDateFrom != "" && req.EvaluateDateTo != ""{
		f, _ := strconv.Atoi(req.EvaluateDateFrom)
		t, _ := strconv.Atoi(req.EvaluateDateTo)

		if f > t {
			var r =  PurchaseResults{empty_res,http.StatusBadRequest}
			return r, fmt.Errorf("Error: %s", "Date not valid")
		}
	}

	// 検索処理
	res, err := rep.CSVRepository(customer_code, customer_name, specialied_field, evaluate_date_from, evaluate_date_to, continued_evaluate_result)
	if err != nil {
		var r = PurchaseResults{empty_res, http.StatusNotFound}
		return r, err
	} else if len(res) == 0 {
		var r= PurchaseResults{empty_res, http.StatusNoContent}
		return r, err
	}

	// CSV出力
	if err := util.PurchaseCSV(c, res); err != nil {
		var r =  PurchaseResults{empty_res,http.StatusInternalServerError}
		return r, err
	}

	// 戻り値処理
	var r =  PurchaseResults{res,http.StatusOK}
	return r, err
}