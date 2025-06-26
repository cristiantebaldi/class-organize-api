package turmarepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) GetByID(id int32) (*domain.Turma, error) {
	turma := domain.Turma{}

	err := repository.db.Get(&turma, `SELECT * FROM turmas WHERE id = $1;`, id)

	if err != nil {
		return nil, fmt.Errorf("%s", "ERROR:"+err.Error())
	}

	return &turma, nil
}
