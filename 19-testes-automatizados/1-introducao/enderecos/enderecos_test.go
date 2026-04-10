// O teste de uma função nunca fica no mesmo arquivo da função em si

// Esse arquivo terá um 'teste de unidade'
// Testará uma pequena parte do código
package enderecos

import "testing"

type cenarioDeTest struct {
	enderecoInserido string
	retornoEsperado  string
}

// 't *testing.T' é o parâmetro e assinatura de uma função teste
// A função obrigatoriamente precisa começar com 'Test' e, como boa prática,
// em seguida, recebe o nome da função que será testada, como por exemplo:
// 'Test' + 'TipoDeEndereco':
func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTest{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada da Serra", "Estrada"},
		{"RUA FRATES", "Rua"},
		{"AVENIDA REBOLÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}
