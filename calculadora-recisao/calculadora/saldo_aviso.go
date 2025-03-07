package calculadorarecisao

func CalcularSaldoAviso(salario float64, avisoPrevio AvisoPrevio, justaCusa bool) (aviso float64) {
	if justaCusa {
		return
	}
	switch avisoPrevio {
	case Trabalhado:
		return
	case NãoCumprido:
		return -salario
	default:
		return salario
	}
}
