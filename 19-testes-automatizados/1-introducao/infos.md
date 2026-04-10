## Testes Automatizados em Go

Testes automatizados são fundamentais para garantir que alterações no código não quebrem funcionalidades existentes (testes de regressão).

### Regras do Pacote `testing`
- **Nomeação**: Arquivos devem terminar em `_test.go`.
- **Assinatura**: Funções de teste devem começar com a palavra `Test` e receber o parâmetro `*testing.T`.
- **Comando**: Utilizamos `go test` no terminal para executar as verificações.

### Por que testar?
1. **Confiança**: Permite refatorar o código sabendo que a lógica principal continua válida.
2. **Documentação**: Testes bem escritos servem como exemplos de como a função deve se comportar.
3. **Segurança**: Ajuda a prever cenários de erro e entradas inválidas antes que cheguem ao usuário.