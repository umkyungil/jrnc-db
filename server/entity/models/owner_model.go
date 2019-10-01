package models

import (
	"time"
)

// 発注元
type Owner struct {
	OwnerCode					int				`json:"owner_code" gorm:"primary_key"`				// 発注者コード
	OwnerName					string			`json:"owner_name" gorm:"size:60; not null"`					// 発注者名
	OwnerNameK         			string	        `json:"owner_name_k" gorm:"size:30"`				// 発注者カナ名
	ZipCd1                  	string          `json:"zip_cd1" gorm:"size:3"`						// 郵便番号１
	ZipCd2                  	string          `json:"zip_cd2" gorm:"size:4"`						// 郵便番号２
	Address                		string          `json:"address" gorm:"size:60"`						// 住所1
	OwnerClassType				int				`json:"owner_class_type" gorm:"size:3"`				// 発注分類区分
	OwnerOrganType				int				`json:"owner_organ_type" gorm:"size:3"`				// 発注機関
	DeletedDate					*time.Time		`json:"deleted_date" gorm:"size:8"`
	CreatedDate 				*time.Time  	`json:"created_date" gorm:"size:8"`
	UpdatedDate 				*time.Time 		`json:"updated_date" gorm:"size:8"`
}