## Estruturas de Concorrência: Select

O `select` é uma estrutura de controle exclusiva para canais. Ele permite que uma Goroutine aguarde múltiplas operações de comunicação, executando o primeiro caso que estiver pronto.

### Por que usar o Select?
Em sistemas de alta performance (como ferramentas de segurança ou backends), frequentemente lidamos com fontes de dados que possuem velocidades diferentes. O `select` evita que uma resposta rápida (ex: 500ms) fique "presa" esperando uma resposta lenta (ex: 2s).

### Funcionamento
- O `select` bloqueia a execução até que um de seus casos possa prosseguir.
- Se múltiplos casos estiverem prontos ao mesmo tempo, o Go escolhe um de forma **aleatória** para garantir justiça (fairness).
- **Default Case**: Assim como no `switch`, podemos usar um `default` para que o `select` seja não-bloqueante (ele executa o default se nenhum canal estiver pronto).

### Exemplo Prático
No código desenvolvido, temos:
1. **Canal 1**: Envia dados a cada 0.5s.
2. **Canal 2**: Envia dados a cada 2s.
O `select` garante que as mensagens do Canal 1 sejam processadas assim que chegarem, sem esperar o ciclo de 2 segundos do Canal 2.