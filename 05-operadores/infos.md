## Operadores

### Aritméticos
Operadores básicos para cálculos matemáticos: `+`, `-`, `/`, `*`, `%`.

**Nota:**
Não é possível realizar operações entre variáveis de tipos diferentes (ex: `int8` com `int32`). O Go exige que os tipos sejam idênticos ou que haja uma conversão explícita.
Além disso, a divisão entre dois números inteiros sempre resultará em um número inteiro (ex: `10 / 4` resultará em `2`, e não `2.5`).

### Atribuição 
Formas de atribuir valores a variáveis:
- `var variavel1 string = "Valor"` (Declaração explícita)
- `variavel2 := "Valor"` (Operador curto/inferência de tipo)

### Relacionais
Sempre retornam um valor booleano (`true` ou `false`):
- `==`: Igual a
- `!=`: Diferente de
- `>` : Maior que
- `<` : Menor que
- `>=`: Maior ou igual a
- `<=`: Menor ou igual a

### Lógicos
Usados para combinar condições:
- `&&`: Operador "E" (AND) - Verdadeiro se ambos forem verdadeiros.
- `||`: Operador "OU" (OR) - Verdadeiro se pelo menos um for verdadeiro.
- `! `: Operador de Negação (NOT) - Inverte o valor booleano.

### Unários
Formas simplificadas de alterar o valor de uma variável:
- `++`: Incremento (soma 1)
- `--`: Decremento (subtrai 1)
- `+=`, `-=`, `*=`, `/=`, `%=`: Operadores de atribuição composta (ex: `n += 5` é o mesmo que `n = n + 5`).

### Ternários (Atenção)
O **Go não possui operador ternário** (aquele `condicao ? true : false` comum em C# ou Java). Para manter o código legível e explícito, utiliza-se obrigatoriamente a estrutura `if/else`.