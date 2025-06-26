# Class Organize - Sistema de Gerenciamento de Salas, Turmas e Professores

Este é um sistema de back-end desenvolvido em Go para gerenciar a alocação de salas, turmas e professores em uma instituição de ensino.

## Funcionalidades

- Gerenciamento de professores (CRUD)
- Gerenciamento de salas (CRUD)
- Gerenciamento de turmas (CRUD)
- Alocação de professores, salas e turmas (CRUD)
- Verificação de disponibilidade de salas
- Consulta de alocações por professor, sala ou turma
- **Alocação automática de professore, sala e turma**

## Requisitos

- Go 1.16 ou superior
- PostgreSQL 12 ou superior

## Configuração

1. Clone o repositório
2. Configure o arquivo `config.json` com as credenciais do banco de dados:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=class_organize
```

3. Instale as dependências:

```bash
go mod tidy
```

4. Execute a aplicação:

```bash
go run main.go
```

O servidor será iniciado na porta 3000.

## API Endpoints

### Professores

- `GET /professor` - Listar todos os professores
- `GET /professor/{id}` - Obter um professor específico
- `POST /professor` - Criar um novo professor
- `PUT /professor/{id}` - Atualizar um professor
- `DELETE /professor/{id}` - Remover um professor

### Salas

- `GET /sala` - Listar todas as salas
- `GET /sala/{id}` - Obter uma sala específica
- `POST /sala` - Criar uma nova sala
- `PUT /sala/{id}` - Atualizar uma sala
- `DELETE /sala/{id}` - Remover uma sala

### Turmas

- `GET /turma` - Listar todas as turmas
- `GET /turma/{id}` - Obter uma turma específica
- `POST /turma` - Criar uma nova turma
- `PUT /turma/{id}` - Atualizar uma turma
- `DELETE /turma/{id}` - Remover uma turma

### Alocações

- `GET /alocacao` - Listar todas as alocações
- `GET /alocacao{id}` - Obter uma alocação específica
- `POST /alocacao` - Criar uma nova alocação
- `PUT /alocacao{id}` - Atualizar uma alocação
- `DELETE /alocacao{id}` - Remover uma alocação

### Consultas Especiais

- `GET /alocacao/sala/{id}` - Listar alocações por sala
- `GET /alocacao/professor/{id}` - Listar alocações por professor
- `GET /alocacao/turma/{id}` - Listar alocações por turma

## Exemplos de Uso

### Criar um Professor

```bash
curl -X POST http://localhost:3000/professor \
  -H "Content-Type: application/json" \
  -d '{"nome":"João Silva","email":"joao@exemplo.com","formacao":"Mestrado em Computação","disciplina":"Programação Web"}'
```

### Criar uma Sala

```bash
curl -X POST http://localhost:3000/sala \
  -H "Content-Type: application/json" \
  -d '{"numero":"101","capacidade":40,"bloco":"A","tipo":"Laboratório"}'
```

### Criar uma Turma

```bash
curl -X POST http://localhost:3000/turma \
  -H "Content-Type: application/json" \
  -d '{"nome":"Turma A","curso":"Sistemas de Informação","periodo":"Noturno","quant_alunos":35}'
```

### Criar uma Alocação

```bash
curl -X POST http://localhost:3000/alocacao \
  -H "Content-Type: application/json" \
  -d '{"professor_id":1,"sala_id":1,"turma_id":1,"dia_semana":"Segunda","horario_inicio":"19:00","horario_fim":"22:30"}'
```

### Criar alocação automática
```bash
curl -X POST http://localhost:3000/alocacao/automatico \
  -H "Content-Type: application/json" \
  -d '{
  "dia_semana":"Segunda","horario_inicio":"19:00",
  "horario_fim":"22:30"
  }'
```

## Licença

Este projeto está licenciado sob a licença MIT.