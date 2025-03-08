package calculadorarecisao

import (
	"time"
)


func CalcularSaldoDécimoTerceiro(salário float64, dataContratação, dataDemissao time.Time, avisoPrevio AvisoPrevio) float64 {
	mesesTrabalhados := getMesesTrabalhadosDesdeJaneiro(dataDemissao, avisoPrevio)
	salarioBruto := salário * float64(mesesTrabalhados) / 12
	return retirarDescontos(salarioBruto)
}

func getMesesTrabalhadosDesdeJaneiro(dataDemissao time.Time, avisoPrevio AvisoPrevio) int {
	dataInicial := Date(dataDemissao.Year() - 1, 12, 15)
	return getMesesTrabalhados(dataInicial, dataDemissao, avisoPrevio)
}