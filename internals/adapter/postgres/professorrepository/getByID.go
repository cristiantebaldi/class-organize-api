package professorrepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) GetByID(id int32) (*domain.Professor, error) {
	professor := domain.Professor{}

	err := repository.db.Get(&professor, `SELECT * FROM professores WHERE id = $1;`, id)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &professor, nil
}