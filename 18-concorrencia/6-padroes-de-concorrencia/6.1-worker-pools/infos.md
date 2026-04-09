## Padrões de Concorrência: Worker Pools

O padrão **Worker Pool** consiste em criar uma fila de tarefas e um número limitado de Goroutines (trabalhadores) para processá-las simultaneamente.

### Por que utilizar?
- **Eficiência**: Permite processar grandes volumes de dados em paralelo.
- **Controle de Recursos**: Evita que o sistema crie Goroutines infinitas, o que poderia causar falta de memória. Você define exatamente quantos "trabalhadores" estarão ativos.

### Canais Direcionais
Uma boa prática em Go é restringir o que uma função pode fazer com um canal:
- `<-chan`: Canal de leitura (apenas recebe dados).
- `chan<-`: Canal de escrita (apenas envia dados).
Isso aumenta a segurança do código, evitando efeitos colaterais indesejados.

### Fluxo de Execução
1. Criamos dois canais (Tarefas e Resultados).
2. Disparamos os Workers (Goroutines).
3. Alimentamos o canal de tarefas.
4. Fechamos o canal de tarefas (avisa aos workers que o trabalho acabou).
5. Coletamos os resultados do canal de saída.