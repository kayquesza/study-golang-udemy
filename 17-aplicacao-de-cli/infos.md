## Projeto: Aplicação de Linha de Comando (CLI)

Este projeto consiste numa ferramenta de terminal para buscar informações de rede utilizando a biblioteca `urfave/cli`.

### Conceitos Aplicados
- **Gerenciamento de Módulos**: Uso de `go mod init` e `go get` para gerenciar dependências externas.
- **Refatoração e DRY**: Centralização de configurações de `Flags` para reuso em múltiplos comandos.
- **Interação com Rede (net package)**:
    - `LookupIP`: Busca os endereços IP vinculados a um host.
    - `LookupNS`: Identifica os Servidores de Nome (Name Servers) responsáveis pelo domínio.
- **Estrutura de Pacotes**: Manutenção da separação entre a inicialização (`main`) e a lógica de comandos (`app`).

### Bibliotecas Utilizadas
- `github.com/urfave/cli`: Framework profissional para construção de CLIs em Go.

### Como Executar
```bash
# Para buscar IPs
go run main.go ip --host google.com

# Para buscar Servidores de Nome
go run main.go servidores --host google.com