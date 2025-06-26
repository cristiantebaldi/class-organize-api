package professorservice

import (
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/go-playground/validator/v10"
)

type service struct {
	validator *validator.Validate
	usecase   domain.ProfessorUsecase
}

func New(validator *validator.Validate, usecase domain.ProfessorUsecase) domain.ProfessorHTTPService {
	return &service{
		validator: validator,
		usecase:   usecase,
	}
}
