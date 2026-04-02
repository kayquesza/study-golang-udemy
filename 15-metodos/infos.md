## Métodos

Em Go, **Métodos** são funções que estão anexadas a um tipo específico (geralmente uma `struct`). Eles representam as ações que aquele tipo pode executar.

### Sintaxe e o "Receiver"
A principal diferença na declaração é a presença do **receptor** (receiver) antes do nome da função.

```go
func (variavel Tipo) NomeDoMetodo() {
    // corpo do método
}