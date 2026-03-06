Markdown
## Loops (Estruturas de Repetição)

Em Go, a única estrutura de repetição disponível é o `for`. No entanto, ele é flexível o suficiente para substituir o `while` de outras linguagens.

### Formatos de Uso

1. **Estilo "While"**: Utiliza apenas uma condição de parada.
   ```go
   for i < 10 { ... }
   ```

2. **Estilo Tradicional**: Com inicialização, condição e incremento.

```go
for j := 0; j < 10; j++ { ... }
```

3. **Loop Infinito**: Executa o bloco para sempre (ou até encontrar um break ou return).

```go
for { ... }
```

### A Cláusula `range`

Utilizada para iterar sobre coleções (Arrays, Slices, Maps e Strings).

**Retorno**: Sempre fornece dois valores. O primeiro é o índice (ou chave) e o segundo é o valor.

**Strings**: Ao iterar strings, o valor retornado é uma `rune` (número da tabela Unicode/ASCII).

**Maps**: Itera sobre chaves e valores. Diferente de arrays, a ordem de iteração em maps não é garantida.

**Nota**: Não é possível utilizar `range` diretamente em uma `struct`.