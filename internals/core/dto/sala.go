package dto

import (
	"encoding/json"
	"io"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gocarina/gocsv"
)

type SalaRequestBody struct {
	Numero     string `json:"numero"`
	Capacidade int    `json:"capacidade"`
	Bloco      string `json:"bloco"`
	Tipo       string `json:"tipo"`
}

func FromJsonSalaRequestBody(
	body io.ReadCloser,
	validate *validator.Validate,
) (*SalaRequestBody, error) {
	var sala SalaRequestBody

	if err := json.NewDecoder(body).Decode(&sala); err != nil {
		return nil, err
	}
	if err := validate.Struct(sala); err != nil {
		return nil, err
	}

	return &sala, nil
}

func FromCSVSalaSlice(
	file *os.File,
) (*[]SalaRequestBody, error) {
	salas := []SalaRequestBody{}
	if err := gocsv.UnmarshalFile(file, &salas); err != nil {
		return nil, err
	}

	return &salas, nil
}
