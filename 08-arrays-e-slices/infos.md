## Arrays 

### Informações Importantes
- **Tamanho Fixo:** Uma vez declarado com `[n]`, o tamanho não muda.
- **Burlar o Tamanho:** Usar `[...]` faz o compilador contar os elementos para você, mas o array resultante continua sendo fixo.
- **Tipagem Única:** Todos os elementos devem ser do mesmo tipo.
- **Zero Values:** Elementos não inicializados assumem o valor padrão (0 para `int`, `""` para `string`, `false` para `bool`).

## Slices

### Informações Importantes
- **Flexíveis:** Não possuem tamanho fixo na declaração (`[]tipo`).
- **Estrutura de Ponteiro:** Um slice não armazena dados; ele aponta para um array interno.
- **Reflexão:** Slices e Arrays são tipos diferentes para o Go (um `[5]int` é diferente de um `[]int`).

### Funções Essenciais
- `append(slice, valor)`: Adiciona um elemento e, se necessário, aumenta a capacidade do array interno.
- `make([]tipo, len, cap)`: Cria um slice com controle total sobre o tamanho inicial e a capacidade máxima.
- `len()`: Retorna o número de elementos atuais.
- `cap()`: Retorna a capacidade máxima antes de uma nova realocação de memória.