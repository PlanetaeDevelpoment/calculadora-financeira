package calculadorarecisao

import "time"

const jurosMensal = 0.0025

func CalcularSaldoFGTS(salário float64, dataContratação, dataDemissao time.Time, avisoPrevio AvisoPrevio, motivo Motivo, justaCausa bool) (saldo float64) {
	if justaCausa {
		return
	}
	mesesTrabalhados := getMesesTrabalhados(dataContratação, dataDemissao, avisoPrevio)
	saldo = getCalcularSaldoFGTS(mesesTrabalhados, salário)
	if motivo == DemissãoSemJustaCausa {
		saldo *= 1.4
	}
	return
}

func getCalcularSaldoFGTS(mesesTrabalhados int, salario float64) (saldo float64) {
	base := salario / 12
	for range mesesTrabalhados {
		saldo += base
		saldo *= (1 + jurosMensal)
	}
	return saldo
}
