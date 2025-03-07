package calculadorarecisao

import "time"

func CalcularSaldoFérias(salario float64, dataContratacao, dataDemissao time.Time, avisoPrevio AvisoPrevio, feriasVencidas bool) (saldoFerias float64) {
	mesesTrabalhados := getMesesTrabalhados(dataContratacao, dataDemissao, avisoPrevio)
	for mesesTrabalhados > 12 {
		mesesTrabalhados -= 12
	}
	if feriasVencidas {
		mesesTrabalhados += 12
	}
	saldoFerias = getSaldoFerias(salario, mesesTrabalhados)
	return
}

func getMesesTrabalhados(dataInicio, dataFim time.Time, avisoPrevio AvisoPrevio) (mesesTrabalhados int) {
	if avisoPrevio == Trabalhado {
		dataFim = dataFim.AddDate(0, 0, getTempoDeAviso(dataInicio, dataFim))
	}
	mesesTrabalhados = int(dataFim.Sub(dataInicio).Hours() / (24 * 30))
	return
}

func getSaldoFerias(salario float64, mesesTrabalhados int) (saldoFerias float64) {
	saldoFerias = (salario * float64(mesesTrabalhados) / 12) * 4 / 3
	return
}
