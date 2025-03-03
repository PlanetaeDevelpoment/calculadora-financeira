package calculadorarecisao_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	calculadora "github.com/EdmilsonRodrigues/calculadora-financeira/calculadora-recisao"
)

func TestSaldoSalário(t *testing.T) {
	testCases := []struct {
		salário         float64
		saldo           float64
		dataContratação time.Time
		dataDemissao    time.Time
		avisoPrevio     calculadora.AvisoPrevio
	}{
		{1000, 92.5, calculadora.Date(2024, 1, 1), calculadora.Date(2025, 1, 1), calculadora.Trabalhado},
		{2000, 370, calculadora.Date(2024, 1, 4), calculadora.Date(2025, 1, 4), calculadora.Trabalhado},
		{1500, 46.25, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 1), calculadora.Trabalhado},
		{1500, 1433.27, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 12, 1), calculadora.Trabalhado},
		{1500, 277.5, calculadora.Date(2024, 1, 1), calculadora.Date(2026, 1, 1), calculadora.Trabalhado},
		{2500, 2294.88, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado},
		{3000, 2710.04, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Indenizado},
		{3500, 3090.04, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Dispensado},
		{4000, 3464.04, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.NãoCumprido},
		{4500, 3809.59, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado},
		{8000, 6022.05, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado},
		{8500, 6368.58, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Teste de saldo Salário: %v", tc.salário), func(t *testing.T) {
			got := calculadora.SaldoSalário(tc.salário, tc.dataContratação, tc.dataDemissao, tc.avisoPrevio)
			want := tc.saldo

			assertFloatEqual(t, got, want)
		})
	}
}

func TestSaldoFérias(t *testing.T) {
	testCases := []struct {
		salário         float64
		saldo           float64
		dataContratação time.Time
		dataDemissao    time.Time
		avisoPrevio     calculadora.AvisoPrevio
	}{
		{1000, 1416.67, calculadora.Date(2024, 1, 1), calculadora.Date(2025, 1, 1), calculadora.Trabalhado},
		{2000, 2833.33, calculadora.Date(2024, 1, 4), calculadora.Date(2025, 1, 4), calculadora.Trabalhado},
		{1500, 2000, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 1), calculadora.Trabalhado},
		{1500, 2000, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 12, 1), calculadora.Trabalhado},
		{1500, 4166.67, calculadora.Date(2024, 1, 1), calculadora.Date(2026, 1, 1), calculadora.Trabalhado},
		{4500, 6000, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado},
		{4500, 5500, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Indenizado},
		{4500, 5500, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Dispensado},
		{4500, 5500, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.NãoCumprido},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Teste de saldo de Férias: %+v", tc.salário), func(t *testing.T) {
			got := calculadora.SaldoFérias(tc.salário, tc.dataContratação, tc.dataDemissao, tc.avisoPrevio)
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
