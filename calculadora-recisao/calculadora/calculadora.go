package calculadorarecisao

import (
	"time"
)

type AvisoPrevio int
type Motivo int

const (
	Trabalhado  = 0
	Indenizado  = 1
	Dispensado  = 2
	NãoCumprido = 3

	DemissãoSemJustaCausa		= 4
	DemissãoComJustaCausa 		= 5
	FuncionárioPediuDemissão 	= 6
)


type RequisiçãoRecisão struct {
	Salário             float64
	DataContratação     time.Time
	DataDemissão        time.Time
	AvisoPrevio         AvisoPrevio
	Motivo              Motivo
	JustaCausa          bool
	FériasVencidas		bool
}

type RequisiçãoResposta struct {
	SaldoFGTS            float64
	SaldoAviso           float64
	SaldoSalário         float64
	SaldoFérias          float64
	SaldoTotal           float64
}

func (req *RequisiçãoResposta) getSaldoTotal () {
	req.SaldoTotal = req.SaldoFGTS + req.SaldoAviso + req.SaldoSalário + req.SaldoFérias
}


func CalcularRecisão(requisição RequisiçãoRecisão) RequisiçãoResposta {
	resposta := RequisiçãoResposta{
		SaldoFGTS:            CalcularSaldoFGTS(requisição.Salário, requisição.DataContratação, requisição.DataDemissão, requisição.AvisoPrevio, requisição.Motivo, requisição.JustaCausa),
		SaldoAviso:           CalcularSaldoAviso(requisição.Salário, requisição.AvisoPrevio, requisição.JustaCausa),
		SaldoSalário:         CalcularSaldoSalário(requisição.Salário, requisição.DataContratação, requisição.DataDemissão, requisição.AvisoPrevio),
		SaldoFérias:          CalcularSaldoFérias(requisição.Salário, requisição.DataContratação, requisição.DataDemissão, requisição.AvisoPrevio, requisição.FériasVencidas),
	}
	resposta.getSaldoTotal()	
	return resposta
} 

func Date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
