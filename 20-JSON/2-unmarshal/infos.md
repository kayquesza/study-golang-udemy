## Manipulação de JSON em Go

O pacote padrão `encoding/json` é utilizado para converter estruturas de dados de Go para JSON (**Marshal**) e vice-versa (**Unmarshal**).

### Unmarshal (Decodificação)
Transforma um slice de bytes (`[]byte`) contendo dados em formato JSON numa Struct ou num Map de Go.

- **Uso de Ponteiros**: É obrigatório passar o endereço de memória da variável de destino (ex: `&usuario`), para que a função possa preencher os campos originais.
- **Mapeamento por Tags**: O processo utiliza as `json tags` definidas na struct para saber em qual campo colocar cada valor do JSON.
- **Tratamento de Strings**: Geralmente convertemos strings para `[]byte` antes de realizar a operação.

### Diferenças no Consumo
- **Para Structs**: O Unmarshal valida os tipos. Se tentares colocar um texto num campo `int`, a função retornará um erro.
- **Para Maps**: Oferece flexibilidade total, mas exige cuidado com a conversão de tipos (type assertion) ao ler os dados posteriormente.