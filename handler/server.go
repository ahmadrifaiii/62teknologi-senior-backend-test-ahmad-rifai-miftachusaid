package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"62tech.co/service/config/env"
	"62tech.co/service/utl/middleware/logging"
	"62tech.co/service/utl/middleware/secure"

	"io/ioutil"
	"net/http"

	"os"
	"os/signal"
	"time"

	"62tech.co/service/pkg/swagger"
	_ "62tech.co/service/pkg/swagger/docs"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	mysqlMigrate "62tech.co/service/utl/database/migrate/mysql"
)

const DefaultPort = 8080

// HTTPServerMain main function for serving services over http

func (s *Service) HTTPServerMain() *echo.Echo {
	// to active swagger
	swagger.Init()

	e := echo.New()

	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())
	e.Use(secure.Headers())
	e.Use(logging.Logging())

	// administrator group
	adm := e.Group("/api/v1")

	// domain module business
	BusinessModule := adm.Group("/business")
	s.BusinessModule.HandleRest(BusinessModule)

	// domain module system
	SystemModule := adm.Group("/system")
	s.SystemModule.HandleRest(SystemModule)

	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.GET("/migration", mysqlMigrate.Migration)

	data, _ := json.MarshalIndent(e.Routes(), "", "  ")

	ioutil.WriteFile("routes.json", data, 0644)

	return e
}

func (s *Service) StartServer() {
	server := s.HTTPServerMain()
	listenerPort := fmt.Sprintf(":%v", env.Conf.HttpPort)
	if err := server.StartServer(&http.Server{
		Addr:         listenerPort,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}); err != nil {
		server.Logger.Fatal(err.Error())
	}
}

func (s *Service) ShutdownServer() {
	server := s.HTTPServerMain()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err.Error())
	}
}
