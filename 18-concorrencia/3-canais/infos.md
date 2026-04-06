## Canais (Channels)

Canais são os condutos que permitem a comunicação e sincronização entre Goroutines. Eles garantem que os dados trafeguem de forma segura entre diferentes fluxos de execução.

### Operações Básicas
- **Criação**: `canal := make(chan tipo)`
- **Enviar dado**: `canal <- valor` (Bloqueia até que alguém receba)
- **Receber dado**: `valor := <-canal` (Bloqueia até que alguém envie)

### Sincronização e Fechamento
- **Bloqueio**: O canal força as goroutines a esperarem umas pelas outras, servindo como um mecanismo de sincronização natural.
- **Fechamento**: O comando `close(canal)` avisa a quem está recebendo que não haverá mais dados. Tentar enviar para um canal fechado causa um pânico.
- **Iteração**: O uso de `for valor := range canal` é a forma recomendada de ler todos os dados até que o canal seja fechado.

### O Padrão "Comma Ok"
Podemos verificar se um canal ainda está aberto ao receber um dado:
`valor, aberto := <-canal`
Se `aberto` for `false`, o canal foi fechado e o valor retornado é o "Zero Value" do tipo.