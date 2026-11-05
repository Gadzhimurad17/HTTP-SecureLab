package errors

import "github.com/Gagzhimurad17/NullPoint/web/backend/content/internal/domain/models"

type AppErrors struct {
	Err     error
	Message string
	Kind    models.ErrorKind
}

func (err *AppErrors) Error() string {
	if err.Err != nil {
		return err.Message + ": " + err.Error()
	}

	return err.Message
}
