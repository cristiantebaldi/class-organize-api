package alocacaorepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) Delete(id int32) (*domain.Alocacao, error) {
	alocacao := domain.Alocacao{}

	err := repository.db.
		QueryRowx(`DELETE FROM alocacoes WHERE id = $1 returning *;`, id).
		StructScan(&alocacao)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &alocacao, nil
}
