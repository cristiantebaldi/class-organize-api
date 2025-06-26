package turmarepository

import (
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
)

func (repository repository) Fetch() (*[]domain.Turma, error) {
	turmas := []domain.Turma{}

	query := `SELECT id, nome, curso, periodo, quant_alunos FROM turmas`
	{
		err := repository.db.Select(&turmas, query)

		if err != nil {
			return nil, fmt.Errorf("%s", "ERROR:" + err.Error())
		}
	}

	return &turmas, nil
}
