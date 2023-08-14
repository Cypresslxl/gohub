package migrations

import (
	"database/sql"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		City         string `json:"type:varchar(10)"`
		Introduction string `gorm:"type:varchar(255)"`
		Avatar       string `gorm:"type:varchar(255);default:null"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&User{}, "City")
		migrator.DropColumn(&User{}, "Introduction")
		migrator.DropColumn(&User{}, "Avatar")
		//migrator.DropTable(&User{})
	}

	migrate.Add("2023_08_14_192128_add_fields_to_user", up, down)
}
