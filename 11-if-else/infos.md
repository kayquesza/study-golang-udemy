## Estruturas de Controle: if/else

O `if` em Go é a ferramenta fundamental para criar fluxos condicionais no programa.

### Sintaxe Básica
Diferente de outras linguagens, o Go não utiliza parênteses para envolver a condição.

```go
if numero > 15 {
    // bloco de código
} else if numero < 0 {
    // bloco de código
} else {
    // bloco de código
}
```

### If com Inicialização (If Init)

Uma funcionalidade poderosa do Go é a capacidade de inicializar uma variável diretamente na declaração do if.

**Sintaxe:** if variavel := valor; condicao { ... }

### Vantagens:

**Escopo Limitado:** A variável criada existe apenas dentro do bloco if/else.

**Código Limpo:** Ideal para capturar retornos de funções que só serão usados naquela validação específica (muito comum no tratamento de erros).

### Regras Importantes

A condição deve resultar em um valor booleano (true ou false). O Go não possui o conceito de valores "truthy" ou "falsy" (como 0 ou strings vazias sendo tratados como false).

As chaves {} são obrigatórias, mesmo para blocos de apenas uma linha.