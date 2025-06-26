package di

import (
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/http/salaservice"
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres"
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres/salarepository"
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/usecase/salausecase"
	"github.com/go-playground/validator/v10"
)

func ConfigSalaDI(conn postgres.DatabaseConfig, validate *validator.Validate) domain.SalaHTTPService {
	salaRepository := salarepository.New(conn)
	salaUseCase := salausecase.New(salaRepository)
	salaService := salaservice.New(validate, salaUseCase)

	return salaService
}