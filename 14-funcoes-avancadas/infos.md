## Funções Avançadas em Go

Esta seção detalha funcionalidades que permitem maior controle sobre a lógica, o fluxo de execução e a manipulação de dados em funções.

### 1. Retorno Nomeado
Permite nomear as variáveis de retorno diretamente na assinatura da função.
- **Vantagem:** O `return` pode ser usado sem argumentos, e as variáveis são pré-inicializadas.

### 2. Funções Variáticas
Funções que aceitam uma quantidade variável de parâmetros usando o prefixo `...`.
- **Regra:** O parâmetro variático deve ser o último da função.

### 3. Funções Anônimas
Funções que não possuem um nome definido. Podem ser executadas imediatamente após a declaração ou armazenadas em variáveis.

### 4. Recursividade
Técnica onde uma função chama a si mesma para resolver problemas que podem ser divididos em subproblemas menores.
- **Importante:** Sempre deve possuir uma condição de parada clara para evitar loops infinitos e erros de memória.

### 5. Cláusula Defer
Adia a execução de um comando ou função até o momento imediatamente anterior ao `return` da função onde foi declarada.
- **Uso Comum:** Limpeza de recursos e fechamento de conexões.

### 6. Panic e Recover
- **Panic:** Interrompe a execução do programa de forma abrupta (usado em erros fatais).
- **Recover:** Recupera um programa que está entrando em pânico, interrompendo a queda do sistema. **Nota:** O `recover` só funciona dentro de uma função com `defer`.

### 7. Closure
Uma função que "envolve" e referencia variáveis que estão fora do seu corpo, mantendo o acesso a elas mesmo após a função original ter terminado sua execução.