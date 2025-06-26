package salarepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) GetByID(id int32) (*domain.Sala, error) {
	sala := domain.Sala{}

	err := repository.db.Get(&sala, `SELECT * FROM salas WHERE id = $1;`, id)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:"+err.Error())
	}

	return &sala, nil
}
