package handler

import (
	"62tech.co/service/domain/v1/business"

	"62tech.co/service/domain/v1/system"

	"62tech.co/service/config"
	"62tech.co/service/config/database"
)

type Service struct {
	SystemModule   *system.Module
	BusinessModule *business.Module
}

func InitHandler() *Service {

	// mysql init
	MySQLConnection, err := database.MysqlDB()
	if err != nil {
		panic(err)
	}

	config := config.Configuration{
		MysqlDB: MySQLConnection,
	}

	// set service modular
	moduleSystem := system.InitModule(config)
	moduleBusiness := business.InitModule(config)

	return &Service{
		SystemModule:   moduleSystem,
		BusinessModule: moduleBusiness,
	}
}
