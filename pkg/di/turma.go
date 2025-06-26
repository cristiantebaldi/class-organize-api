package di

import (
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/http/turmaservice"
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres"
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres/turmarepository"
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/usecase/turmausecase"
	"github.com/go-playground/validator/v10"
)

func ConfigTurmaDI(conn postgres.DatabaseConfig, validate *validator.Validate) domain.TurmaHTTPService {
	turmaRepository := turmarepository.New(conn)
	turmaUseCase := turmausecase.New(turmaRepository)
	turmaService := turmaservice.New(validate, turmaUseCase)

	return turmaService
}