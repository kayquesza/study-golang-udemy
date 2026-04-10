// O teste de uma função nunca fica no mesmo arquivo da função em si

// Esse arquivo terá um 'teste de unidade'
// Testará uma pequena parte do código
package enderecos

import "testing"

// 't *testing.T' é o parâmetro e assinatura de uma função teste
// A função obrigatoriamente precisa começar com 'Test' e, como boa prática,
// em seguida, recebe o nome da função que será testada, como por exemplo:
// 'Test' + 'TipoDeEndereco':
func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Rua Paulista"

	tipoDeEnderecoEsperado := "Avenida"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s.", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}

}

// Para rodar o test no terminal:
// 'go test'
