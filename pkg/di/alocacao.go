package di

import (
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/http/alocacaoservice"
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres"
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres/alocacaorepository"
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres/professorrepository"
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres/salarepository"
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres/turmarepository"
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/usecase/alocacaousecase"
	"github.com/go-playground/validator/v10"
)

func ConfigAlocacaoDI(conn postgres.DatabaseConfig, validate *validator.Validate) domain.AlocacaoHTTPService {
	alocacaoRepository := alocacaorepository.New(conn)
	professorRepository := professorrepository.New(conn)
	salaRepository := salarepository.New(conn)
	turmaRepository := turmarepository.New(conn)

	alocacaoUseCase := alocacaousecase.New(alocacaoRepository, professorRepository, salaRepository, turmaRepository)
	alocacaoService := alocacaoservice.New(validate, alocacaoUseCase)

	return alocacaoService
}