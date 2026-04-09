## Padrões de Concorrência: Multiplexador (Fan-In)

O padrão **Multiplexador** (também conhecido como *Fan-In*) consiste em fundir dois ou mais canais de entrada num único canal de saída.

### Como Funciona
- **Entrada Múltipla**: Recebe diversos canais (geralmente gerados pelo padrão *Generator*).
- **Processamento Central**: Utiliza uma goroutine interna com um `select` para monitorizar todas as entradas simultaneamente.
- **Saída Única**: Consolida todas as mensagens recebidas e envia-as para um único canal de retorno.

### Vantagens
- **Simplicidade no Consumo**: Quem consome os dados (a função `main`, por exemplo) não precisa de saber quantas fontes existem; apenas lê de um único canal.
- **Não Bloqueante**: Graças ao `select`, o multiplexador não fica preso à espera de uma fonte lenta se outra fonte já tiver dados prontos.

### Exemplo de Aplicação
Essencial em sistemas que agregam logs de diferentes servidores, monitorizam múltiplos sensores ou recolhem resultados de diferentes motores de busca para apresentar ao utilizador final.