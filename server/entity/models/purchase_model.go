package models

import (
	"time"
)

// 購買取引台帳マスター
type Purchase struct {
	Id							int64			`json:"id" gorm:"primary_key; AUTO_INCREMENT"`					// 購買取引先台帳マスターID
	CustommerNo					int				`json:"custommer_no" gorm:"size:4"`								// 顧客番号（現行システムカーラムで不使用）
	SpecialiedField         	string	        `json:"specialied_field" gorm:"size:100"`						// 取引先：専門分野
	Officer						string			`json:"officer" gorm:"size:100"`								// 取引先：担当者
	IntroducerType          	string	        `json:"introducer_type" gorm:"size:4"`							// 紹介者区分（１：無し、２：顧客、３：その他）
	Introducer              	string		 	`json:"introducer" gorm:"size:100"`								// 紹介者
	IntroducerLocated			string		 	`json:"introducer_located" gorm:"size:100"`						// 紹介者所属先
	RequestSection          	string	        `json:"request_section" gorm:"size:50"`							// 評価依頼部門
	RequestDate             	string	 		`json:"request_date" gorm:"size:8"`								// 評価依頼日
	EvaluateDate				string	 		`json:"evaluate_date" gorm:"size:8"`							// 評価実施日
	Evaluate1               	string	        `json:"evaluate1" gorm:"size:4"`								// 評価１区分（１：一般、２：個人）
	Evaluate1no1            	string	        `json:"evaluate1_1" gorm:"column:evaluate1_1; size:4"`			// 評価１ポイント（一般）
	Evaluate1no2            	string	        `json:"evaluate1_2" gorm:"column:evaluate1_2; size:4"`			// 評価１ポイント（個人商店など）
	Evaluate2               	string	        `json:"evaluate2" gorm:"size:4"`								// 評価２ポイント
	Evaluate3               	string	        `json:"evaluate3" gorm:"size:4"`								// 評価３ポイント
	EvaluatePoint				string			`json:"evaluate_point" gorm:"size:4"`							// 評価ポイント計
	EvaluateResult				string			`json:"evaluate_result" gorm:"size:4"`							// 判定結果（１：取引不可、２：取引可）
	ContractType				string			`json:"contract_type" gorm:"size:4"`							// 契約形態（１：基本契約、２：随時契約、３：１回限りの契約）
	ManagementDocument			string			`json:"management_document" gorm:"size:4"`						// 経営評価資料（１：無し、２：有り）
	Comment						string		   	`json:"comment" gorm:"size:500"`								// 特記事故

	// 継続評価情報
	ContinuedEvaluate1        	string			`json:"continued_evaluate1" gorm:"size:4"`						// 取組姿勢
	ContinuedEvaluate2          string      	`json:"continued_evaluate2" gorm:"size:4"`						// 専門技術
	ContinuedEvaluate3       	string      	`json:"continued_evaluate3" gorm:"size:4"`						// 照査体制
	ContinuedEvaluate4      	string      	`json:"continued_evaluate4" gorm:"size:4"`						// 成果レベル
	ContinuedEvaluatePoint  	string			`json:"continued_evaluate_point" gorm:"size:4"`					// 継続評価ポイント計
	ContinuedEvaluateResult 	string      	`json:"continued_evaluate_result" gorm:"size:4"`				// 注文停止・抹殺（１：注文停止、２：抹消）
	BasicContractDate       	string  		`json:"basic_contract_date" gorm:"size:8"`						// 基本契約日
	BusinessEvaluationDocument  string			`json:"business_evaluation_document" gorm:"size:4"`				// 業務評価無し（１：無し、２：あり）
	TradingField				string			`json:"trading_field" gorm:"size:100"`							// 取引分野
	TradingFieldDetails			string			`json:"trading_field_details" gorm:"size:200"`					// 取引分野細目

	//取引先情報
	CustomerCode				string		    `json:"customer_code" gorm:"size:4"`							// 取引先コード
	JapaneseSyllabary			string			`json:"japanese_syllabary" gorm:"size:4"`						// 五十音
	CustomerName            	string		    `json:"customer_name" gorm:"size:100"`							// 取引先名称
	CustomerNameK          		string        	`json:"customer_name_k" gorm:"size:100"`						// 取引先名称カナ
	RepresentativeName      	string          `json:"representative_name" gorm:"size:100"`					// 代表者名
	RepresentativeNameK     	string          `json:"representative_name_k" gorm:"size:100"`					// 代表者名カナ
	Address1                	string          `json:"address1" gorm:"size:200"`								// 住所1
	Address2                	string          `json:"address2" gorm:"size:200"`								// 住所2
	AddressK                	string	 	   	`json:"address_k" gorm:"size:200"`								// 住所カナ
	ZipCd                  		string          `json:"zip_cd" gorm:"size:10"`									// 郵便番号
	TelNo                   	string			`json:"tel_no" gorm:"size:14"`									// 電話番号
	FaxNo               		string        	`json:"fax_no" gorm:"size:14"`									// ファックス番号
	Capital                 	string          `json:"capital" gorm:"size:8"`									// 資本金（コンマ無し）
	ISO							string			`json:"iso" gorm:"size:4"`										// ISO有無（１：無し、２：あり）

	//その他情報
	DeletedDate					*time.Time		`json:"deleted_date" gorm:"size:8"`
	CreatedDate 				*time.Time  	`json:"created_date" gorm:"size:8"`
	UpdatedDate 				*time.Time 		`json:"updated_date" gorm:"size:8"`
}




