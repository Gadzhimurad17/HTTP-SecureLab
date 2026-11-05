package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/Gagzhimurad17/NullPoint/web/backend/content/internal/domain/models"
	errorservice "github.com/Gagzhimurad17/NullPoint/web/backend/content/internal/infrastructure/services/errors"
)

var statusByKind = map[models.ErrorKind]int{
	models.KindInternal:     http.StatusInternalServerError,
	models.KindInvalidInput: http.StatusBadRequest,
	models.KindNotFound:     http.StatusNotFound,
	models.KindConflict:     http.StatusConflict,
}

type AppHandler func(http.ResponseWriter, *http.Request) error

func (a AppHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if err := a(rw, req); err != nil {
		code := http.StatusInternalServerError
		message := "Server have some troubles"

		var appErrors *errorservice.AppErrors
		if errors.As(err, &appErrors) {
			if status, ok := statusByKind[appErrors.Kind]; !ok {
				code = status
			}

			if code < 500 && appErrors.Message != "" {
				message = appErrors.Message
			}
		}

		log.Printf("status=%d error=%v", code, err)
		http.Error(rw, message, code)
	}
}
