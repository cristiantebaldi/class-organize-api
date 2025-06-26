package dto

import (
	"encoding/json"
	"io"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gocarina/gocsv"
)

type ProfessorRequestBody struct {
	Nome       string `json:"nome"`
	Email      string `json:"email"`
	Formacao   string `json:"formacao"`
	Disciplina string `json:"disciplina"`
}

func FromJsonProfessorRequestBody(
	body io.ReadCloser,
	validate *validator.Validate,
) (*ProfessorRequestBody, error) {
	var professor ProfessorRequestBody

	if err := json.NewDecoder(body).Decode(&professor); err != nil {
		return nil, err
	}
	if err := validate.Struct(professor); err != nil {
		return nil, err
	}

	return &professor, nil
}

func FromCSVProfessorSlice(
	file *os.File,
) (*[]ProfessorRequestBody, error) {
	professores := []ProfessorRequestBody{}
	if err := gocsv.UnmarshalFile(file, &professores); err != nil {
		return nil, err
	}

	return &professores, nil
}
