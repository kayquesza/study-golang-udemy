package main

import (
	"fmt"
)

// 1. A Interface (O Contrato)
type ScannerSeguranca interface {
	Escanear() error
}

// 2. As Structs (Os Dados)
type SAST struct { // Static Application Security Testing
	Linguagem string
}

type DAST struct { // Dynamic Application Security Testing
	URL string
}

type ScannerRede struct {
	IPRange string
}

// 3. Implementação os métodos (Cumprindo o Contrato)
func (s SAST) Escanear() error {
	fmt.Printf("Executando análise estática no código %s...\nBuscando vulnerabilidades no código-fonte.\n\n", s.Linguagem)
	return nil
}

func (d DAST) Escanear() error {
	fmt.Printf("Iniciando teste dinâmico na aplicação %s...\nSimulando ataques em tempo de execução.\n\n", d.URL)
	return nil
}

func (sr ScannerRede) Escanear() error {
	fmt.Printf("Varrendo a rede no range %s...\nVerificando portas abertas e serviços vulneráveis.", sr.IPRange)
	return nil
}

// 4. A função que usa a interface (O 'Maestro')
func IniciarPipelineSeguranca(scanners []ScannerSeguranca) {

	fmt.Printf("\n--- INICIANDO PIPELINE DE SEGURANÇA ---\n\n")

	for _, s := range scanners {
		s.Escanear()
	}

	fmt.Printf("\n\n--- PIPELINE FINALIZADO ---\n\n")

}

func main() {

	// Instâncias com dados de configuração
	sast := SAST{Linguagem: "Golang"}
	dast := DAST{URL: "https://minha-api.com.br"}
	rede := ScannerRede{IPRange: "192.168.0.1/24"}

	// O evento que queremos registrar
	scanners := []ScannerSeguranca{sast, dast, rede}

	// Executando a auditoria em diferentes destinos
	IniciarPipelineSeguranca(scanners)

}
