## Projeto: Aplicação de Linha de Comando (CLI)

Este projeto consiste em uma ferramenta de terminal para buscar informações de rede (IPs e nomes de servidores) utilizando a biblioteca `urfave/cli`.

### Conceitos Aplicados
- **Gerenciamento de Módulos**: Uso de `go mod init` e `go get` para gerenciar dependências externas.
- **Estrutura de Pacotes**: Divisão entre o pacote principal (`main`) e pacotes de lógica interna (`app`).
- **Ponteiros de Funções**: A função `Gerar()` retorna um ponteiro `*cli.App`, garantindo que estamos manipulando a mesma instância da aplicação em todo o fluxo.

### Bibliotecas Utilizadas
- `github.com/urfave/cli`: Framework para construção de interfaces de linha de comando em Go.

### Como Executar
```bash
go run main.go ip --host google.com