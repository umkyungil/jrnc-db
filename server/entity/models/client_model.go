package models

import (
	"time"
)

// 取引先
type Client struct {
	ClientCode					string		    `json:"client_code" gorm:"primary_key; size:5"`		// 取引先コード
	ClientName            		string		    `json:"client_name" gorm:"size:40"`					// 取引先名称
	ClientNameK          		string        	`json:"client_name_k" gorm:"size:30"`				// 取引先名称カナ
	ZipCd1                  	string          `json:"zip_cd1" gorm:"size:3"`						// 郵便番号１
	ZipCd2                  	string          `json:"zip_cd2" gorm:"size:4"`						// 郵便番号２
	Address                		string          `json:"address1" gorm:"size:60"`					// 住所
	ContractType				string			`json:"contract_type" gorm:"size:50"`				// 契約内容（１：取引不可、２：取引可）
	OrderFlag					string			`json:"order_flag" gorm:"size:50"`					// 注文FLG（０：注文不可、２：注文可）
	DeletedDate					*time.Time		`json:"deleted_date" gorm:"size:8"`
	CreatedDate 				*time.Time  	`json:"created_date" gorm:"size:8"`
	UpdatedDate 				*time.Time 		`json:"updated_date" gorm:"size:8"`
}




