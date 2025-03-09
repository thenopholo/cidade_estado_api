# Cidade Estado API

Uma API em Go para busca de estados brasileiros utilizando o Gin framework.

## 📋 Sobre o Projeto

Este projeto é uma API REST que fornece informações sobre estados brasileiros. A API consome dados da [Brasil API](https://brasilapi.com.br/), processando e disponibilizando essas informações através de endpoints RESTful.

## 🏗️ Arquitetura

O projeto segue uma arquitetura limpa (Clean Architecture), organizando o código em camadas com responsabilidades bem definidas:

```
cidade_estado_api/
├── cmd/                        # Ponto de entrada da aplicação
│   └── main.go                 # Inicializa o servidor
├── internal/                   # Código que não deve ser importado por outros projetos
│   ├── app/                    # Lógica da aplicação
│   │   ├── bootstrap/          # Inicialização e configuração do servidor
│   │   │   └── server.go
│   │   └── handlers/           # Manipuladores HTTP
│   │       └── locations/
│   │           ├── dto/        # Objetos de Transferência de Dados
│   │           │   └── state_response.go
│   │           └── location_handler.go
│   ├── domain/                 # Regras de negócio e entidades
│   │   └── entities/
│   │       └── stete_entity.go
│   └── infrastructure/         # Implementações externas
│       └── repositories/
│           └── location/
│               ├── dto/
│               │   └── brazil_api_response.go
│               └── location_repository.go
```

## 🧩 Conceitos Implementados

### 1. Clean Architecture (Arquitetura Limpa)

O projeto separa claramente as responsabilidades em camadas:

- **Domain (Domínio)**: Contém as entidades de negócio (StateEntity)
- **Application (Aplicação)**: Contém os handlers e DTOs para API
- **Infrastructure (Infraestrutura)**: Contém repositórios e integrações externas

### 2. Dependency Injection (Injeção de Dependência)

Os componentes são acoplados através de suas interfaces e injetados onde necessário:

```go
// Injetando o repositório no handler
locationRepository := repositories.NewLocationRepository()
locationHandler := locations.NewLocationhanedler(locationRepository)
```

### 3. DTO (Data Transfer Objects)

DTOs são usados para transferir dados entre camadas sem expor detalhes de implementação:

```go
// DTO para resposta da API
type StateResponse struct {
    Acronym string `json:"acronym,omitempty"`
    Nome    string `json:"nome,omitempty"`
}

// DTO para resposta da Brasil API
type BrazilApiResponse struct {
    Sigla string `json:"sigla,omitempty"`
    Nome  string `json:"nome,omitempty"`
}
```

### 4. Repository Pattern (Padrão Repositório)

Abstrai o acesso a fontes de dados externas:

```go
func (l *LocationRepository) GetStates() ([]entities.StateEntity, error) {
    // Implementação para buscar dados da Brasil API
}
```

### 5. REST API com Gin

Utiliza o framework Gin para criar endpoints RESTful:

```go
func configureRoutes(e *gin.Engine) {
    // ...
    g := e.Group("/api/v1")
    {
        g.GET("/states", locationHandler.GetAllStates)
    }
}
```

### 6. Tratamento de Contexto e Timeout

Implementação de timeouts para chamadas HTTP:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req = req.WithContext(ctx)
```

## 🚀 Como Executar o Projeto

### Pré-requisitos

- Go 1.23 ou superior
- Git

### Passos para Execução

1. Clone o repositório:
   ```bash
   git clone https://github.com/thenopholo/cidade_estado_api.git
   cd cidade_estado_api
   ```

2. Baixe as dependências:
   ```bash
   go mod download
   ```

3. Execute a aplicação:
   ```bash
   go run cmd/main.go
   ```

4. O servidor estará disponível em `http://localhost:8080`

## 📡 Utilizando a API

### Listar Todos os Estados

**Endpoint**: `GET /api/v1/states`

**Exemplo de Requisição**:
```bash
curl -X GET http://localhost:8080/api/v1/states
```

**Exemplo de Resposta**:
```json
[
  {
    "acronym": "AC",
    "nome": "Acre"
  },
  {
    "acronym": "AL",
    "nome": "Alagoas"
  },
  // ... outros estados
]
```

## 🧪 Testando a API

Você pode testar a API usando ferramentas como:

- [curl](https://curl.se/) (linha de comando)
- [Postman](https://www.postman.com/) (interface gráfica)
- [Insomnia](https://insomnia.rest/) (interface gráfica)

Exemplo com Postman:
1. Abra o Postman
2. Crie uma nova requisição GET
3. Digite a URL: `http://localhost:8080/api/v1/states`
4. Clique em "Send"
5. Visualize a resposta JSON com a lista de estados brasileiros

## 🔄 Fluxo de Dados

1. O cliente faz uma requisição HTTP GET para `/api/v1/states`
2. O Gin roteia a requisição para o `GetAllStates` do `LocationHandler`
3. O handler chama o método `GetStates` do `LocationRepository`
4. O repositório faz uma chamada HTTP para a Brasil API
5. Os dados são decodificados de JSON para o DTO `BrazilApiResponse`
6. Os dados são convertidos para entidades de domínio `StateEntity`
7. O handler converte as entidades para DTOs `StateResponse`
8. A resposta é serializada em JSON e retornada ao cliente

## 📝 Próximos Passos (Possíveis Melhorias)

- Implementar busca de cidades por estado
- Adicionar cache para reduzir chamadas à Brasil API
- Implementar testes unitários e de integração
- Adicionar documentação com Swagger
- Implementar autenticação e autorização
- Adicionar logs estruturados
- Implementar healthcheck
- Containerizar a aplicação com Docker