## Composição (Pseudo-herança)

Em Go, não existe herança entre structs. Em vez disso, utilizamos a **composição** através de **campos anônimos**.

### Campos Anônimos
Ao inserir uma struct dentro de outra sem definir um nome para o campo, todos os campos da struct interna são "promovidos" para a struct externa.

**Exemplo:**
```go
type pessoa struct {
    nome string
}

type estudante struct {
    pessoa // Campo anônimo
    curso  string
}