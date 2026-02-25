package server

import (
	"net/http"

	// "github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"htmx-templ-sqlc/cmd/web"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))
	e.GET("/home", echo.WrapHandler(http.HandlerFunc(HomeHandler)))
	e.GET("/media", echo.WrapHandler(http.HandlerFunc(MediaHandler)))
	e.GET("/people", echo.WrapHandler(http.HandlerFunc(PeopleHandler)))
	e.POST("/people", echo.WrapHandler(http.HandlerFunc(PeopleTabHandler)))
	e.GET("/summer_program", echo.WrapHandler(http.HandlerFunc(SummerProgramHandler)))
	// e.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
	// e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))
	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
