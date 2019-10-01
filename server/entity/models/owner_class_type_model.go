package models

import (
	"time"
)

// 発注分類区分
type OwnerClassType struct {
	TypeCode			int				`json:"type_code" gorm:"primary_key"`			// 発注分類区分
	TypeName			string			`json:"type_name" gorm:"size:60; not null"`		// 発注分類名
	DeletedDate			*time.Time		`json:"deleted_date" gorm:"size:8"`
	CreatedDate 		*time.Time  	`json:"created_date" gorm:"size:8"`
	UpdatedDate 		*time.Time 		`json:"updated_date" gorm:"size:8"`
}