package calculadorarecisao_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	calculadora "github.com/EdmilsonRodrigues/calculadora-financeira/calculadora-recisao/calculadora"
)

func TestCalcularSaldoSalário(t *testing.T) {
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
			got := calculadora.CalcularSaldoSalário(tc.salário, tc.dataContratação, tc.dataDemissao, tc.avisoPrevio)
			want := tc.saldo

			assertFloatEqual(t, got, want)
		})
	}
}

func TestCalcularSaldoFérias(t *testing.T) {
	testCases := []struct {
		salário         float64
		saldo           float64
		dataContratação time.Time
		dataDemissao    time.Time
		avisoPrevio     calculadora.AvisoPrevio
		feriasVencidas  bool
	}{
		{1000, 1444.44, calculadora.Date(2024, 1, 1), calculadora.Date(2025, 1, 1), calculadora.Trabalhado, true},
		{2000, 2888.89, calculadora.Date(2024, 1, 4), calculadora.Date(2025, 1, 4), calculadora.Trabalhado, true},
		{1500, 1833.33, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 1), calculadora.Trabalhado, false},
		{1500, 2000, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 12, 1), calculadora.Trabalhado, false},
		{1500, 2166.67, calculadora.Date(2024, 1, 1), calculadora.Date(2026, 1, 1), calculadora.Trabalhado, true},
		{4500, 6000, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado, false},
		{4500, 5500, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Indenizado, false},
		{4500, 5500, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Dispensado, false},
		{4500, 5500, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.NãoCumprido, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Teste de saldo de Férias: %+v %+v", tc.salário, tc.avisoPrevio), func(t *testing.T) {
			got := calculadora.CalcularSaldoFérias(tc.salário, tc.dataContratação, tc.dataDemissao, tc.avisoPrevio, tc.feriasVencidas)
			want := tc.saldo

			assertFloatEqual(t, got, want)
		})
	}
}

func TestCalcularSaldoAviso(t *testing.T) {
	testCases := []struct {
		salário     float64
		saldo       float64
		avisoPrevio calculadora.AvisoPrevio
		justaCausa  bool
	}{
		{1000, 0, calculadora.Trabalhado, false},
		{1000, 1000, calculadora.Indenizado, false},
		{1000, 1000, calculadora.Dispensado, false},
		{1000, -1000, calculadora.NãoCumprido, false},
		{1000, 0, calculadora.Trabalhado, true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Teste de saldo de Aviso: %+v %+v", tc.salário, tc.avisoPrevio), func(t *testing.T) {
			got := calculadora.CalcularSaldoAviso(tc.salário, tc.avisoPrevio, tc.justaCausa)
			want := tc.saldo
			assertFloatEqual(t, got, want)
		})
	}
}

func TestSaldoDécimoTerceiro(t *testing.T) {
	testCases := []struct {
		salário         float64
		saldo           float64
		dataContratação time.Time
		dataDemissao    time.Time
		avisoPrevio     calculadora.AvisoPrevio
	}{
		{1000, 77.08, calculadora.Date(2024, 1, 1), calculadora.Date(2025, 1, 1), calculadora.Trabalhado},
		{2000, 154.17, calculadora.Date(2024, 1, 4), calculadora.Date(2025, 1, 4), calculadora.Trabalhado},
		{1500, 1271.88, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 1), calculadora.Trabalhado},
		{1500, 1387.5, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 12, 1), calculadora.Trabalhado},
		{1500, 115.63, calculadora.Date(2024, 1, 1), calculadora.Date(2026, 1, 1), calculadora.Trabalhado},
		{2500, 2294.88, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado},
		{3000, 2505.32, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Indenizado},
		{3500, 2871.87, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Dispensado},
		{4000, 3214.71, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.NãoCumprido},
		{4500, 3809.59, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado},
		{8000, 6022.05, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado},
		{8500, 6368.58, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Teste de saldo de Décimo Terceiro: %+v %+v", tc.salário, tc.avisoPrevio), func(t *testing.T) {
			got := calculadora.CalcularSaldoDécimoTerceiro(tc.salário, tc.dataContratação, tc.dataDemissao, tc.avisoPrevio)
			want := tc.saldo

			assertFloatEqual(t, got, want)
		})
	}
}

