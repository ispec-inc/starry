package presenter

import (
	"net/http"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/apperror"
)

var code2status = map[apperror.Code]int{
	apperror.CodeError:        http.StatusInternalServerError,
	apperror.CodeInvalid:      http.StatusBadRequest,
	apperror.CodeUnauthorized: http.StatusUnauthorized,
	apperror.CodeNotFound:     http.StatusNotFound,
}
