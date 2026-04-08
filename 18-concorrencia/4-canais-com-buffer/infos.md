## Canais com Buffer

Canais com Buffer permitem o envio de dados sem que haja necessariamente um receptor pronto no momento exato do envio. Eles possuem uma capacidade limitada definida na criação.

### Funcionamento
- **Sintaxe**: `canal := make(chan tipo, capacidade)`
- **Envio**: O envio só bloqueia a execução quando o buffer atinge sua capacidade máxima.
- **Recebimento**: O recebimento só bloqueia quando o buffer está completamente vazio.

### Diferença Prática
- **Canal sem Buffer**: A sincronização é imediata (Rendezvous). O emissor e o receptor devem estar prontos simultaneamente.
- **Canal com Buffer**: A sincronização é assíncrona até que o limite do buffer seja atingido. Funciona como uma "fila" temporária.

### Risco de Deadlock
Mesmo com buffer, o programa entrará em deadlock se tentarmos enviar dados para um canal cheio ou receber de um canal vazio dentro da mesma Goroutine (geralmente a `main`), pois não haverá outro fluxo para desbloquear a operação.