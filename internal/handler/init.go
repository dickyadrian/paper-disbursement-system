package server

import "github.com/dickyadrian/paper-disbursement-system/application"

type Handler struct {
	*application.App
}

func NewHandler(
	app *application.App,
) *Handler {
	return &Handler{app}
}
