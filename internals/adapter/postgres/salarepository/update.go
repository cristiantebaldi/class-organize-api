package salarepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (repository repository) Update(
	id int32,
	salaRequest *dto.SalaRequestBody,
) (*domain.Sala, error) {
	sala := domain.Sala{}

	err := repository.db.QueryRowx(`
		UPDATE salas SET 
		numero = $1, 
		capacidade = $2, 
		bloco = $3, 
		tipo = $4 
		WHERE id = $5 
		returning *;
	`,
		salaRequest.Numero,
		salaRequest.Capacidade,
		salaRequest.Bloco,
		salaRequest.Tipo,
		id,
	).StructScan(&sala)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &sala, nil
}
