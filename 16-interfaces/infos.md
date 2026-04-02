## Interfaces

Interfaces em Go são conjuntos de assinaturas de métodos. Elas definem um **comportamento** (o que um tipo faz) em vez de focar na estrutura (o que um tipo é).

### Conceitos Chave
- **Contratos**: Se uma struct possui todos os métodos definidos em uma interface, dizemos que ela implementa essa interface de forma implícita.
- **Polimorfismo**: Permite que diferentes tipos sejam tratados como o mesmo tipo de interface, desde que cumpram o contrato.
- **Assinatura**: Uma interface contém apenas o nome do método, parâmetros de entrada e tipos de retorno.

### Interface Vazia (`interface{}`)
Representa um tipo que pode conter qualquer valor. Como não exige nenhum método, todos os tipos a satisfazem.
- **Uso**: Útil em funções que lidam com dados desconhecidos (ex: `fmt.Println`), mas deve ser evitada para manter a segurança do código.

### Cálculo de Área (Exemplo Matemático)
A interface `forma` garante que qualquer struct que implemente o método `area()` possa ter seu valor calculado, seja um retângulo ou um círculo, onde: