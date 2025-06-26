package di

import (
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/http/professorservice"
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres"
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres/professorrepository"
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/usecase/professorusecase"
	"github.com/go-playground/validator/v10"
)

func ConfigProfessorDI(conn postgres.DatabaseConfig, validate *validator.Validate) domain.ProfessorHTTPService {
	professorRepository := professorrepository.New(conn)
	professorUseCase := professorusecase.New(professorRepository)
	professorService := professorservice.New(validate, professorUseCase)

	return professorService
}