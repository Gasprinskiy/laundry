package global

import (
	"errors"
	"net/http"
)

var (
	ErrNoData        = errors.New("данные не найдены")
	ErrInvalidParam  = errors.New("не верные параметры")
	ErrInternalError = errors.New("внутреняя ошибка сервера")
)

var ErrStatusCodes = map[error]int{
	ErrNoData:        http.StatusNotFound,
	ErrInternalError: http.StatusInternalServerError,
	ErrInvalidParam:  http.StatusBadRequest,
}
