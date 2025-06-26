package turmaservice

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/go-playground/validator/v10"
)

type service struct {
	validator *validator.Validate
	usecase   domain.TurmaUsecase
}

func New(validator *validator.Validate, usecase domain.TurmaUsecase) domain.TurmaHTTPService {
	return &service{
		validator: validator,
		usecase:   usecase,
	}
}
