## Maps

Os **Maps** são coleções de pares **chave-valor**. São a estrutura ideal para buscas rápidas quando você possui um identificador único (como um ID, CPF ou Nome) e precisa encontrar um dado associado.

### Sintaxe Básica
A declaração segue o formato: `map[tipo_da_chave]tipo_do_valor`.

```go
usuario := map[string]string{
    "nome":      "Kayque",
    "faculdade": "GRAN",
}
```

### Regras e Características

**Chaves Únicas:** Não podem existir duas chaves idênticas. Se você atribuir um novo valor a uma chave existente, o valor antigo é sobrescrito.

**Tipagem Fixa:** Uma vez definido, o map só aceita chaves e valores dos tipos especificados.

**Acesso:** O acesso é feito via colchetes: valor := mapa["chave"].

### Operações Comuns

**Adicionar/Alterar:** mapa["nova_chave"] = "valor"

**Remover:** delete(mapa, "chave")