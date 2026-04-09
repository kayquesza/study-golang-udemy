## Padrões de Concorrência: Generator (Gerador)

O padrão **Generator** é utilizado para encapsular a criação de uma Goroutine e de um canal, retornando esse canal para que o chamador possa consumir dados de forma simples.

### Características Principais
- **Abstração**: O chamador não precisa saber que existe uma Goroutine rodando internamente.
- **Canais Direcionais**: O retorno costuma ser um canal de "apenas leitura" (`<-chan`), garantindo que o consumidor não interfira no envio de dados.
- **Autonomia**: A função geradora é responsável por iniciar seu próprio fluxo de trabalho.

### Vantagens
- **Código Limpo**: A função `main` fica muito mais legível, focando apenas no processamento dos dados recebidos.
- **Desacoplamento**: A lógica de produção de dados fica isolada da lógica de consumo.

### Exemplo de Uso
Ideal para sistemas que precisam de um fluxo contínuo de informações, como leitores de sensores, geradores de IDs únicos ou streams de logs.