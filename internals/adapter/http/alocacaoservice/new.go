package alocacaoservice

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/go-playground/validator/v10"
)

type service struct {
	validator *validator.Validate
	usecase   domain.AlocacaoUsecase
}

func New(validator *validator.Validate, usecase domain.AlocacaoUsecase) domain.AlocacaoHTTPService {
	return &service{
		validator: validator,
		usecase:   usecase,
	}
}
