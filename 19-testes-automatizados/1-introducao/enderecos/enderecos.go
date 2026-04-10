package enderecos

import "strings"

// 'TipoDeEndereco' verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"} // tipos válidos para começo de endereço

	enderecoEmLetraMiniscula := strings.ToLower(endereco)                        // converte para a letra minuscula (a primeira palavra capturada)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMiniscula, " ")[0] // vendo se a primeira palavra, com a letra miniscula, bate com algum dos tipos validos

	enderecoTemUmTipoValido := false // inicia como falso, até que, se o item em questão contiver a palavra válida, torna-se um tipo válido
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco) // se for válido, retorna a palavra em questão e deixa a inicial maiuscula
	}

	return "Tipo Inválido"

}
