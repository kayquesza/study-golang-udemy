## Funções

Uma série de passos que serão seguidos pelo programa.

As funções em **Go** podem ter *paramêtros* e *retornos*.
- **Paramêtros** são dados que são colocados na função para que ela funcione.
- **Retornos** é o que a função retorna para o usuário.

Sintaxe de uma função:
```go
    func Funcao(<parametros>) <retorno> {
        return <o que deve retornar>
    }
```

**Exemplo* de sintaxe:
```go
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
```
**Nota:**
- `n1 int8, n2 int8` correspondem aos *paramêtros*.
- `int8` correspode ao tipo de *retorno*.
- `return n1 + n2` corresponde a uma operação que será feita e retornada ao usuário.