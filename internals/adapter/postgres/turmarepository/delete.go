package turmarepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) Delete(id int32) (*domain.Turma, error) {
	turma := domain.Turma{}

	err := repository.db.
		QueryRowx(`DELETE FROM turmas WHERE id = $1 returning *;`, id).
		StructScan(&turma)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
	}

	return &turma, nil
}
