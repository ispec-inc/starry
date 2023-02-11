package presenter

import "github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/apperror"

type errorResponse struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func newErrorResponse(code apperror.Code, detail string) errorResponse {
	return errorResponse{
		Code:   string(code),
		Detail: detail,
	}
}
