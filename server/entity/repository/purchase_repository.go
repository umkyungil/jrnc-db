package repository

import (
	"github.com/gin-gonic/gin"
	db "jrnc-db/server/driver/mysql"
	"jrnc-db/server/entity/models"
	"jrnc-db/server/entity/util"
	"time"
)

// 全件検索
func (r Repository)GetAllModelRepository() ([]models.Purchase, error) {
	db := db.GetDB()
	var res []models.Purchase

	// SQL作成
	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// ID検索
func (r Repository)GetByIdRepository(id string) (models.Purchase, error) {
	db := db.GetDB()
	var p models.Purchase

	// SQL作成
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

// 条件検索
func (r Repository)GetByConditionsRepository(customer_code string, customer_name string, specialied_field string, evaluate_date_from string,
							  evaluate_date_to string, continued_evaluate_result string) ([]models.Purchase, error) {
	db := db.GetDB()
	var res []models.Purchase

	// 現在日付取得
	today := time.Now().Format(util.FORMAT_DATE)
	t, _ := time.Parse(util.FORMAT_DATE, evaluate_date_to)
	base_date := t.AddDate(-1,0,0).Format(util.FORMAT_DATE)

	// SQL作成
	tx := db.Where("deleted_date IS NULL")
	// 取引先コード
	if customer_code != "" {
		tx = tx.Where("customer_code = ?", customer_code)
	}
	// 取引先名称
	if customer_name != "" {
		tx = tx.Where("customer_name LIKE ?", "%"+customer_name+"%")
	}
	// 専門分野
	if specialied_field != "" {
		tx = tx.Where("specialied_field LIKE ?", "%"+specialied_field+"%")
	}
	// 評価実施日
	if evaluate_date_from != "" && evaluate_date_to != ""{
		tx = tx.Where("evaluate_date BETWEEN ? AND ?", evaluate_date_from, evaluate_date_to)
	}
	// 評価実施日
	if evaluate_date_from != "" && evaluate_date_to == ""{
		tx = tx.Where("evaluate_date BETWEEN ? AND ?", evaluate_date_from, today)
	}
	// 評価実施日（１年前〜）
	if evaluate_date_from == "" && evaluate_date_to != ""{
		tx = tx.Where("evaluate_date BETWEEN ? AND ?", base_date, evaluate_date_to)
	}
	// 注文停止・抹殺
	if continued_evaluate_result != "" {
		tx = tx.Where("continued_evaluate_result = ?", continued_evaluate_result)
	}

	if err := tx.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

// 新規登録
func (r Repository)CreateModelRepository(req *models.Purchase) (*models.Purchase, error) {
	db := db.GetDB()

	// 作成日付設定
	time := time.Now()
	req.CreatedDate = &time

	res := db.Create(&req)
	if res.Error != nil {
		return req, res.Error
	}

	return req, nil
}

// 更新登録
func (r Repository)UpdateByIdRepository(id string, c *gin.Context) (models.Purchase, error) {
	db := db.GetDB()
	var p models.Purchase

	// SQL作成
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	if err := c.BindJSON(&p); err != nil {
		return p, err
	}

	// 更新日付設定
	time := time.Now()
	p.UpdatedDate = &time

	result := db.Save(&p)
	if result.Error != nil {
		return p, result.Error
	}

	return p, nil
}

// 削除登録
func (r Repository)DeleteByIdRepository(id string, c *gin.Context) (error) {
	db := db.GetDB()
	var p models.Purchase

	// SQL作成
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return err
	}
	if err := c.BindJSON(&p); err != nil {
		return err
	}

	// 削除日付設定
	time := time.Now()
	p.DeletedDate = &time

	result := db.Save(&p)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// CSV検索
func (r Repository)CSVRepository(customer_code string, customer_name string, specialied_field string, evaluate_date_from string,
	evaluate_date_to string, continued_evaluate_result string) ([]models.Purchase, error) {
	db := db.GetDB()
	var res []models.Purchase

	// 現在日付取得
	today := time.Now().Format(util.FORMAT_DATE)
	t, _ := time.Parse(util.FORMAT_DATE, evaluate_date_to)
	base_date := t.AddDate(-1,0,0).Format(util.FORMAT_DATE)

	// SQL作成
	tx := db.Where("deleted_date IS NULL")
	// 取引先コード
	if customer_code != "" {
		tx = tx.Where("customer_code = ?", customer_code)
	}
	// 取引先名称
	if customer_name != "" {
		tx = tx.Where("customer_name LIKE ?", "%"+customer_name+"%")
	}
	// 専門分野
	if specialied_field != "" {
		tx = tx.Where("specialied_field LIKE ?", "%"+specialied_field+"%")
	}
	// 評価実施日
	if evaluate_date_from != "" && evaluate_date_to != ""{
		tx = tx.Where("evaluate_date BETWEEN ? AND ?", evaluate_date_from, evaluate_date_to)
	}
	// 評価実施日
	if evaluate_date_from != "" && evaluate_date_to == ""{
		tx = tx.Where("evaluate_date BETWEEN ? AND ?", evaluate_date_from, today)
	}
	// 評価実施日（１年前〜）
	if evaluate_date_from == "" && evaluate_date_to != ""{
		tx = tx.Where("evaluate_date BETWEEN ? AND ?", base_date, evaluate_date_to)
	}
	// 注文停止・抹殺
	if continued_evaluate_result != "" {
		tx = tx.Where("continued_evaluate_result = ?", continued_evaluate_result)
	}

	if err := tx.Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}