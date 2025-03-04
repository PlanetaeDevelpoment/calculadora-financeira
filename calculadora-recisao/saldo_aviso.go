package calculadorarecisao


func SaldoAviso(salario float64, avisoPrevio AvisoPrevio) (aviso float64) {
	switch avisoPrevio {
	case Trabalhado:
		return
	case NãoCumprido:
		return -salario
	default:
		return salario
	}	
}