package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dickyadrian/paper-disbursement-system/application"
	server "github.com/dickyadrian/paper-disbursement-system/internal/handler"
	"github.com/labstack/echo/v4"
)

func RunServer(ctx context.Context, app *application.App) {
	port := app.AppConfig.Port

	e := echo.New()
	h := server.NewHandler(app)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	v1 := e.Group("/v1")

	v1.POST("/users", h.CreateUser)
	v1.GET("/users/:id", h.GetUser)
	v1.POST("/disbursements", h.CreateDisbursement)

	done := make(chan bool)
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal(err)
		}

		done <- true
	}()

	<-ctx.Done()
	if err := e.Shutdown(ctx); err != nil {
		fmt.Printf("Error when shutting down server")
	}

	<-done
	fmt.Printf("Server shut down")
}
