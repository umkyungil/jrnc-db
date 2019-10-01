package mysql

import (
	"database/sql"
	"log"
	"os"

	"jrnc-db/server/entity/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db    *gorm.DB
	err   error
	sqldb *sql.DB
)

// Init is initialize db from main function
func InitDev() {
	// 環境変数へ設定
	DBMS := "mysql"
	USER := "root"
	PASS := "!kium0612"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "jrnc-db"

	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+"?"+"parseTime=true"
	db, err = gorm.Open(DBMS, CONNECT)

	/*DBMS := os.Getenv("CRM_DBMS_DEV")
	CONNECTION := os.Getenv("CRM_DB_CONNECTION_DEV")
	db, err = gorm.Open(DBMS, CONNECTION)*/

	if err != nil {
		panic(err)
	}

	//単数形のテーブル作成設定
	db.SingularTable(true)
	//defer db.Close()
	db.LogMode(true)

	// テーブル初期化
	AutoMigration()

}

// InitGAE is initialize db(CloudSQL) from main function
func InitGAE() {
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("CLOUDSQL_USER")
	PASS := os.Getenv("CLOUDSQL_PASSWORD")
	CONNECTIONNAME := os.Getenv("CLOUDSQL_CONNECTION_NAME")
	DBNAME := os.Getenv("CLOUDSQL_DB_NAME")

	CONNECT := USER + ":" + PASS + "@unix(/cloudsql/" + CONNECTIONNAME + ")/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"
	var err error
	db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	// 開発環境はgormのデバッグモードにする
	if os.Getenv("GAE_ENV") == "dev" {
		db.LogMode(true)
	}
	AutoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// AutoMigration is initialize db(CloudSQL)
func AutoMigration() {
	db.AutoMigrate(&models.Purchase{}, &models.Client{}, &models.Owner{}, &models.OwnerClassType{}, &models.OwnerOrganType{})
}

