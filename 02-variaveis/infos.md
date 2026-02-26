## Variáveis e Seus Tipos

Como **Go** é fortemente tipado, **toda** variável tem obrigatoriamente um tipo, mas não necessariamente precisa estar explicíto. 

Toda variável *deve* ser utilizada. Caso não seja, o **Go** avisará com um erro informando a variável e falta de seu uso.

### Valores Zero (Zero Values)
Se declarada uma variável sem atribuição de um valor, ela não fica *nula* ou com *lixo de memória*. Ela assume um **valor zero** padrão:
- `int`: 0
- `string`: "" (string vazia)
- `bool`: false
- `float`: 0.0

### Escopo e a Restrição do `:=`
- O operador curto `:=`**só pode ser usado dentro de funções**.
- Para declarar variáveis no escopo global (fora da `func main`), se é obrigado a usar a palavra-chave `var`.