package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"goBackendExample/internal/model"
)

var db *gorm.DB

func Setup() error {
	var err error

	db, err = gorm.Open(postgres.New(
		postgres.Config{
			DSN: "host=localhost port=5432 user=postgres dbname=example_db password=postgres sslmode=disable " +
				"TimeZone=Asia/Shanghai",
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)

	if err != nil {
		log.Panic(err)
		return err
	}

	// auto migrate
	db.AutoMigrate(&model.User{})

	return nil
}
