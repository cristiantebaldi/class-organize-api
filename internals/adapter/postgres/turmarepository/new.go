package turmarepository

import (
	"github.com/cristiantebaldi/class-organize-api/internals/adapter/postgres"
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

type repository struct {
	db postgres.DatabaseConfig
}

// New returns contract implementation of AlocacaoRepository
func New(db postgres.DatabaseConfig) domain.TurmaRepository {
	return &repository{
		db: db,
	}
}