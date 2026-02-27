## Ponteiros

Um **Ponteiro** é uma variável que armazena o **endereço de memória** de outra variável, em vez de armazenar o valor em si.

### Operadores Fundamentais

1.  **Operador de Endereço (`&`)**: Utilizado para obter o endereço de memória de uma variável.
    - Exemplo: `ponteiro = &variavel`
2.  **Operador de Desreferenciação (`*`)**: Utilizado para acessar ou modificar o valor que está armazenado no endereço para o qual o ponteiro aponta.
    - Exemplo: `fmt.Println(*ponteiro)`

### Diferença entre Cópia e Ponteiro

- **Cópia por Valor**: Quando atribuímos uma variável a outra (`v2 = v1`), o Go cria uma cópia exata do dado. Alterar uma não afeta a outra.
- **Referência (Ponteiro)**: Quando atribuímos o endereço (`p = &v1`), ambas as variáveis estão "olhando" para o mesmo lugar na memória. Se o valor na memória mudar, o ponteiro refletirá essa mudança.

### Por que usar Ponteiros?

- **Performance**: Evita copiar grandes estruturas de dados (como `structs` pesadas) na memória ao passá-las para funções.
- **Consistência**: Permite que uma função modifique o valor original de uma variável definida em outro escopo.