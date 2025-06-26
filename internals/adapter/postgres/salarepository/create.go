package salarepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (repository repository) Create(
	salaRequest *dto.SalaRequestBody,
) (*domain.Sala, error) {
	sala := domain.Sala{}

	err := repository.db.QueryRowx(`
		INSERT INTO salas (numero, capacidade, bloco, tipo) 
		VALUES ($1, $2, $3, $4) returning *;
	`,
		salaRequest.Numero,
		salaRequest.Capacidade,
		salaRequest.Bloco,
		salaRequest.Tipo,
	).StructScan(&sala)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &sala, nil
}