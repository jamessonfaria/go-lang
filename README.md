# go-lang

Guia rapido com comandos essenciais para trabalhar com Go no dia a dia.

## Comandos Basicos

### Ver versao do Go

```bash
go version
```

Mostra a versao instalada do Go.

### Ver ambiente do Go

```bash
go env
```

Exibe as variaveis de ambiente usadas pela ferramenta.

### Abrir ajuda geral

```bash
go help
```

Mostra a ajuda principal da CLI do Go.

## Gerenciamento de Modulos

### Iniciar um projeto

```bash
go mod init nome-do-modulo
```

Cria o arquivo `go.mod` e inicia um novo modulo.

### Organizar dependencias

```bash
go mod tidy
```

Limpa dependencias nao utilizadas e baixa as necessarias.

### Baixar dependencias

```bash
go mod download
```

Baixa as dependencias sem compilar o projeto.

### Usar pasta vendor

```bash
go mod vendor
```

Copia as dependencias para o diretorio `vendor`.

## Executar e Compilar

### Executar sem gerar binario

```bash
go run main.go
```

Executa o codigo diretamente.

### Compilar o projeto

```bash
go build
```

Gera o executavel padrao do projeto.

### Compilar com nome personalizado

```bash
go build -o app
```

Define o nome do binario gerado.

### Instalar binario

```bash
go install
```

Instala o executavel no diretorio de binarios do Go.

## Testes

### Rodar todos os testes

```bash
go test ./...
```

Executa os testes de todos os pacotes do projeto.

### Rodar testes em modo verboso

```bash
go test -v
```

Exibe mais detalhes durante a execucao.

### Ver cobertura de testes

```bash
go test -cover
```

Mostra a cobertura de testes.

## Gerenciar Pacotes

### Adicionar ou atualizar dependencia

```bash
go get pacote
```

Adiciona ou atualiza uma dependencia.

### Listar pacotes do projeto

```bash
go list ./...
```

Lista todos os pacotes encontrados no projeto.

## Formatacao e Qualidade

### Formatar codigo

```bash
go fmt ./...
```

Formata automaticamente o codigo Go.

### Verificar problemas comuns

```bash
go vet ./...
```

Ajuda a encontrar erros e padroes suspeitos.

## Documentacao

### Ver documentacao de funcao

```bash
go doc fmt.Println
```

Mostra a documentacao de uma funcao especifica.

### Ver documentacao de pacote

```bash
go doc pacote
```

Mostra a documentacao de um pacote.

## Outros Comandos Uteis

### Limpar arquivos gerados

```bash
go clean
```

Remove arquivos criados durante build e outras operacoes.

### Executar geracao automatica

```bash
go generate
```

Executa diretivas `//go:generate` do projeto.

## Fluxo Tipico de um Projeto Novo

```bash
go mod init meu-projeto
go run main.go
go build
go test ./...
go mod tidy
```

Sequencia comum para iniciar, executar, compilar, testar e organizar dependencias.
