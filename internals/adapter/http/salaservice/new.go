package salaservice

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/go-playground/validator/v10"
)

type service struct {
	validator *validator.Validate
	usecase   domain.SalaUsecase
}

func New(validator *validator.Validate, usecase domain.SalaUsecase) domain.SalaHTTPService {
	return &service{
		validator: validator,
		usecase:   usecase,
	}
}
