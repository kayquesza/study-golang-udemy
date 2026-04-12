## Testes Avançados: Sub-testes e Organização

Neste módulo, exploramos como organizar testes complexos que envolvem diferentes tipos ou comportamentos sob uma mesma função de teste principal.

### O Poder do `t.Run`
O método `t.Run` permite a criação de sub-testes. Isso traz vantagens como:
- **Isolamento**: Cada sub-teste pode ser executado individualmente via terminal.
- **Legibilidade**: O código do teste fica organizado por contexto ou tipo de dado.
- **Relatórios**: Falhas são reportadas com o caminho completo (ex: `TestArea/Circulo`).

### Controle de Fluxo: Errorf vs Fatalf
- **t.Errorf**: Registra a falha, mas permite que o código continue executando.
- **t.Fatalf**: Registra a falha e interrompe a execução do sub-teste ou teste atual no mesmo instante.

### Testando Interfaces
Através de interfaces, podemos criar testes que validam o contrato (comportamento) em vez da implementação específica, permitindo testar estruturas diferentes (como Retângulos e Círculos) de forma coesa.