## Estruturas de Controle: Switch

O `switch` é utilizado para simplificar múltiplas seleções condicionais, servindo como uma alternativa mais legível ao uso excessivo de `if/else`.

### Características Principais
- **Break Implícito:** Ao contrário de linguagens como C, o Go não exige o comando `break` ao final de cada caso; ele sai do bloco automaticamente após encontrar a primeira condição verdadeira.
- **Múltiplas Condições:** É possível avaliar múltiplos valores em um único caso: `case "janeiro", "fevereiro", "março":`.
- **Default:** O bloco `default` é executado se nenhuma das condições anteriores for atendida.

### Formatos de Uso
1. **Com variável:** Avalia o valor de uma variável específica.
2. **Sem variável:** Cada `case` contém uma condição booleana independente (funciona como um `if/else` limpo).

### Cláusula Fallthrough
O comando `fallthrough` é utilizado quando desejamos que o código execute o próximo bloco de comando logo abaixo, sem validar a condição do próximo `case`. É um recurso que deve ser usado com cuidado, pois ignora a lógica da próxima condição.