package salarepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) Fetch() (*[]domain.Sala, error) {
	salas := []domain.Sala{}

	query := `SELECT id, numero, capacidade, bloco, tipo FROM salas`
	{
		err := repository.db.Select(&salas, query)

		if err != nil {
			return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
		}
	}

	return &salas, nil
}
