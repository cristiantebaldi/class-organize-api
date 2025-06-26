package dto

import (
	"encoding/json"
	"io"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gocarina/gocsv"
)

type AlocacaoRequestBody struct {
	ProfessorID   int       `json:"professor_id"`
	SalaID        int       `json:"sala_id"`
	TurmaID       int       `json:"turma_id"`
	DiaSemana     string    `json:"dia_semana"`
	HorarioInicio string    `json:"horario_inicio"`
	HorarioFim    string    `json:"horario_fim"`
}

func FromJsonAlocacaoRequestBody(
	body io.ReadCloser,
	validate *validator.Validate,
) (*AlocacaoRequestBody, error) {
	var alocacao AlocacaoRequestBody

	if err := json.NewDecoder(body).Decode(&alocacao); err != nil {
		return nil, err
	}
	if err := validate.Struct(alocacao); err != nil {
		return nil, err
	}

	return &alocacao, nil
}

func FromCSVAlocacaoSlice(
	file *os.File,
) (*[]AlocacaoRequestBody, error) {
	alocacoes := []AlocacaoRequestBody{}
	if err := gocsv.UnmarshalFile(file, &alocacoes); err != nil {
		return nil, err
	}

	return &alocacoes, nil
}
