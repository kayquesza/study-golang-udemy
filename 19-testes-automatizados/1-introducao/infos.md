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

### Table Driven Tests (Testes Orientados a Dados)

Esta é a forma recomendada de escrever testes em Go quando temos uma função que precisa ser validada com diversos conjuntos de entradas e saídas.

#### Vantagens
- **Manutenibilidade**: Adicionar novos casos de teste exige apenas uma nova linha em uma struct.
- **Legibilidade**: O código do teste fica separado dos dados do teste.
- **Cobertura**: Facilita a visualização de quais cenários já foram cobertos (casos de sucesso, erro, strings vazias, etc).

#### Estrutura Básica
1. Definimos uma `struct` local para representar o cenário (entrada e saída esperada).
2. Criamos um `slice` dessa struct com todos os casos que queremos validar.
3. Usamos um loop `range` para executar a função e comparar os resultados.