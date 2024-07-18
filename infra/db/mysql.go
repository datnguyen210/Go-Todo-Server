package db

import (
	"database/sql"
	"fmt"

	"datnguyen/todo/config"

	"github.com/go-sql-driver/mysql"
)

func InitMySQL(cfg *config.Config) (*sql.DB, error) {
	mysqlConfig := mysql.Config{
		User:   cfg.DBUser,
		Passwd: cfg.DBPassword,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort),
		DBName: cfg.DBName,
	}

	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, fmt.Errorf("db.Ping: %v", pingErr)
	}
	fmt.Println("Connected!")

	return db, nil
}
