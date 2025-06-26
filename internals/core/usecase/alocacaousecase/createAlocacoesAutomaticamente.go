package alocacaousecase

import (
	"math/rand"
	"fmt"

	"github.com/cristiantebaldi/class-organize-api/internals/core/domain"
	"github.com/cristiantebaldi/class-organize-api/internals/core/dto"
)

func (usecase usecase) CreateAlocacoesAutomaticamente(alocacaoRequest *dto.AlocacaoRequestBody) (*[]domain.Alocacao, error) {
	salas, err := usecase.salaRepository.Fetch()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar salas: %w", err)
	}

	salasAlocadas, err := usecase.repository.GetSalasAlocadas(alocacaoRequest.DiaSemana, alocacaoRequest.HorarioInicio, alocacaoRequest.HorarioFim)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar salas alocadas: %w", err)
	}

	salasAlocadasIDs := make(map[int]bool)
	for _, sala := range *salasAlocadas {
		salasAlocadasIDs[sala.ID] = true
	}

	var salasDisponiveis []domain.Sala
	for _, sala := range *salas {
		if !salasAlocadasIDs[sala.ID] {
			salasDisponiveis = append(salasDisponiveis, sala)
		}
	}

	turmas, err := usecase.turmaRepository.Fetch()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar turmas: %w", err)
	}

	turmasAlocadas, err := usecase.repository.GetTurmasAlocadas(alocacaoRequest.DiaSemana, alocacaoRequest.HorarioInicio, alocacaoRequest.HorarioFim)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar turmas alocadas: %w", err)
	}

	turmasAlocadasIDs := make(map[int]bool)
	for _, turma := range *turmasAlocadas {
		turmasAlocadasIDs[turma.ID] = true
	}

	var turmasDisponiveis []domain.Turma
	for _, turma := range *turmas {
		if !turmasAlocadasIDs[turma.ID] {
			turmasDisponiveis = append(turmasDisponiveis, turma)
		}
	}

	professores, err := usecase.professorRepository.Fetch()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar professores: %w", err)
	}

	professoresAlocados, err := usecase.repository.GetProfessoresAlocados(alocacaoRequest.DiaSemana, alocacaoRequest.HorarioInicio, alocacaoRequest.HorarioFim)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar professores alocados: %w", err)
	}


	professoresAlocadosIDs := make(map[int]bool)
	for _, professor := range *professoresAlocados {
		professoresAlocadosIDs[professor.ID] = true
	}

	var professoresDisponiveis []domain.Professor
	for _, professor := range *professores {
		if !professoresAlocadosIDs[professor.ID] {
			professoresDisponiveis = append(professoresDisponiveis, professor)
		}
	}

	if len(professoresDisponiveis) == 0 || len(salasDisponiveis) == 0 || len(turmasDisponiveis) == 0 {
		return nil, fmt.Errorf("não há recursos suficientes para fazer alocações")
	}

	var alocacoesCriadas []domain.Alocacao

	numAlocacoes := min(len(professoresDisponiveis), min(len(salasDisponiveis), len(turmasDisponiveis)))

	rand.Shuffle(len(professoresDisponiveis), func(i, j int) {
		professoresDisponiveis[i], professoresDisponiveis[j] = professoresDisponiveis[j], professoresDisponiveis[i]
	})

	rand.Shuffle(len(salasDisponiveis), func(i, j int) {
		salasDisponiveis[i], salasDisponiveis[j] = salasDisponiveis[j], salasDisponiveis[i]
	})

	rand.Shuffle(len(turmasDisponiveis), func(i, j int) {
		turmasDisponiveis[i], turmasDisponiveis[j] = turmasDisponiveis[j], turmasDisponiveis[i]
	})

	for i := 0; i < numAlocacoes; i++ {
		novaAlocacao := &dto.AlocacaoRequestBody{
			ProfessorID:   professoresDisponiveis[i].ID,
			SalaID:        salasDisponiveis[i].ID,
			TurmaID:       turmasDisponiveis[i].ID,
			DiaSemana:     alocacaoRequest.DiaSemana,
			HorarioInicio: alocacaoRequest.HorarioInicio,
			HorarioFim:    alocacaoRequest.HorarioFim,
		}

		alocacaoCriada, err := usecase.repository.Create(novaAlocacao)
		if err != nil {
			continue
		}

		alocacoesCriadas = append(alocacoesCriadas, *alocacaoCriada)
	}

	if len(alocacoesCriadas) == 0 {
		return nil, fmt.Errorf("não foi possível criar nenhuma alocação")
	}

	return &alocacoesCriadas, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
