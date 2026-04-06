## Concorrência em Go: Goroutines

A **concorrência** em Go é baseada no modelo de CSP (Communicating Sequential Processes), onde as **Goroutines** são a unidade fundamental.

### O que são Goroutines?
São funções ou métodos que são executados de forma independente e simultânea em relação ao fluxo principal do programa. Diferente de *Threads* de sistemas operacionais, as Goroutines são extremamente leves (consomem poucos KB de memória).

### Como funcionam?
- Utilizamos a palavra-chave `go` antes da chamada da função.
- O programa não espera o retorno da função para seguir para a próxima linha.
- **Importante:** Se a função `main` (a thread principal) terminar, todas as outras goroutines serão encerradas imediatamente, independentemente de terem terminado o seu trabalho ou não.

### Concorrência vs Paralelismo
- **Concorrência:** Lidar com muitas coisas ao mesmo tempo (organização).
- **Paralelismo:** Fazer muitas coisas ao mesmo tempo (execução simultânea em múltiplos núcleos de CPU).