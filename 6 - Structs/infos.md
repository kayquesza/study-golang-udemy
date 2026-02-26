## Structs

**Structs**: uma coleção de campos, onde cada campo tem um *nome* e um *tipo*.

Sintaxe de uma **Struct**:

```go
type <nomeDaStruct> struct {
    <campo1> <tipo>
    <campo2> <tipo>
}
```

**Exemplo:**
```go
type usuario struct {
	nome  string
	idade uint8
}
``` 