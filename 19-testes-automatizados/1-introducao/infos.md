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

### Anatomia de um Teste Unitário

Para testar uma unidade (função ou método) de forma isolada, seguimos este padrão:

1.  **Cenário**: Definimos uma entrada (ex: "Rua Paulista") e o resultado esperado (ex: "Rua").
2.  **Execução**: Chamamos a função real com a entrada definida.
3.  **Comparação**: Verificamos se o resultado recebido é igual ao esperado.
4.  **Notificação de Falha**: Se forem diferentes, usamos `t.Errorf` para registrar o erro.

### Comandos de Terminal
- `go test`: Executa todos os testes do diretório atual.
- `go test -v`: (Verbose) Mostra o nome de cada função de teste que está sendo executada e o tempo gasto.
- `go test ./...`: Executa todos os testes do projeto, incluindo subpastas (útil para sua pasta de cursos).