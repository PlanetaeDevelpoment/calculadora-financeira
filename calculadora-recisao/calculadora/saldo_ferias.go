package calculadorarecisao

import "time"



func SaldoFérias(salario float64, dataContratacao, dataDemissao time.Time, avisoPrevio AvisoPrevio) (saldoFerias float64) {
	mesesTrabalhados := getMesesTrabalhados(dataContratacao, dataDemissao, avisoPrevio)
	saldoFerias = getSaldoFerias(salario, mesesTrabalhados)
	return
}

func getMesesTrabalhados(dataContratacao, dataDemissao time.Time, avisoPrevio AvisoPrevio) (mesesTrabalhados int) {
	if avisoPrevio == Trabalhado {
		dataDemissao = dataDemissao.AddDate(0, 0, getTempoDeAviso(dataContratacao, dataDemissao))
	}
	mesesTrabalhados = int(dataDemissao.Sub(dataContratacao).Hours() / (24 * 30))
	return	
}

func getSaldoFerias(salario float64, mesesTrabalhados int) (saldoFerias float64) {
	saldoFerias = (salario * float64(mesesTrabalhados) / 12) * 4/3
	return
}