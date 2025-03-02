package main_test

import (
	"fmt"
	"testing"
	"time"

	calculadora "github.com/EdmilsonRodrigues/calculadora-recisao"
)

func TestSaldoSalario(t *testing.T) {
	testCases := []struct {
		salario     float64
		saldo       float64
		dataContratação time.Time
		dataDemissao time.Time
	}{
		{1000, 100, calculadora.Date(2024, 1, 1), calculadora.Date(2025, 1, 1)},
		{2000, 400, calculadora.Date(2024, 1, 4), calculadora.Date(2025, 1, 4)},
		{1500, 0, calculadora.Date(2024, 1, 1), calculadora.Date(2025, 12, 1)},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Teste de saldo Salário: %v", tc.salario), func(t *testing.T) {
			got := calculadora.SaldoSalario(tc.salario, tc.dataDemissao)
			want := tc.saldo

			assertEqual(t, got, want)
		})
	}
}

func assertEqual[C comparable](t testing.TB, got, want C) {
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
