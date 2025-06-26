package professorrepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) Fetch() (*[]domain.Professor, error) {
	professores := []domain.Professor{}

	query := `SELECT id, nome, email, formacao, disciplina FROM professores`
	{
		err := repository.db.Select(&professores, query)

		if err != nil {
			return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
		}
	}

	return &professores, nil
}
