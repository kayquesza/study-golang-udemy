## Sincronização: WaitGroup

O pacote `sync` fornece a estrutura **WaitGroup**, essencial para gerenciar o ciclo de vida de múltiplas Goroutines e garantir que a função principal espere a conclusão de todas as tarefas.

### Os Três Pilares
1.  **Add(int)**: Incrementa o contador interno com o número de Goroutines que serão disparadas.
2.  **Done()**: Decrementa o contador em 1. Deve ser chamado obrigatoriamente ao fim de cada Goroutine.
3.  **Wait()**: Bloqueia a execução do programa até que o contador atinja zero.

### Observações Técnicas
- **Encapsulamento**: O uso de funções anônimas para chamar o `Done()` evita que a lógica de sincronização "polua" a função de negócio (como a função `escrever`).
- **Risco de Deadlock**: Se o número de chamadas ao `Done()` for menor que o valor definido no `Add()`, o programa entrará em pânico por nunca conseguir sair do estado de espera.