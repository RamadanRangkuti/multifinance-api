package main

import (
	"database/sql"

	"github.com/RamadanRangkuti/multifinance-api/config"
	userHandler "github.com/RamadanRangkuti/multifinance-api/handlers/users"
	"github.com/RamadanRangkuti/multifinance-api/repository/users"
	"github.com/RamadanRangkuti/multifinance-api/routes"
	userSvc "github.com/RamadanRangkuti/multifinance-api/usecase/users"
	"github.com/go-playground/validator"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	dbConn, err := config.ConnectToDatabase(config.Connection{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})
	if err != nil {
		return
	}
	defer dbConn.Close()

	validator := validator.New()

	routes := setupRoutes(dbConn, validator)
	routes.Run(cfg.AppPort)
}

func setupRoutes(db *sql.DB, validator *validator.Validate) *routes.Routes {
	userStore := users.NewStore(db)
	userSvc := userSvc.NewUserSvc(userStore)
	userHandler := userHandler.NewHandler(userSvc, validator)

	return &routes.Routes{
		User: userHandler,
	}
}
