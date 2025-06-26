CREATE TABLE IF NOT EXISTS professores (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    formacao VARCHAR(100),
    disciplina VARCHAR(100) NOT NULL
);


CREATE TABLE IF NOT EXISTS salas (
    id SERIAL PRIMARY KEY,
    numero VARCHAR(20) NOT NULL,
    capacidade INT NOT NULL,
    bloco VARCHAR(50),
    tipo VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS turmas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    curso VARCHAR(100) NOT NULL,
    periodo VARCHAR(50),
    quant_alunos INT
);

CREATE TABLE IF NOT EXISTS alocacoes (
    id SERIAL PRIMARY KEY,
    professor_id INT REFERENCES professores(id),
    sala_id INT REFERENCES salas(id),
    turma_id INT REFERENCES turmas(id),
    dia_semana VARCHAR(20) NOT NULL,
    horario_inicio VARCHAR(10) NOT NULL,
    horario_fim VARCHAR(10) NOT NULL,
    CONSTRAINT unique_alocacao UNIQUE(sala_id, dia_semana, horario_inicio)
);