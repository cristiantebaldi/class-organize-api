package professorrepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) Delete(id int32) (*domain.Professor, error) {
	professor := domain.Professor{}

	err := repository.db.
		QueryRowx(`DELETE FROM professores WHERE id = $1 returning *;`, id).
		StructScan(&professor)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &professor, nil
}
