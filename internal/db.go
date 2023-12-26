package internal

import (
	"fmt"

	"golang.org/x/exp/slog"

	"github.com/yuu-senoo/task-log/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB(cnf config.DBConfig) {
	// Connect to the database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)
	slog.Info(fmt.Sprintf("connectiong ... (%s)\n", dsn))
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	slog.Info("Relational database connected.")
}

func MigrateDB() {
	// Migrate the schema
	db.Set("gorm:table_options", "ENGINE = InnoDB").
		AutoMigrate(
			&User{},
			&Project{},
			&Task{},
		)
	slog.Info("Successfully created table.")
}
