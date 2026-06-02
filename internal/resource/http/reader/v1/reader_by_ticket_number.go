package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/web-rabis/httperrors"
)

func (res *ReaderResource) readerByTicketNumber(w http.ResponseWriter, r *http.Request) {
	ticketNumberParam := chi.URLParam(r, "ticketNumber")

	ticketNumber, err := strconv.ParseInt(ticketNumberParam, 10, 64)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	reader, err := res.readerMan.ReaderByTicketNumber(r.Context(), ticketNumber)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	render.JSON(w, r, reader)
}
