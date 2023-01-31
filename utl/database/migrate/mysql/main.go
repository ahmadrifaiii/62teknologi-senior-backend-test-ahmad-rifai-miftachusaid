package mysql

import (
	"database/sql"
	"fmt"
	"net/http"

	"62tech.co/service/utl/response"

	"62tech.co/service/config/env"

	"62tech.co/service/model"
	"github.com/labstack/echo/v4"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/gommon/log"
)

// @Summary Migration
// @Description get the status of server.
// @Tags System Server
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /migration [get]
func Migration(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
		resp      = make(map[string]interface{})
	)

	// logic
	env.LoadEnv()

	condb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		env.Conf.DBUser,
		env.Conf.DBPass,
		env.Conf.DBHost,
		env.Conf.DBPort,
		env.Conf.DBName)

	db, err := sql.Open("mysql", condb)
	if err != nil {
		// log.Errorf("error database %v", err)
		defer db.Close()
		// os.Exit(1)
		return response.Success(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error database %v", err),
			Data:    "",
		})
	}
	err = db.Ping()
	if err != nil {
		log.Errorf("not connect database %v", err)
		// os.Exit(1)
		return response.Success(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("not connect database %v", err),
			Data:    nil,
		})
	}

	resp["db-connection"] = "OK"

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Errorf("error database %v", err)
		defer db.Close()
		return response.Success(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("migration error database %v", err),
			Data:    nil,
		})
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", "../../migration_files"),
		"mysql",
		driver,
	)
	if err != nil {
		log.Errorf("error migration %v", err)
		defer db.Close()
		// os.Exit(1)
		return response.Success(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error migration %v", err),
			Data:    nil,
		})
	}

	m.Steps(2)

	src, dbe := m.Close()
	if dbe != nil || src != nil {
		log.Errorf("error migration %v %v", src, dbe)
		defer db.Close()
		// os.Exit(1)
		return response.Success(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error migration %v", err),
			Data:    nil,
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
	})
}