func TestCalcularSaldoFGTS(t *testing.T) {
	testCases := []struct {
		salário         float64
		saldo           float64
		dataContratação time.Time
		dataDemissao    time.Time
		avisoPrevio     calculadora.AvisoPrevio
		motivo          calculadora.Motivo
		justaCausa      bool
	}{
		{1000, 0, calculadora.Date(2024, 1, 1), calculadora.Date(2025, 1, 1), calculadora.Trabalhado, calculadora.DemissãoComJustaCausa, true},
		{2000, 2204.97, calculadora.Date(2024, 1, 4), calculadora.Date(2025, 1, 4), calculadora.Trabalhado, calculadora.FuncionárioPediuDemissão, false},
		{1500, 1954.12, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 1), calculadora.Trabalhado, calculadora.DemissãoSemJustaCausa, false},
		{1500, 2134.44, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 12, 1), calculadora.Trabalhado, calculadora.DemissãoSemJustaCausa, false},
		{1500, 4520.07, calculadora.Date(2024, 1, 1), calculadora.Date(2026, 1, 1), calculadora.Trabalhado, calculadora.DemissãoSemJustaCausa, false},
		{2500, 3557.40, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado, calculadora.DemissãoSemJustaCausa, false},
		{3000, 3908.23, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Indenizado, calculadora.DemissãoSemJustaCausa, false},
		{3500, 4559.61, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Dispensado, calculadora.DemissãoSemJustaCausa, false},
		{4000, 5210.98, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.NãoCumprido, calculadora.DemissãoSemJustaCausa, false},
		{4500, 6403.32, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado, calculadora.DemissãoSemJustaCausa, false},
		{8000, 11383.68, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado, calculadora.DemissãoSemJustaCausa, false},
		{8500, 12095.16, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado, calculadora.DemissãoSemJustaCausa, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Teste de saldo de FGTS: %+v %+v", tc.salário, tc.avisoPrevio), func(t *testing.T) {
			got := calculadora.CalcularSaldoFGTS(tc.salário, tc.dataContratação, tc.dataDemissao, tc.avisoPrevio, tc.motivo, tc.justaCausa)
			want := tc.saldo
			assertFloatEqual(t, got, want)
		})
	}
}

func TestCalcularRecisão(t *testing.T) {
	testCases := []struct {
		salário         float64
		saldo           float64
		dataContratação time.Time
		dataDemissao    time.Time
		avisoPrevio     calculadora.AvisoPrevio
		motivo          calculadora.Motivo
		justaCausa      bool
		FériasVencidas  bool
	}{
		{1000, 1536.94, calculadora.Date(2024, 1, 1), calculadora.Date(2025, 1, 1), calculadora.Trabalhado, calculadora.DemissãoComJustaCausa, true, true},
		{2000, 5463.85, calculadora.Date(2024, 1, 4), calculadora.Date(2025, 1, 4), calculadora.Trabalhado, calculadora.FuncionárioPediuDemissão, false, true},
		{1500, 3833.70, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 1), calculadora.Trabalhado, calculadora.DemissãoSemJustaCausa, false, false},
		{1500, 5567.71, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 12, 1), calculadora.Trabalhado, calculadora.DemissãoSemJustaCausa, false, false},
		{1500, 6964.24, calculadora.Date(2024, 1, 1), calculadora.Date(2026, 1, 1), calculadora.Trabalhado, calculadora.DemissãoSemJustaCausa, false, true},
		{2500, 9185.61, calculadora.Date(2024, 1, 1), calculadora.Date(2024, 11, 30), calculadora.Trabalhado, calculadora.DemissãoSemJustaCausa, false, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Teste de recisão: %+v %+v", tc.salário, tc.avisoPrevio), func(t *testing.T) {
			got := calculadora.CalcularRecisão(calculadora.RequisiçãoRecisão{tc.salário, tc.dataContratação, tc.dataDemissao, tc.avisoPrevio, tc.motivo, tc.justaCausa, tc.FériasVencidas})
			want := tc.saldo
			assertFloatEqual(t, got.SaldoTotal, want)
		})
	}
}

func assertFloatEqual(t testing.TB, got, want float64) {
	t.Helper()
	if math.Abs(got-want) >= 0.0051 {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
