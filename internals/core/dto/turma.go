package dto

import (
	"encoding/json"
	"io"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gocarina/gocsv"
)

type TurmaRequestBody struct {
	Nome        string `json:"nome"`
	Curso       string `json:"curso"`
	Periodo     string `json:"periodo"`
	QuantAlunos int    `json:"quant_alunos"`
}

func FromJsonTurmaRequestBody(
	body io.ReadCloser,
	validate *validator.Validate,
) (*TurmaRequestBody, error) {
	var turma TurmaRequestBody

	if err := json.NewDecoder(body).Decode(&turma); err != nil {
		return nil, err
	}
	if err := validate.Struct(turma); err != nil {
		return nil, err
	}

	return &turma, nil
}

func FromCSVTurmaSlice(
	file *os.File,
) (*[]TurmaRequestBody, error) {
	turmas := []TurmaRequestBody{}
	if err := gocsv.UnmarshalFile(file, &turmas); err != nil {
		return nil, err
	}

	return &turmas, nil
}
