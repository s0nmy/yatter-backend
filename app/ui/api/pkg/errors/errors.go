package errors

import (
	"errors"
	"log/slog"
	"net/http"
	y_errors "yatter-backend-go/pkg/errors"
	"yatter-backend-go/pkg/errors/code"
)

var codes = map[code.StatusCode]int{
	code.BadRequest:   http.StatusBadRequest,
	code.Unauthorized: http.StatusUnauthorized,
	code.Forbidden:    http.StatusForbidden,
	code.NotFound:     http.StatusNotFound,
	code.Conflict:     http.StatusConflict,
	code.Internal:     http.StatusInternalServerError,
}

func Handle(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	if v, ok := errors.AsType[*y_errors.Status](err); ok {
		slog.Info(v.UIMessage(), "code", v.Code())
		http.Error(w, v.UIMessage(), codes[v.Code()])
	} else {
		slog.Info(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
