package alocacaorepository

import (
	"fmt"
	
	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) GetByID(id int32) (*domain.Alocacao, error) {
	alocacao := domain.Alocacao{}

	err := repository.db.Get(&alocacao, `SELECT * FROM alocacoes WHERE id = $1;`, id)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &alocacao, nil
}