package orm

import (
	"back-end-server-dev/utils/connection"
	"database/sql"
	"gorm.io/gorm"
	"log"
)

type ObjectRelationalMapping struct {
	sql connection.DatabaseInterface
}

func NewObjectRelationalMapping(dsn *string) ObjectRelationalMapping {
	return ObjectRelationalMapping{sql: connection.DatabaseConnection{DsnSQL: *dsn}}
}

//go:generate mockery --name ObjectRelationalMappingInterface
type ObjectRelationalMappingInterface interface {
	Gorm() (*gorm.DB, error)
}

func (orm ObjectRelationalMapping) Gorm() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(orm.sql.MySql(), &gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
		return nil, err
	}

	var healthCheck *sql.DB
	healthCheck, err = db.DB()
	if err != nil {
		return nil, err
	}

	err = healthCheck.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}

//mysql.Open("golang-service-account:STRONG.password79@tcp(34.224.99.112:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=UTC"), &gorm.Config{}
