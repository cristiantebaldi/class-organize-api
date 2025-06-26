package salarepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) Delete(id int32) (*domain.Sala, error) {
	sala := domain.Sala{}

	err := repository.db.
		QueryRowx(`DELETE FROM salas WHERE id = $1 returning *;`, id).
		StructScan(&sala)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &sala, nil
}
