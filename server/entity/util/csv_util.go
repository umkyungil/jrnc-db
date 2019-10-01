package util

import (
	"bytes"
	"encoding/csv"
	_ "github.com/djimenez/iconv-go"
	"github.com/gin-gonic/gin"
	"jrnc-db/server/entity/models"
	"net/http"
	"time"
)

// CSV作成
func PurchaseCSV(c *gin.Context, data []models.Purchase) error{
	headers := []string{"専門分野","担当者","紹介者区分","紹介者","紹介者所属先","評価依頼部門","評価依頼日","評価実施日",
		"評価１区分","評価１ポイント（一般）","評価１ポイント（個人商店など）","評価２ポイント","評価３ポイント","評価ポイント計","判定結果",
		"契約形態","経営評価資料","特記事故","取組姿勢","専門技術","照査体制","成果レベル","継続評価ポイント計","注文停止・抹殺","基本契約日","業務評価無し",
		"取引分野","取引分野細目","取引先コード","五十音","取引先名称","取引先名称カナ","代表者名","代表者名カナ","住所1","住所2","住所カナ","郵便番号",
		"電話番号","ファックス番号","資本金","ISO","削除日","作成日","更新日"}

	b := &bytes.Buffer{}
	w := csv.NewWriter(b)

	// ヘッダ作成
	if err := w.Write(headers); err != nil {
		return  err
	}

	// データ作成
	for _, purchase := range data {
		var record []string
		record = append(record, purchase.SpecialiedField)
		record = append(record, purchase.Officer)
		// 紹介者区分（１：無し、２：顧客、３：その他）
		if purchase.IntroducerType != "" {
			if purchase.IntroducerType == ONE {
				record = append(record, NOT_EXIST)
			} else if purchase.IntroducerType == TWO {
				record = append(record, "顧客")
			} else if purchase.IntroducerType == THREE {
				record = append(record, "その他")
			}
		}
		record = append(record, purchase.Introducer)
		record = append(record, purchase.IntroducerLocated)
		record = append(record, purchase.RequestSection)
		record = append(record, purchase.RequestDate)
		record = append(record, purchase.EvaluateDate)
		// 評価１区分（１：一般、２：個人）
		if purchase.Evaluate1 != "" {
			if purchase.Evaluate1 == ONE {
				record = append(record, "一般")
			} else if purchase.Evaluate1 == TWO {
				record = append(record, "個人")
			}
		}
		record = append(record, purchase.Evaluate1no1)
		record = append(record, purchase.Evaluate1no2)
		record = append(record, purchase.Evaluate2)
		record = append(record, purchase.Evaluate3)
		record = append(record, purchase.EvaluatePoint)
		// 判定結果（１：取引不可、２：取引可）
		if purchase.EvaluateResult != "" {
			if purchase.EvaluateResult == ONE {
				record = append(record, "取引不可")
			} else if purchase.EvaluateResult == TWO {
				record = append(record, "取引可")
			}
		}
		// 契約形態（１：基本契約、２：随時契約、３：１回限りの契約）
		if purchase.ContractType != "" {
			if purchase.ContractType == ONE {
				record = append(record, "基本契約")
			} else if purchase.ContractType == TWO {
				record = append(record, "随時契約")
			} else if purchase.ContractType == THREE {
				record = append(record, "１回限りの契約")
			}
		}
		// 経営評価資料（１：無し、２：有り）
		if purchase.ManagementDocument != "" {
			if purchase.ManagementDocument == ONE {
				record = append(record, NOT_EXIST)
			} else if purchase.ManagementDocument == TWO {
				record = append(record, EXIST)
			}
		}
		record = append(record, purchase.Comment)
		record = append(record, purchase.ContinuedEvaluate1)
		record = append(record, purchase.ContinuedEvaluate2)
		record = append(record, purchase.ContinuedEvaluate3)
		record = append(record, purchase.ContinuedEvaluate4)
		record = append(record, purchase.ContinuedEvaluatePoint)
		// 注文停止・抹殺（１：注文停止、２：抹消）
		if purchase.ContinuedEvaluateResult != "" {
			if purchase.ContinuedEvaluateResult == ONE {
				record = append(record, "注文停止")
			} else if purchase.ContinuedEvaluateResult == TWO {
				record = append(record, "抹消")
			}
		}
		record = append(record, purchase.BasicContractDate)
		// 業務評価無し（１：無し、２：有り）
		if purchase.BusinessEvaluationDocument != "" {
			if purchase.BusinessEvaluationDocument == ONE {
				record = append(record, NOT_EXIST)
			} else if purchase.BusinessEvaluationDocument == TWO {
				record = append(record, EXIST)
			}
		}
		record = append(record, purchase.TradingField)
		record = append(record, purchase.TradingFieldDetails)
		record = append(record, purchase.CustomerCode)
		record = append(record, purchase.JapaneseSyllabary)
		record = append(record, purchase.CustomerName)
		record = append(record, purchase.CustomerNameK)
		record = append(record, purchase.RepresentativeName)
		record = append(record, purchase.RepresentativeNameK)
		record = append(record, purchase.Address1)
		record = append(record, purchase.Address2)
		record = append(record, purchase.AddressK)
		record = append(record, purchase.ZipCd)
		record = append(record, purchase.TelNo)
		record = append(record, purchase.FaxNo)
		record = append(record, purchase.Capital)
		// ISO有無（１：無し、２：有り）
		if purchase.ISO != "" {
			if purchase.ISO == ONE {
				record = append(record, NOT_EXIST)
			} else if purchase.ISO == TWO {
				record = append(record, EXIST)
			}
		}
		if purchase.DeletedDate == nil {
			record = append(record, "")
		} else {
			record = append(record, purchase.DeletedDate.Format(FORMAT_DATE))
		}
		if purchase.CreatedDate == nil{
			record = append(record, "")
		} else {
			record = append(record, purchase.CreatedDate.Format(FORMAT_DATE))
		}
		if purchase.UpdatedDate == nil{
			record = append(record, "")
		} else {
			record = append(record, purchase.UpdatedDate.Format(FORMAT_DATE))
		}

		if err := w.Write(record); err != nil {

		}
	}
	w.Flush()

	//ファイル名作成（ミリセコンド）
	filename := CSV_FILENAME + time.Now().Format(MILLI_FORMAT) + EXTENSION

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "text/csv", b.Bytes())

	return nil
}

