package calculadorarecisao

import (
	"math"
	"time"
)

type rangeSalarial struct {
	min, max, porcentagem float64
}

// aviso Prévio idenizado só desconta INSS
// Férias não descota nada
// 13º salário desconta ambos

const (
	baseAviso   = 30
	baseSalário = 30
)

var (
	tabelaINSS = []rangeSalarial{
		{0, 1518, 7.5},
		{1518.01, 2793.88, 9},
		{2793.89, 4190.83, 12},
		{4190.84, 8157.41, 14},
	}
	tabelaIRS = []rangeSalarial{
		{0, 2259.20, 0},
		{2259.21, 2826.65, 7.5},
		{2826.66, 3751.05, 15},
		{3751.06, 4664.68, 22.5},
		{4664.69, math.Inf(1), 27.5},
	}
)

func CalcularSaldoSalário(salário float64, dataContratação, dataDemissao time.Time, avisoPrevio AvisoPrevio) (saldo float64) {
	porcentagemSaldo := getPorcentagemSaldo(dataContratação, dataDemissao, avisoPrevio)
	salarioBruto := salário * porcentagemSaldo
	saldo = retirarDescontos(salarioBruto)
	return
}

func getPorcentagemSaldo(dataContratação, dataDemissao time.Time, avisoPrevio AvisoPrevio) float64 {
	tempoDeAviso := 0
	if avisoPrevio == Trabalhado {
		tempoDeAviso = getTempoDeAviso(dataContratação, dataDemissao)
	}
	dataAposAviso := dataDemissao.AddDate(0, 0, tempoDeAviso)
	return (float64(dataAposAviso.Day()) / baseSalário)
}

func getTempoDeAviso(dataContratação, dataDemissao time.Time) int {
	anosContribuição := getAnosContribuição(dataContratação, dataDemissao)
	return baseAviso + (3 * anosContribuição)
}

func getAnosContribuição(inicio, fim time.Time) int {
	return int(fim.Sub(inicio).Hours() / (24 * 365))
}

func retirarDescontos(salarioBruto float64) (saldo float64) {
	saldo = descontar(salarioBruto)
	return
}

func descontar(salarioBruto float64) (descontado float64) {
	descontado = salarioBruto - getDescontoINSS(salarioBruto)
	descontado -= getDescontoIR(descontado)
	return
}

func getDescontoINSS(salarioBruto float64) float64 {
	return getDesconto(salarioBruto, tabelaINSS)
}

func getDescontoIR(salarioBruto float64) float64 {
	return getDesconto(salarioBruto, tabelaIRS)
}

func getDesconto(salarioBruto float64, tabela []rangeSalarial) (desconto float64) {
	for _, rangeSalarial := range tabela {
		desconto += getDescontoParaRange(salarioBruto, rangeSalarial)
		if salarioBruto <= rangeSalarial.max {
			break
		}
	}
	return
}

func getDescontoParaRange(salarioBruto float64, rangeSalarial rangeSalarial) (desconto float64) {
	menor := salarioBruto
	if rangeSalarial.max < menor {
		menor = rangeSalarial.max
	}
	desconto = (menor - rangeSalarial.min) * (rangeSalarial.porcentagem / 100)
	return
}
