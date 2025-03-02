package main_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	calculadora "github.com/EdmilsonRodrigues/calculadora-recisao"
)

func TestSaldoSalario(t *testing.T) {
	testCases := []struct {
		salario         float64
		saldo           float64
		dataContratação time.Time
		dataDemissao    time.Time
	}{
		{1000, 100, calculadora.Date(2024, 1, 1), calculadora.Date(2025, 1, 1)},
		{2000, 400, calculadora.Date(2024, 1, 4), calculadora.Date(2025, 1, 4)},
		{1500, 50, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 1)},
		{1500, 1550, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 12, 1)},
		{1500, 300, calculadora.Date(2024, 1, 1), calculadora.Date(2026, 1, 1)},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Teste de saldo Salário: %v", tc.salario), func(t *testing.T) {
			got := calculadora.SaldoSalario(tc.salario, tc.dataContratação, tc.dataDemissao)
			want := tc.saldo

			assertFloatEqual(t, got, want)
		})
	}
}

func assertEqual[C comparable](t testing.TB, got, want C) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertFloatEqual(t testing.TB, got, want float64) {
	t.Helper()
	if math.Abs(got-want) > 0.005 {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
