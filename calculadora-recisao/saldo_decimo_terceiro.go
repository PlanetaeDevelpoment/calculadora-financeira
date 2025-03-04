package calculadorarecisao

import "time"


func SaldoDécimoTerceiro(salário float64, dataContratação, dataDemissao time.Time, avisoPrevio AvisoPrevio) float64 {
	mesesTrabalhados := getMesesTrabalhados(dataContratação, dataDemissao, avisoPrevio)
	for {
		if mesesTrabalhados < 12 {
			break			
		}
		mesesTrabalhados -= 12
	}
	salarioBruto := salário * float64(mesesTrabalhados / 12)
	return retirarDescontos(salarioBruto)
}