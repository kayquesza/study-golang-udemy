## Manipulação de JSON em Go

O pacote padrão `encoding/json` é utilizado para converter estruturas de dados de Go para JSON (**Marshal**) e vice-versa (**Unmarshal**).

### Marshal (Codificação)
Transforma uma Struct ou Map em um slice de bytes (`[]byte`) formatado como JSON.

- **Campos Exportáveis**: Apenas campos que começam com letra **Maiúscula** podem ser convertidos para JSON.
- **Struct Tags**: Permitem customizar o nome das chaves no JSON resultante.
  - Exemplo: `Nome string `json:"nome_completo"``

### Mapas vs Structs
- **Structs**: São preferíveis quando o formato do dado é fixo e conhecido. Oferecem maior segurança de tipo.
- **Maps**: São úteis para dados dinâmicos ou quando não queremos definir uma estrutura rígida para o JSON.

### Dica de Depuração
Como o `json.Marshal` retorna bytes, para visualizar o resultado no terminal como texto, utilizamos:
`fmt.Println(bytes.NewBuffer(resultadoBytes))` ou `string(resultadoBytes)`.