// O teste de uma função nunca fica no mesmo arquivo da função em si

// Esse arquivo terá um 'teste de unidade'
// Testará uma pequena parte do código
package enderecos_test

import (
	"introducao-testes/enderecos"
	"testing"
)

type cenarioDeTest struct {
	enderecoInserido string
	retornoEsperado  string
}

// 't *testing.T' é o parâmetro e assinatura de uma função teste
// A função obrigatoriamente precisa começar com 'Test' e, como boa prática,
// em seguida, recebe o nome da função que será testada, como por exemplo:
// 'Test' + 'TipoDeEndereco':
func TestTipoDeEndereco(t *testing.T) {

	t.Parallel() // O teste será rodado em paralelo com o outro teste (ambos precisam ter essa característica)

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
		retornoRecebido := enderecos.TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}

// 'TestQualquer' apenas para entender melhor a autilização do 'go test -v' no terminal
func TestQualquer(t *testing.T) {
	t.Parallel() // O teste será rodado em paralelo com o outro teste (ambos precisam ter essa característica)
	if 3 < 2 {
		t.Errorf("Teste quebrou!")
	}
}

// 'go test --cover' mostra a porcentagem de quanto a função está sendo "coberta"
// 'go test --coverprofile example.txt' gera um arquivo relatando tudo que está coberto ou não
// 'go tool cover --func=example.txt' lê melhor o arquivo gerado anteriormente
// 'go tool cover --html=example.txt' gera um arquivo temporário html mostrando todo conteúdo necessário (o que está coberto e o que não está)
